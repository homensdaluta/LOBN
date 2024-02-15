package main

import (
	"fmt"
	"net"
	"net/url"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

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
			GetAllActiveIPs(inet)
			return inet
		}
	}
	return pcap.Interface{}
}

func GetAllActiveIPs(iface pcap.Interface) {
	fmt.Println(iface)

	handle, _ = pcap.OpenLive(iface.Name, 65535, false, pcap.BlockForever)
	defer handle.Close()
	_ = handle.SetBPFFilter("arp and src host 192.168.1 and dst host 192.168.1.74")
	for i := 2; i < 100; i++ {
		_ = handle.SetBPFFilter("arp and src host 192.168.1 and dst host 192.168.1.192")
		//	ifas, _ := net.Interfaces()
		for {
			createARPPacket(&iface, 192)
			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
			packet, _ := packetSource.NextPacket()
			test := packet.Layer(layers.LayerTypeARP)
			arptest, _ := test.(*layers.ARP)

			fmt.Println(arptest.SourceHwAddress)
			continue
		}
	}
	/*
		if err := writeARP(handle, iface, addr); err != nil {
			log.Printf("error writing packets on %v: %v", iface.Name, err)
			return err
		}*/

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	//packet, _ := packetSource.NextPacket()
	//test := packet.Layer(layers.LayerTypeARP)

	for {
		packet, _ := packetSource.NextPacket()
		test := packet.Layer(layers.LayerTypeARP)
		arptest, _ := test.(*layers.ARP)
		fmt.Println("test")

		fmt.Println(arptest.SourceHwAddress)

		continue
	}

}

func ARPReader() {
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println()
	for {
		packet, _ := packetSource.NextPacket()
		fmt.Println(packet)
		continue
	}

}

func createARPPacket(iface *pcap.Interface, ipTest int) {

	ifas, _ := net.Interfaces()
	fmt.Println(ifas[0].HardwareAddr)
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
		SourceProtAddress: iface.Addresses[0].IP,
		DstHwAddress:      net.HardwareAddr{0, 0, 0, 0, 0, 0},
		DstProtAddress:    []byte{192, 168, 1, byte(ipTest)},
	}
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	fmt.Println("antes")
	fmt.Println(buf)
	gopacket.SerializeLayers(buf, opts, eth, arp)
	fmt.Println("depois")
	fmt.Println(buf)
	if err := handle.WritePacketData(buf.Bytes()); err != nil {
		fmt.Println(err)
	}
}

/*
func ipToHexaByte(ip string) []byte {
	arrayIp := strings.Split(ip, ".")
	var byteArray []byte
	octet0, _ := strconv.Atoi(arrayIp[0])
	octet1, _ := strconv.Atoi(arrayIp[1])
	octet2, _ := strconv.Atoi(arrayIp[2])
	octet3, _ := strconv.Atoi(arrayIp[3])

	b := [4]byte{byte(octet0), byte(octet1), byte(octet2), byte(octet3)}
	fmt.Println(b)
	/*
		for _, ipSplit := range arrayIp {
			a := net.ParseIP(ipSplit)
			fmt.Printf("%b", net.IP.To4(a))
			//hexSplit := hex.EncodeToString(ipSplit)
			hexSplit := []byte(ipSplit)
			fmt.Println(hexSplit)
			//byteArray = append(byteArray, hexSplit[0])
		}
	return byteArray
}
*/
