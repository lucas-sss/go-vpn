package tun

import "fmt"

func printUDPHeader(packet []byte) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("----->UDP Header<----")
	printUDPSrcPort(packet)
	printUDPDstPort(packet)
	printUDPAllLen(packet)
	printUDPChecksum(packet)
	printUDPData(packet)
	fmt.Println("----->UDP Header<---End--")
}

func printUDPSrcPort(packet []byte) {
	fmt.Printf("UDP Header--->SrcPort:%d\n", uint16(packet[0])<<8|uint16(packet[1]))
}

func printUDPDstPort(packet []byte) {
	fmt.Printf("UDP Header--->DstPort:%d\n", uint16(packet[2])<<8|uint16(packet[3]))
}

func printUDPAllLen(packet []byte) {
	fmt.Printf("UDP Header--->AllLen:%d\n", uint16(packet[4])<<8|uint16(packet[5]))
}

func printUDPChecksum(packet []byte) {
	fmt.Printf("UDP Header--->Checksum:%d\n", uint16(packet[6])<<8|uint16(packet[7]))
}

func printUDPData(packet []byte) {
	fmt.Printf("UDP Header--->data:%v\n", string(packet[8:]))
}
