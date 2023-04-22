package tun

import (
	"fmt"
)

func printTCPHeader(packet []byte) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("----->TCP Header<----len:%d\n", len(packet[:]))
	printTCPsrcPort(packet)
	printTCPdstPort(packet)
	printTCPSequenceNumber(packet)
	printTCPACKNumber(packet)
	printTCPHeaderLen(packet)
	printTCPFlagCWR(packet)
	printTCPFlagECE(packet)
	printTCPFlagUrgent(packet)
	printTCPFlagACK(packet)
	printTCPFlagPSH(packet)
	printTCPFlagRST(packet)
	printTCPFlagSYN(packet)
	printTCPFlagFIN(packet)
	printTCPWindowSize(packet)
	printTCPCheckSum(packet)
	printTCPUrgentPointer(packet)
	printTCPData(packet)
	fmt.Println("----->TCP Header<---END---")
}

func printTCPsrcPort(packet []byte) {
	fmt.Printf("TCP Header--->SrcPort:%d\n", int64(uint16(packet[0])<<8|uint16(packet[1])))
}

func printTCPdstPort(packet []byte) {
	fmt.Printf("TCP Header--->DstPort:%d\n", int64(uint16(packet[2])<<8|uint16(packet[3])))
}

func printTCPSequenceNumber(packet []byte) {
	fmt.Printf("TCP Header--->SequenceNumber:%d\n", uint32(packet[4])<<24|uint32(packet[5])<<16|uint32(packet[6])<<8|uint32(packet[7]))
}

func printTCPACKNumber(packet []byte) {
	fmt.Printf("TCP Header--->ACKNum:%d\n", uint32(packet[8])<<24|uint32(packet[9])<<16|uint32(packet[10])<<8|uint32(packet[11]))
}

func printTCPHeaderLen(packet []byte) {
	fmt.Printf("TCP Header--->HeaderLen:%d\n", packet[12]>>4*4)
}

func printTCPFlagCWR(packet []byte) {
	fmt.Printf("TCP Header--->FlagCWR:%d\n", packet[13]&0x80>>7)
}
func printTCPFlagECE(packet []byte) {
	fmt.Printf("TCP Header--->FlagEcho:%d\n", packet[13]&0x40>>6)
}

func printTCPFlagUrgent(packet []byte) {
	fmt.Printf("TCP Header--->FlagUrgent:%d\n", packet[13]&0x20>>5)
}

func printTCPFlagACK(packet []byte) {
	fmt.Printf("TCP Header--->FlagACK:%d\n", packet[13]>>4&0b0001)
	fmt.Printf("TCP Header--->FlagACK:%d\n", packet[13]>>4&0b1)
	//fmt.Printf("TCP Header--->FlagACK:%d\tpacket:%v\n", packet[13]>>4, packet[13])
	fmt.Printf("TCP Header--->FlagACK:%d\n", packet[13]&0x10>>4)
}

func printTCPFlagPSH(packet []byte) {
	fmt.Printf("TCP Header--->FlagPSH:%d\n", packet[13]&0x08>>3)
}

func printTCPFlagRST(packet []byte) {
	fmt.Printf("TCP Header--->FlagRST:%d\n", packet[13]&0x04>>2)
}

func printTCPFlagSYN(packet []byte) {
	fmt.Printf("TCP Header--->FlagSYN:%d\n", packet[13]&0x02>>1)
}

func printTCPFlagFIN(packet []byte) {
	fmt.Printf("TCP Header--->FlagFIN:%d\n", packet[13]&0x01)
}

func printTCPWindowSize(packet []byte) {
	fmt.Printf("TCP Header--->WindowSize:%d\n", uint16(packet[14])<<8|uint16(packet[15]))
}

func printTCPCheckSum(packet []byte) {
	fmt.Printf("TCP Header--->Checksum:%04x\n", uint16(packet[16])<<8|uint16(packet[17]))
}

func printTCPUrgentPointer(packet []byte) {
	fmt.Printf("TCP Header--->UrgentPointer:%d\n", uint16(packet[18])<<8|uint16(packet[19]))
}

func printTCPData(packet []byte) {
	headerLen := packet[12] >> 4 * 4
	p := packet[headerLen:]
	dataLen := len(p)
	if dataLen > 1 {
		fmt.Printf("TCP Header--->Data--->:%v\theaderLen:%v\tdataLen:%d\n", string(p[:dataLen-1]), headerLen, dataLen)
	}
}
