package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/gopacket/pcap"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var (
	err     error
	handle  *pcap.Handle
	ipAddr  net.IP
	macAddr net.HardwareAddr
	target  string
)

func main() {
	start()
	//massivePing()
	/*
		n := 0
		for n < 254 {
			addr := "192.168.1." + strconv.Itoa(n)
			dst, dur, err := Ping(addr)
			if err != nil {
				panic(err)
			}
			log.Printf("Ping %s (%s): %s\n", addr, dst, dur)
			n++
		}*/
}

func grabAddresses(iface string) (macAddr net.HardwareAddr, ipAddr net.IP) {

	netInterface, err := net.InterfaceByName(iface)
	checkError(err)

	macAddr = netInterface.HardwareAddr
	addrs, _ := netInterface.Addrs()
	ipAddr, _, err = net.ParseCIDR(addrs[0].String())
	checkError(err)

	return macAddr, ipAddr
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// it works but its not concurrent
func massivePing() {

	connection, err := icmp.ListenPacket("ip4:icmp", ListenAddr)
	if err != nil {
		fmt.Errorf("Error listening to %v\n", ListenAddr)
		panic(err)
	}
	defer connection.Close()

	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(""),
		},
	}
	messageByte, err := message.Marshal(nil)
	if err != nil {
		fmt.Errorf("Error listening to creating message\n")
		panic(err)
	}

	point := 0
	for point < 254 {
		go pingTest(point, connection, messageByte)
		/*
			addr := "192.168.1." + strconv.Itoa(point)
			dst, err := net.ResolveIPAddr("ip4", addr)
			if err != nil {
				fmt.Errorf("Something went wrong resolving address %v\n", addr)
				panic(err)
			}

			start := time.Now()
			n, err := connection.WriteTo(messageByte, dst)
			if err != nil {
				fmt.Errorf("Something went wrong sending to: %v\n", addr)
				panic(err)
			}

			reply := make([]byte, 1500)
			err = connection.SetReadDeadline(time.Now().Add(10 * time.Second))
			if err != nil {
				fmt.Errorf("could not read reply from %v\n", addr)
				panic(err)
			}
			n, peer, err := connection.ReadFrom(reply)
			_ = peer
			if err != nil {
				fmt.Errorf("could not open reply from  %v\n", addr)
				panic(err)
			}
			duration := time.Since(start)

			// Pack it up boys, we're done here
			rm, err := icmp.ParseMessage(ProtocolICMP, reply[:n])
			if err != nil {
				fmt.Errorf("could not unpack to icmp from %v\n", addr)
				panic(err)
			}
			switch rm.Type {
			case ipv4.ICMPTypeEchoReply:
				log.Printf("Ping %s (%s): %s\n", addr, dst, duration)
			default:
				log.Printf("Ping to %s failed\n", addr)
			}*/
		point++
	}
	time.Sleep(60 * time.Second)
}

func pingTest(point int, connection *icmp.PacketConn, messageByte []byte) {
	addr := "192.168.1." + strconv.Itoa(point)
	dst, err := net.ResolveIPAddr("ip4", addr)
	if err != nil {
		fmt.Errorf("Something went wrong resolving address %v\n", addr)
		panic(err)
	}

	start := time.Now()
	n, err := connection.WriteTo(messageByte, dst)
	if err != nil {
		fmt.Errorf("Something went wrong sending to: %v\n", addr)
		panic(err)
	}

	reply := make([]byte, 1500)
	err = connection.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Errorf("could not read reply from %v\n", addr)
		panic(err)
	}
	n, peer, err := connection.ReadFrom(reply)
	_ = peer
	if err != nil {
		fmt.Errorf("could not open reply from  %v\n", addr)
		panic(err)
	}
	duration := time.Since(start)

	// Pack it up boys, we're done here
	rm, err := icmp.ParseMessage(ProtocolICMP, reply[:n])
	if err != nil {
		fmt.Errorf("could not unpack to icmp from %v\n", addr)
		panic(err)
	}
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("Ping %s (%s): %s\n", addr, dst, duration)
	default:
		log.Printf("Ping to %s failed\n", addr)
	}
	point++
}

////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	// Stolen from https://godoc.org/golang.org/x/net/internal/iana,
	// can't import "internal" packages
	ProtocolICMP = 1
	//ProtocolIPv6ICMP = 58
)

// Default to listen on all IPv4 interfaces
var ListenAddr = "0.0.0.0"

// Mostly based on https://github.com/golang/net/blob/master/icmp/ping_test.go
// All ye beware, there be dragons below...

func Ping(addr string) (*net.IPAddr, time.Duration, error) {
	// Start listening for icmp replies
	c, err := icmp.ListenPacket("ip4:icmp", ListenAddr)
	if err != nil {
		return nil, 0, err
	}
	defer c.Close()

	// Resolve any DNS (if used) and get the real IP of the target
	dst, err := net.ResolveIPAddr("ip4", addr)
	if err != nil {
		panic(err)
		return nil, 0, err
	}

	// Make a new ICMP message
	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1, //<< uint(seq), // TODO
			Data: []byte(""),
		},
	}
	b, err := m.Marshal(nil)
	if err != nil {
		return dst, 0, err
	}

	// Send it
	start := time.Now()
	n, err := c.WriteTo(b, dst)
	if err != nil {
		return dst, 0, err
	} else if n != len(b) {
		return dst, 0, fmt.Errorf("got %v; want %v", n, len(b))
	}

	// Wait for a reply
	reply := make([]byte, 1500)
	err = c.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return dst, 0, err
	}
	n, peer, err := c.ReadFrom(reply)
	_ = peer
	if err != nil {
		return dst, 0, err
	}
	duration := time.Since(start)

	// Pack it up boys, we're done here
	rm, err := icmp.ParseMessage(ProtocolICMP, reply[:n])
	if err != nil {
		return dst, 0, err
	}
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		return dst, duration, nil
	default:
		return dst, 0, nil
	}
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
