package tun

import (
	"encoding/binary"
	"fmt"
	"go-vpn/common"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/net/icmp"
)

func CreateUDP(localAddr string) (*net.UDPConn, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		log.Fatalln("failed to get udp socket:", err)
		return nil, err
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln("failed to listen on udp socket:", err)
		return nil, err
	}

	return conn, nil

}

func TunToUDP(udpConn *net.UDPConn, remoteAddr string, tunFile *os.File) {
	packet := make([]byte, 1024*64)
	size := 0
	var err error
	for {
		//读取到以太网帧
		if size, err = tunFile.Read(packet); err != nil {
			return
		}
		//获取以太网帧内部数据类型
		// 0x0800代表IP协议帧
		// 0x0806代表ARP协议帧
		// 0x8864代表PPPoE
		// 0x86dd代表IPv6
		// te := MACType(packet[:size])
		mt := fmt.Sprintf("%x", MACType(packet[:size]))
		fmt.Println("以太网帧类型：", mt)
		//动态获取ipv4头部长度
		printMACHeader(packet[:size])

		//16进制0800代表ip报文
		if strings.EqualFold(mt, "0800") {
			b := packet[14:size]
			printIPv4Header(b)
			hl := getIPv4HeaderLen(b)
			if b[9] == 1 { //icmp
				icmpPacket := b[hl:]
				printICMPHeader(icmpPacket)
			}

			if b[9] == 6 { //tcp
				tcpPacket := b[hl:]
				printTCPHeader(tcpPacket)
			}

			if b[9] == 17 { //udp
				udpPacket := b[hl:]
				printUDPHeader(udpPacket)
			}
		}

		//arp报文
		if strings.EqualFold(mt, "0806") {
			b := packet[14:size]
			printARPHeader(b)
		}

		// b := packet[:size]
		// srcIP := common.GetSrcIP(b)
		// dstIP := common.GetDstIP(b)
		// fmt.Printf("tunToUDP--->Msg Protocol type: %v(1=ICMP, 6=TCP, 17=UDP)\tsrcIP:%v--->dstIP:%v", packet[9], srcIP, dstIP)

		rAddr, err := net.ResolveUDPAddr("udp", remoteAddr)
		if err != nil {
			log.Fatalln("failed to get udp socket:", err)
			return
		}
		if size, err = udpConn.WriteTo(packet[:size], rAddr); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("tunToUDP--->Write Msg To UDP Conn OK! size:%d\n", size)
	}
}

func UdpToTun(udpConn *net.UDPConn, tunFile *os.File) {
	var packet = make([]byte, 1024*64)
	var size int
	var err error
	var addr net.Addr

	for {
		if size, addr, err = udpConn.ReadFrom(packet); err != nil {
			continue
		}

		size, err = tunFile.Write(packet[:size])
		if err != nil {
			continue
		}
		fmt.Printf("udpToTun--->Write Msg To /dev/net/tun OK! size:%d\tsrcIP:%v\n", size, addr)
	}
}

func TunToIcmp(icmpconn *icmp.PacketConn, tunFile *os.File) {
	var srcIP string
	packet := make([]byte, 1024*64)
	size := 0
	var err error
	for {
		if size, err = tunFile.Read(packet); err != nil {
			return
		}
		fmt.Printf("Msg Length: %d\n", binary.BigEndian.Uint16(packet[2:4]))
		fmt.Printf("Msg Protocol: %d (1=ICMP, 6=TCP, 17=UDP)\tsize:%d\n", packet[9], size)

		b := packet[:size]
		srcIP = common.GetSrcIP(b)
		dstIP := common.GetDstIP(b)
		fmt.Printf("Msg srcIP: %s\tdstIP:%v\n", srcIP, dstIP)

		var raddr = net.IPAddr{IP: net.ParseIP(dstIP)}

		b = b[20:size]

		if size, err = icmpconn.WriteTo(b, &raddr); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Write Msg To Icmp Conn OK! size:%d\n", size)
	}
}

func IcmpToTun(tunIp string, icmpconn *icmp.PacketConn, tunFile *os.File) {
	var sb = make([]byte, 1024*64)
	var addr net.Addr
	var size int
	var err error

	for {
		if size, addr, err = icmpconn.ReadFrom(sb); err != nil {
			continue
		}

		ipHeader := common.CreateIPv4Header(net.ParseIP(addr.String()), net.ParseIP(tunIp), os.Getpid())
		iphb, err := ipHeader.Marshal()
		if err != nil {
			continue
		}
		fmt.Printf("Reply MSG Length: %d\n", binary.BigEndian.Uint16(iphb[2:4]))
		fmt.Printf("Reply MSG Protocol: %d (1=ICMP, 6=TCP, 17=UDP)\n", iphb[9])
		dstIP := common.GetDstIP(iphb)
		fmt.Printf("Reply src IP: %s\tdstIP:%v\n", addr, dstIP)

		var rep = make([]byte, 84)
		rep = append(iphb, sb[:size]...)

		size, err = tunFile.Write(rep)
		if err != nil {
			continue
		}
		fmt.Printf("Write Msg To /dev/net/tun OK! size:%d\ttime:%v\n", size, time.Now())
	}
}
