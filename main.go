package main

import (
	"flag"
	"fmt"
	"go-vpn/tun"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	DEFAULT_TUN_NAME   = "tun1"
	DEFAULT_TUN_DEVICE = "/dev/net/tun"
	ifnameSize         = 16
)

var tunName string
var tunDevice string
var bindAddr string
var remoteAddr string
var tunIP string

func init() {
	flag.StringVar(&tunName, "tunname", DEFAULT_TUN_NAME, "The name of tun device")
	flag.StringVar(&tunDevice, "tundev", DEFAULT_TUN_DEVICE, "The file descriptor of tun device")
	flag.StringVar(&bindAddr, "bindaddr", "", "The address[ip:port] is bound when the go-vpn service is running, the ip can not use 127.0.0.1 or localhost")
	flag.StringVar(&remoteAddr, "remoteaddr", "", "Peer go-vpn service address[ip:port]")
	flag.StringVar(&tunIP, "tunip", "", "The file descriptor of tun device")
	flag.Parse()
	if "" == bindAddr {
		fmt.Printf("requird param of bindaddr[example: 192.168.10.2:8181]\n")
		os.Exit(1)
	}
	if "" == remoteAddr {
		fmt.Printf("requird param of remoteaddr[example: 192.168.20.2:8282]\n")
		os.Exit(1)
	}
	if "" == tunIP {
		fmt.Printf("requird param of tunip[example: 10.9.0.2]\n")
		os.Exit(1)
	}
}

var exitChan chan os.Signal

func exitHandle() {
	s := <-exitChan
	fmt.Println("\nCapture signal: ", s)
	//关闭tun设备
	tun.CloseTun()
	fmt.Println("go-vpn exit")
	os.Exit(0) //使用os.Exit强行关掉
}

func main() {
	exitChan = make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	go exitHandle()

	fmt.Printf("======>Now----Tun---VPN---UDP<======\n")
	tunFile, err := tun.CreateTun(tunName, tunDevice, tunIP)
	if err != nil {
		fmt.Printf("ICMP Listen Packet Failed! err:%v\n", err.Error())
		return
	}
	defer tunFile.Close()

	udpConn, err := tun.CreateUDP(bindAddr)
	if err != nil {
		fmt.Printf("UDP conn Failed! err:%v\n", err.Error())
		return
	}
	defer udpConn.Close()

	go tun.TunToUDP(udpConn, remoteAddr, tunFile)

	go tun.UdpToTun(udpConn, tunFile)

	time.Sleep(time.Hour)
}
