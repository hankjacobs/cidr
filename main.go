package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: cidr <CIDR address>")
		os.Exit(1)
	}
	cidr := os.Args[1]
	start, end, err := getIPRange(cidr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%v - %v\n", start, end)
}

func getIPRange(cidr string) (start net.IP, end net.IP, err error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, nil, err
	}

	start = ip.Mask(ipnet.Mask)
	ones, total := ipnet.Mask.Size()
	zeros := total - ones
	count := (1 << uint32(zeros)) // 2 ^ zeros

	end = make(net.IP, len(start))
	copy(end, start)

	i := len(end) - 1
	for i >= 0 {
		if count%256 == 0 {
			end[i] = 255
			count /= 256
		} else {
			end[i] += byte(count - 1)
			break
		}
		i--
	}

	return
}
