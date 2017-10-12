package network

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// Device ...
type Device struct {
	IP           string
	MAC          string
	Manufacturer string
}

// ScanOnIPv4WithCIDR scans for devices on the given interface
func ScanOnIPv4WithCIDR(iface, cidrIP string) ([]Device, error) {
	ipParts := strings.Split(cidrIP, "/")
	if len(ipParts) != 2 {
		return nil, errors.New("Invalid CIDR IP - " + cidrIP)
	}
	ip := ipParts[0]
	mask := ipParts[1]

	// nmap args
	argTarget := ip[:strings.LastIndex(ip, ".")] + ".0/" + mask

	// TODO: check if it's running with su privileges
	cmd := exec.Command("nmap", "-oN", "-", "-sP", argTarget, "-T", "insane", "--exclude", ip)

	//cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	l := len(lines)

	if l < 5 {
		return nil, errors.New("nmap invalid response")
	}

	var devices []Device
	for i := 1; i < l-2; i += 3 {
		devices = append(devices, Device{
			IP:           lines[i][21:],
			MAC:          lines[i+2][13:30],
			Manufacturer: lines[i+2][32 : len(lines[i+2])-1],
		})
	}

	return devices, nil
}
