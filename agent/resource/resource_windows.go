package resource

import (
	"bytes"
	"os"
)

func init() {
	if s, err := os.ReadFile("/sys/class/dmi/id/product_serial"); err == nil {
		hostSerial = string(bytes.TrimSpace(s))
	}
	if s, err := os.ReadFile("/sys/class/dmi/id/product_uuid"); err == nil {
		hostID = string(bytes.TrimSpace(s))
	}
	if s, err := os.ReadFile("/sys/class/dmi/id/product_name"); err == nil {
		hostModel = string(bytes.TrimSpace(s))
	}
	if s, err := os.ReadFile("/sys/class/dmi/id/sys_vendor"); err == nil {
		hostVendor = string(bytes.TrimSpace(s))
	}
}

func GetDNS() string {
	return "windows-NAN"
}

func GetGateway() string {
	return "windows-NAN"
}
