package common

import (
	"net"

	"golang.org/x/net/ipv4"
)

func CreateIPv4Header(src, dst net.IP, id int) *ipv4.Header {

	iph := &ipv4.Header{
		Version:  ipv4.Version,
		Len:      ipv4.HeaderLen,
		TOS:      0x00,
		TotalLen: ipv4.HeaderLen + 64,
		ID:       id,
		Flags:    ipv4.DontFragment,
		FragOff:  0,
		TTL:      64,
		Protocol: 1,
		Checksum: 0,
		Src:      src,
		Dst:      dst,
	}

	h, _ := iph.Marshal()

	iph.Checksum = int(checkSum(h))

	return iph
}

func checkSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += sum >> 16

	return uint16(^sum)
}

func GetSrcIP(packet []byte) string {
	key := ""
	if IsIPv4(packet) && len(packet) >= 20 {
		key = GetIPv4Src(packet).To4().String()
	}

	return key
}

func IsIPv4(packet []byte) bool {
	flag := packet[0] >> 4
	return flag == 4
}

func GetDstIP(packet []byte) string {
	key := ""
	if IsIPv4(packet) && len(packet) >= 20 {
		key = GetIPv4Dst(packet).To4().String()
	}

	return key
}

func GetIPv4Src(packet []byte) net.IP {
	return net.IPv4(packet[12], packet[13], packet[14], packet[15])
}

func GetIPv4Dst(packet []byte) net.IP {
	return net.IPv4(packet[16], packet[17], packet[18], packet[19])
}
