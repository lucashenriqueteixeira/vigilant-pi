package main

import (
	"fmt"
	"log"

	"github.com/petersondmg/vigilant-pi/lib/config"
	"github.com/petersondmg/vigilant-pi/lib/network"
)

func main() {
	settings := config.Read()
	cidrIP := network.GetInterfaceIPv4CIDR(settings.Interface)
	devices, err := network.ScanOnIPv4WithCIDR(settings.Interface, cidrIP)

	if err != nil {
		log.Fatalf("Error scanning network: %s", err.Error())
	}

	for _, device := range devices {
		fmt.Printf("%s - %s - %s\n", device.IP, device.MAC, device.Manufacturer)
	}
}
