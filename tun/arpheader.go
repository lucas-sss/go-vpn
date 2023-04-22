package tun

import (
	"fmt"
	"net"
)

func printARPHeader(packet []byte) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("----->ARP Header<----")
	printARPTYPE(packet)
	printARPProTYPE(packet)
	printARPHardwareLen(packet)
	printARPPLen(packet)
	printARPOp(packet)
	printARPSrcHardwareMAC(packet)
	printARPSrcIP(packet)
	printARPDstHardwareMAC(packet)
	printARPDstIP(packet)
	fmt.Println("----->ARP Header<---END---")

}
func printARPTYPE(packet []byte) {
	fmt.Printf("ARP Header--->Type:%d\n", uint16(packet[0])<<8|uint16(packet[1]))
}

func printARPProTYPE(packet []byte) {
	fmt.Printf("ARP Header--->ProtocolType:%04x\n", uint16(packet[2])<<8|uint16(packet[3]))
}

func printARPHardwareLen(packet []byte) {
	fmt.Printf("ARP Header--->HardwareLen:%d\n", packet[4])
}

func printARPPLen(packet []byte) {
	fmt.Printf("ARP Header--->ProtocolLen:%d\n", packet[5])
}

func printARPOp(packet []byte) {
	fmt.Printf("ARP Header--->op:%d\n", uint16(packet[6])<<8|uint16(packet[7]))
}

func printARPSrcHardwareMAC(packet []byte) {
	fmt.Printf("ARP Header--->SrcHardwarMAC:%x:%x:%x:%x:%x:%x\n", packet[8], packet[9], packet[10], packet[11], packet[12], packet[13])
}

func printARPSrcIP(packet []byte) {
	fmt.Printf("ARP Header--->SrcIP:%d.%d.%d.%d\n", packet[14], packet[15], packet[16], packet[17])
}

func printARPDstHardwareMAC(packet []byte) {
	fmt.Printf("ARP Header--->DstHardwareMAC:%x:%x:%x:%x:%x:%x\n", packet[18], packet[19], packet[20], packet[21], packet[22], packet[23])
}

func printARPDstIP(packet []byte) {
	fmt.Printf("ARP Header--->DstIP:%d.%d.%d.%d\n", packet[24], packet[25], packet[26], packet[27])
}

func GetIPv4SrcARP(packet []byte) net.IP {
	return net.IPv4(packet[14], packet[15], packet[16], packet[17])
}

func GetIPv4DstARP(packet []byte) net.IP {
	return net.IPv4(packet[24], packet[25], packet[26], packet[27])
}
