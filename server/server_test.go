package server

import (
	"testing"
)

func TestGetWhoisServer(t *testing.T) {
	tests := []struct {
		tld      string
		expected string
	}{
		{".com", "whois.verisign-grs.com"},
		{".in", "whois.registry.in"},
		{"1.1.1.1", "whois.iana.org"},
		{".com.mm", "whois-servers.net"},
	}

	for i, tt := range tests {
		server := GetWhoisServer(tt.tld)

		if server != tt.expected {
			t.Fatalf("tests[%d] - whois server wrong. tld=%q expected=%q, got=%q", i, tt.tld, tt.expected, server)
		}
	}
}
