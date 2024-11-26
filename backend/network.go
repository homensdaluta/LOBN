package main

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type ArpMap struct {
	IP  string
	MAC string
}

func GetSystemInfo() []pcap.Interface {

	var devices []pcap.Interface
	devices, _ = pcap.FindAllDevs()
	return devices
}

func GetInterfaceInfo(name string) pcap.Interface {
	decodedValue, _ := url.QueryUnescape(name)
	var devices []pcap.Interface
	devices, _ = pcap.FindAllDevs()
	for _, inet := range devices {
		if decodedValue == inet.Name {

			return inet
		}
	}
	return pcap.Interface{}
}

func GetAllActiveIPs(ifaceName string, ifaceIP string, timeout int64, timebetweenPackets int64) []ArpMap {
	fmt.Println(ifaceName)
	handle, _ = pcap.OpenLive(ifaceName, 65535, false, 1*time.Second)
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	defer handle.Close()
	_ = handle.SetBPFFilter("arp and dst host " + ifaceIP)

	for i := 1; i < 254; i++ {
		time.Sleep(time.Duration(timebetweenPackets) * time.Millisecond)
		go createARPPacket(ifaceIP, i)
	}

	arpSlice := ARPReader(timeout, packetSource)
	sliceArp := make([]ArpMap, 0)

	for _, arpPacket := range arpSlice {
		srcMAC := strings.ToUpper(byteToMACString(arpPacket.SourceHwAddress))
		srcIp := byteToIpv4String(arpPacket.SourceProtAddress)
		sliceArp = append(sliceArp, ArpMap{srcIp, srcMAC})
	}

	return sliceArp
}

func ARPReader(timeout int64, packetSource *gopacket.PacketSource) []*layers.ARP {
	timeDelay := 4 * time.Second
	var arpSlice []*layers.ARP
	arpSlice = make([]*layers.ARP, 0)
	var endTime <-chan time.Time

	for loopVar := true; loopVar; {
		if loopVar && endTime == nil {
			endTime = time.After(timeDelay)
		}
		select {
		case <-endTime:
			endTime = nil
			loopVar = false
		default:
			packet, _ := packetSource.NextPacket()
			if packet != nil {
				test := packet.Layer(layers.LayerTypeARP).(*layers.ARP)
				arpSlice = append(arpSlice, test)
			}
			continue
		}
	}
	//fmt.Println(len(arpSlice))
	return arpSlice
}

func createARPPacket(ifaceIP string, ipTest int) {

	ifas, _ := net.Interfaces()
	eth := &layers.Ethernet{
		SrcMAC:       ifas[0].HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}

	arp := &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   ifas[0].HardwareAddr,
		SourceProtAddress: stringToByteIP(ifaceIP),
		DstHwAddress:      net.HardwareAddr{0, 0, 0, 0, 0, 0},
		DstProtAddress:    []byte{192, 168, 1, byte(ipTest)},
	}
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	gopacket.SerializeLayers(buf, opts, eth, arp)
	if err := handle.WritePacketData(buf.Bytes()); err != nil {
		fmt.Println(err)
	}
}
