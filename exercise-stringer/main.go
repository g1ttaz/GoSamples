package main

import "fmt"

// IPAddr is a IP v4 address
type IPAddr [4]byte

func (ipaddr *IPAddr) String() string {
	if ipaddr == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%d.%d.%d.%d", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
