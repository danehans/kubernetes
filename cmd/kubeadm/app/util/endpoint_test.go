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
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"testing"
)

func TestGenerateMasterEndpoint(t *testing.T) {
	var tests = []struct {
		protocol string
		address  string
		port     int32
		endpoint string
		expected bool
	}{
		{
			protocol: "https",
			address:  "1.2.3.4",
			port:     6443,
			endpoint: "https://1.2.3.4:6443",
			expected: true,
		},
		{
			protocol: "HTTP ",
			address:  "1.2.3.4",
			port:     6443,
			endpoint: "http://1.2.3.4:6443",
			expected: true,
		},
		{
			protocol: "https",
			address:  "1.2.3.4",
			port:     6443,
			endpoint: "https://[1.2.3.4]:6443",
			expected: false,
		},
		{
			protocol: "https",
			address:  "2001:db8::4",
			port:     6443,
			endpoint: "https://[2001:db8::4]:6443",
			expected: true,
		},
		{
			protocol: "https",
			address:  "2001:db8::4",
			port:     6443,
			endpoint: "https://2001:db8::4:6443",
			expected: false,
		},
		{
			protocol: "ftp",
			address:  "1.2.3.4",
			port:     3446,
			endpoint: "",
			expected: true,
		},
	}
	for _, rt := range tests {
		actual := kubeadmutil.GenerateMasterEndpoint(rt.protocol, rt.address, rt.port)
		//actual := "https" + "://" + "1.2.3.4" + ":" + "6443"
		if actual != rt.endpoint && rt.expected {
			t.Errorf(
				"failed GenerateMasterEndpoint:\n\texpected: %s\n\t  actual: %s",
				rt.endpoint,
				(actual),
			)
		}
	}
}
