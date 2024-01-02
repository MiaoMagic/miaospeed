//go:build core

package clash

import (
	"net/url"
	"strconv"

	"github.com/moshaoli688/clash/constant"
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
		DstIP:   nil,
		DstPort: constant.Port(portInt), // 将整数转换为 constant.Port 类型
	}
	return
}
