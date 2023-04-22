package tun

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

const (
	defaultTunDev = "/dev/net/tun"
	ifnameSize    = 16
)

type ifreqFlags struct {
	IfrnName  [ifnameSize]byte
	IfruFlags uint16
}

var tunFile *os.File

func CreateTun(tunName, tunDevice, tunIP string) (*os.File, error) {
	err := addTun(tunName)
	if err != nil {
		return nil, err
	}

	err = configTun(tunName, tunIP)
	if err != nil {
		return nil, err
	}

	tunFile, ifName, err := openTun(tunName, tunDevice)
	if err != nil {
		return nil, err
	}
	fmt.Printf("create tun interface name: %s\n", ifName)

	return tunFile, nil
}

// func ioctl(fd int, request, argp uintptr) error {
// 	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), request, argp)
// 	if errno != 0 {
// 		fmt.Errorf("ioctl failed with '%s'\n", errno)
// 		return fmt.Errorf("ioctl failed with '%s'", errno)
// 	}
// 	return nil
// }

// func fromZeroTerm(s []byte) string {
// 	return string(bytes.TrimRight(s, "\000"))
// }

func openTun(name, tunDevice string) (*os.File, string, error) {
	tun, err := os.OpenFile(tunDevice, os.O_RDWR|syscall.O_NONBLOCK, 0)
	if err != nil {
		fmt.Printf("OpenTun Failed! err:%v", err.Error())
		return nil, "", err
	}
	tunFile = tun

	// var ifr ifreqFlags
	// copy(ifr.IfrnName[:len(ifr.IfrnName)-1], []byte(name+"\000"))
	// ifr.IfruFlags = syscall.IFF_TUN | syscall.IFF_NO_PI

	// err = ioctl(int(tun.Fd()), syscall.TUNSETIFF, uintptr(unsafe.Pointer(&ifr)))
	// if err != nil {
	// 	fmt.Printf("OpenTun Failed! err:%v\n", err.Error())
	// 	return nil, "", err
	// }

	// return tun, ifName, nil

	// Create a new TUN device
	ifr, err := unix.NewIfreq(name)
	if err != nil {
		return nil, "", err
	}
	// ifr.flags = IFF_TUN
	ifr.SetUint16(unix.IFF_TUN)

	_, _, errno := unix.Syscall(unix.SYS_IOCTL, tun.Fd(), uintptr(unix.TUNSETIFF), uintptr(unsafe.Pointer(ifr)))
	if errno != 0 {
		return nil, "", fmt.Errorf(errno.Error())
	}

	// Set Tun Interface Not Persistent
	_, _, errno = unix.Syscall(unix.SYS_IOCTL, tun.Fd(), uintptr(unix.TUNSETPERSIST), uintptr(0))
	if errno != 0 {
		return nil, "", fmt.Errorf(errno.Error())
	}
	return tun, ifr.Name(), nil
}

func addTun(tunName string) error {
	la := netlink.LinkAttrs{
		Name:  tunName,
		Index: 8,
		MTU:   1500,
	}
	tun := netlink.Tuntap{
		LinkAttrs: la,
		Mode:      netlink.TUNTAP_MODE_TUN,
	}

	l, err := netlink.LinkByName(tunName)
	if err == nil {
		netlink.LinkSetDown(l)
		netlink.LinkDel(l)
	}

	err = netlink.LinkAdd(&tun)
	if err != nil {
		return err
	}
	return nil
}

func configTun(tunName, tunIP string) error {
	l, err := netlink.LinkByName(tunName)
	if err != nil {
		return err
	}

	ip, err := netlink.ParseIPNet(fmt.Sprintf("%s/%d", tunIP, 24))
	if err != nil {
		return err
	}

	addr := &netlink.Addr{IPNet: ip, Label: ""}
	if err = netlink.AddrAdd(l, addr); err != nil {
		return err
	}

	err = netlink.LinkSetUp(l)
	if err != nil {
		return err
	}

	return nil
}

func CloseTun() {
	if tunFile != nil {
		tunFile.Close()
	}
}
