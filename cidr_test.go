package main

import (
	"fmt"
	"testing"
)

func TestCidrRange(t *testing.T) {
	tests := []struct {
		cidr        string
		expected    string
		shouldError bool
	}{
		{"192.168.1.1/32", "192.168.1.1 - 192.168.1.1", false},
		{"192.168.1.1/24", "192.168.1.0 - 192.168.1.255", false},
		{"192.168.1.1/18", "192.168.0.0 - 192.168.63.255", false},
		{"192.168.1.1/16", "192.168.0.0 - 192.168.255.255", false},
		{"192.168.1.1/8", "192.0.0.0 - 192.255.255.255", false},
		{"192.168.1.1/0", "0.0.0.0 - 255.255.255.255", false},
		{"192.168.1.1/-1", "", true},
		{"192.168.1.1/33", "", true},
		{"192.168.1.1/256", "", true},
		{"192.168.256.1/256", "", true},
		{"1920.168.0.1/256", "", true},
		{"192.-1.0.1/256", "", true},
	}

	for _, c := range tests {
		start, end, err := getIPRange(c.cidr)
		if c.shouldError {
			if err == nil {
				t.Errorf("%v: Expected error but did not get one", c.cidr)
			}

			continue
		}
		received := fmt.Sprintf("%v - %v", start, end)
		if received != c.expected {
			t.Errorf("%v: Expected %v but received %v", c.cidr, c.expected, received)
		}
	}
}
