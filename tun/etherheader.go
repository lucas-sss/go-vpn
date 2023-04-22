package tun

import (
	"fmt"
	"net"
)

func printMACHeader(packet []byte) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("----->MAC Header<----")
	printSrcMACEther(packet)
	printDstMACEther(packet)
	printTypeEther(packet)
	fmt.Println("----->MAC Header<--end--")
}

func printDstMACEther(packet []byte) {
	fmt.Printf("Ether Header--->DstMac:%x:%x:%x:%x:%x:%x\n", packet[0], packet[1], packet[2], packet[3], packet[4], packet[5])
}

func printSrcMACEther(packet []byte) {
	fmt.Printf("Ether Header--->SrcMac:%x:%x:%x:%x:%x:%x\n", packet[6], packet[7], packet[8], packet[9], packet[10], packet[11])
}
func printTypeEther(packet []byte) {
	fmt.Printf("Ether Header--->Type:%04x\n", uint16(packet[12])<<8|uint16(packet[13]))
}

func MACDestination(macFrame []byte) net.HardwareAddr {
	return net.HardwareAddr(macFrame[:6])
}

func MACSource(macFrame []byte) net.HardwareAddr {
	return net.HardwareAddr(macFrame[6:12])
}

// 读取以太网帧的第12、13个字节，这两个字节代表数据类型
func MACType(macFrame []byte) []byte {
	return macFrame[12:14]
}

func MACTypeARP(macFrame []byte) []byte {
	return macFrame[2:4]
}
