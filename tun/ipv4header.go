package tun

import (
	"fmt"
	"net"
	"strconv"
)

// 打印ip报文头的详情
func printIPv4Header(packet []byte) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("----->IP Header<----")
	printVersionIPv4(packet)
	printHeaderLenIPv4(packet)
	printServiceTypeIPv4(packet)
	printAllLenIPv4(packet)
	printIdentificationIPv4(packet)
	printFlagsIPv4(packet)
	printFragmentOffsetIPv4(packet)
	printTTLIPv4(packet)
	printProtocolIPv4(packet)
	printChecksumIPv4(packet)
	printSrcIPv4(packet)
	printDstIPv4(packet)
	fmt.Println("----->IP Header<---End---")
}

func getIPv4HeaderLen(packet []byte) int {
	header := packet[0]
	headerLen := header & 0x0f * 4
	hl, _ := strconv.Atoi(fmt.Sprintf("%d", headerLen))
	return hl
}

func printVersionIPv4(packet []byte) {
	header := packet[0]
	fmt.Printf("IPv4 Header--->Version:%d\n", header>>4)
}

func printHeaderLenIPv4(packet []byte) {
	header := packet[0]
	fmt.Printf("IPv4 Header--->HeaderLen:%d byte\n", header&0x0f*4)
}

func printServiceTypeIPv4(packet []byte) {
	st := packet[1]
	fmt.Printf("IPv4 Header--->ServiceType:%v\n", st)
}

func printAllLenIPv4(packet []byte) {
	fmt.Printf("IPv4 Header--->AllLen:%d\n", uint16(packet[2])<<8|uint16(packet[3]))
}

func printIdentificationIPv4(packet []byte) {
	id := uint16(packet[4])<<8 | uint16(packet[5])
	fmt.Printf("IPv4 Header--->Identification:%x\t%d\n", id, id)
}

func printFlagsIPv4(packet []byte) {
	// 向右移动5位，相当于去掉后5位。数值降低
	fmt.Printf("IPv4 Header--->Flags:%03b\n", packet[6]>>5)
}
func printFragmentOffsetIPv4(packet []byte) {
	// 向左移动3位，相当于将前3位去掉。数值可能增加
	fmt.Printf("IPv4 Header--->FragmentOffset:%013b\n", uint16(packet[6])<<3|uint16(packet[7]))
}

func printTTLIPv4(packet []byte) {
	fmt.Printf("IPv4 Header--->TTL:%d\n", packet[8])
}

func printProtocolIPv4(packet []byte) {
	fmt.Printf("IPv4 Header--->ProtocolType:%d\n", packet[9])
}

func printChecksumIPv4(packet []byte) {
	fmt.Printf("IPv4 Header--->Checksum:%d\n", uint16(packet[10])<<8|uint16(packet[11]))
}

func printSrcIPv4(packet []byte) net.IP {
	fmt.Printf("IPv4 Header--->SrcIP:%d.%d.%d.%d\n", packet[12], packet[13], packet[14], packet[15])
	return net.IPv4(packet[12], packet[13], packet[14], packet[15])
}

func printDstIPv4(packet []byte) net.IP {
	fmt.Printf("IPv4 Header--->DstIP:%d.%d.%d.%d\n", packet[16], packet[17], packet[18], packet[19])
	return net.IPv4(packet[16], packet[17], packet[18], packet[19])
}
