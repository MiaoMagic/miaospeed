package interfaces

import "github.com/moshaoli688/miaospeed/utils/structs"

type ProxyType string

const (
	Shadowsocks  ProxyType = "Shadowsocks"
	ShadowsocksR ProxyType = "ShadowsocksR"
	Snell        ProxyType = "Snell"
	Socks5       ProxyType = "Socks5"
	Http         ProxyType = "Http"
	Vmess        ProxyType = "Vmess"
	Trojan       ProxyType = "Trojan"

	Vless     ProxyType = "Vless"
	Hysteria  ProxyType = "Hysteria"
	Hysteria2 ProxyType = "Hysteria2"

	Tuic ProxyType = "Tuic"

	ProxyInvalid ProxyType = "Invalid"
)

var AllProxyTypes = []ProxyType{
	Shadowsocks, ShadowsocksR, Snell, Socks5, Http, Vmess, Trojan,
	Vless, Hysteria, Hysteria2, Tuic,
}

func Valid(proxyType ProxyType) bool {
	return structs.Contains(AllProxyTypes, proxyType)
}

func Parse(proxyType string) ProxyType {
	pType := ProxyType(proxyType)
	if Valid(pType) {
		return pType
	}
	return ProxyInvalid
}

type ProxyInfo struct {
	Name    string
	Address string
	Type    ProxyType
}

func (pi *ProxyInfo) Map() map[string]string {
	return map[string]string{
		"Name":    pi.Name,
		"Address": pi.Address,
		"Type":    string(pi.Type),
	}
}
