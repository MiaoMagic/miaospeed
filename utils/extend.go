package utils

import (
	"net"
	urllib "net/url"
)

func DomainPreheating(url string) net.IP {
	Domain, _ := urllib.Parse(url)
	Preheating, _ := net.LookupIP(Domain.Hostname())
	for _, PreheatingIp := range Preheating {
		if PreheatingIp == nil {
			DLogf("Task Poll | Task Domain Preheating Error, poll=%s Domain=%s Parsing failure", "DomainPreheating", Domain.Hostname())
		}
		return PreheatingIp
	}
	return nil
}
