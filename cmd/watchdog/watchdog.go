package watchdog

import (
	"fmt"
	"log"

	"github.com/petersondmg/vigilant-pi/lib/config"
	"github.com/petersondmg/vigilant-pi/lib/network"
	"github.com/urfave/cli"
)

// Command ...
func Command() cli.Command {
	return cli.Command{
		Name:        "watch",
		Description: "starts monitoring and recording videos",
		Action: func(c *cli.Context) error {
			search()
			return nil
		},
	}
}

func search() {
	cidrIP := network.GetInterfaceIPv4CIDR(config.Current.Interface)
	devices, err := network.ScanOnIPv4WithCIDR(config.Current.Interface, cidrIP)

	if err != nil {
		log.Fatalf("Error scanning network: %s", err.Error())
	}

	for _, device := range devices {
		fmt.Printf("%s - %s - %s\n", device.IP, device.MAC, device.Manufacturer)
	}
}
