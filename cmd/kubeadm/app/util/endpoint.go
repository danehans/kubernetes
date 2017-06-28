/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"fmt"
	"net"
	"strings"
)

// GenerateMasterEndpoint generates a properly formated Master URL.
func GenerateMasterEndpoint(protocol, address string, port int32) string {
	// Make sure protocol is valid and properly formatted
	proto := strings.TrimSpace(strings.ToLower(protocol))
	if proto != "http" && proto != "https" {
		return ""
	}

	// Generate the properly formatted API server endpoint based on IP address type.
	var masterEndpoint string
	masterIP := net.ParseIP(address)
	switch {
	case masterIP == nil:
		return ""
	case masterIP.To4() != nil:
		masterEndpoint = fmt.Sprintf("%s://%s:%d", proto, masterIP, port)
	case masterIP.To16() != nil:
		masterEndpoint = fmt.Sprintf("%s://[%s]:%d", proto, masterIP, port)
	}
	return masterEndpoint
}
