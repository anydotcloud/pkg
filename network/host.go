package network

import (
	"fmt"
	"net"
	"strings"
)

// GetLocalIPv4Address returns the IPv4 address of the local host. It makes a
// call to Google's DNS service at 8.8.8.8 from the local host to determine the
// local IP address. If no outbound network connectivity is allowed, returns
// "0.0.0.0"
func GetLocalIPv4Address() string {
	c, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Warning: unable to make a UDP request to 8.8.8.8:80 " +
			"to determine local host address. Using 0.0.0.0")
		return "0.0.0.0"
	}
	defer c.Close()
	addr := c.LocalAddr().String()
	return addr[:strings.LastIndex(addr, ":")]
}
