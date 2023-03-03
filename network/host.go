// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

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
