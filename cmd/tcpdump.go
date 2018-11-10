package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func tracePacket() {
	handle, err := pcap.OpenLive(device, snapshopLen, promiscuous, timeout)
	if err != nil {
		log.Fatalf("failed to open device %s, err: %v\n", device, err)
	}
	defer handle.Close()

	err = handle.SetBPFFilter(tcpdumpFilter)
	if err != nil {
		log.Fatalf("failed to set monitoring filter '%s', err: %v\n", tcpdumpFilter, err)
	}

	packetSrc := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSrc.Packets() {
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {
	if dump {
		fmt.Println(packet.Dump())
	}

	var (
		proto   layers.IPProtocol
		srcIP   net.IP
		srcPort layers.TCPPort
		dstIP   net.IP
		dstPort layers.TCPPort
		payload []byte
	)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		srcIP = ip.SrcIP
		dstIP = ip.DstIP
		proto = ip.Protocol
	}
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		srcPort = tcp.SrcPort
		dstPort = tcp.DstPort
		payload = tcp.LayerPayload()
	}
	cy := color.New(color.FgCyan).SprintfFunc()
	mag := color.New(color.FgMagenta).SprintfFunc()
	if len(payload) == 0 {
		return
	}
	fmt.Printf("%s - %s:%s >> %s:%s\n", proto, srcIP, cy("%d", srcPort), dstIP, mag("%d", dstPort))
	color.New(color.FgGreen).Printf("\n%s\n", payload)

	if app {
		appLayer := packet.ApplicationLayer()
		if appLayer != nil {
			ap := appLayer.Payload()
			color.New(color.FgHiGreen).Println(ap)
		}
	}

	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("error decoding some parts of the packet:", err)
	}

	fmt.Println("--------------------\n")
}
