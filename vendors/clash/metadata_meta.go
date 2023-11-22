//go:build meta

package clash

import (
	"github.com/Dreamacro/clash/constant"
	"net/netip"
	"net/url"
	"strconv"
)

func urlToMetadata(rawURL string, network constant.NetWork) (addr constant.Metadata, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}

	portStr := u.Port()
	if portStr == "" {
		switch u.Scheme {
		case "https":
			portStr = "443"
		case "http":
			portStr = "80"
		default:
			return
		}
	}
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		return
	}
	addr = constant.Metadata{
		NetWork: network,
		Host:    u.Hostname(),
		DstIP:   netip.Addr{},
		DstPort: uint16(portInt),
	}
	return
}
