package network

import (
	"log"
	"net"
	"strings"
)

// GetInterfaceIPv4CIDR returns the IPv4 of the given
// inteface. Fatal if not found
func GetInterfaceIPv4CIDR(iface string) string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Can't get current interfaces %s", err.Error())
	}
	for _, i := range ifaces {
		if i.Name != iface {
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			log.Fatalf("Can't get addrs of interface %s", err.Error())
		}

		for _, addr := range addrs {
			// var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				strIP := v.String()
				if strings.Contains(strIP, ".") {
					return strIP
				}
				// case *net.IPAddr:
				// 	ip = v.IP
			}
		}

		log.Fatalf("Can't get addrs of interface %s. Checkout config file (config.yml if not changed in env)", iface)
	}

	log.Fatalf("Interface %s not found. Checkout config file (config.yml if not chanced in env)", iface)
	return ""
}
