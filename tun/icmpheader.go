package tun

import "fmt"

func printICMPHeader(packet []byte) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("----->ICMP Header<----")
	printICMPType(packet)
	printICMPCode(packet)
	printICMPCheckSum(packet)
	printICMPIdentification(packet)
	printICMPSeqNum(packet)
	printICMPTimestamp(packet)
	printICMPData(packet)
	fmt.Println("----->ICMP Header<----End----")
}

func printICMPType(packet []byte) {
	fmt.Printf("ICMP Header--->Type:%d\n", packet[0])
}

func printICMPCode(packet []byte) {
	fmt.Printf("ICMP Header--->Code:%d\n", packet[1])
}

func printICMPCheckSum(packet []byte) {
	fmt.Printf("ICMP Header--->Checksum:%04x\n", uint16(packet[2])<<8|uint16(packet[3]))
}

func printICMPIdentification(packet []byte) {
	fmt.Printf("TICMP Header--->Identification(process):%d\n", uint16(packet[4])<<8|uint16(packet[5]))
}

func printICMPSeqNum(packet []byte) {
	fmt.Printf("ICMP Header--->SeqNum:%d\n", uint16(packet[6])<<8|uint16(packet[7]))
}

func printICMPTimestamp(packet []byte) {
	fmt.Printf("ICMP Header --->timestamp:%02x %02x %02x %02x %02x %02x %02x %02x\n", packet[8], packet[9], packet[10], packet[11], packet[12], packet[13], packet[14], packet[15])
}

func printICMPData(packet []byte) {
	fmt.Printf("ICMP Header--->data:%v\n", string(packet[20:]))
}
