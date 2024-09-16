package clash

import (
	"github.com/Dreamacro/clash/adapter"
	"github.com/Dreamacro/clash/constant"
	"github.com/moshaoli688/miaospeed/interfaces"
	"github.com/moshaoli688/miaospeed/utils"
	"gopkg.in/yaml.v2"
	"strings"
)

func parseProxy(proxyName, proxyPayload string) constant.Proxy {
	var payload map[string]any
	yaml.Unmarshal([]byte(proxyPayload), &payload)
	proxy, err := adapter.ParseProxy(payload)

	if err != nil {
		utils.DLogf("Vendor Parser | Parse clash profile error, error=%v", err.Error())
	}

	return proxy
}

func extractFirstProxy(proxyName, proxyPayload string) constant.Proxy {
	if processProxyPayload(proxyPayload) {
		return nil
	}
	proxy := parseProxy(proxyName, proxyPayload)

	if proxy != nil && interfaces.Parse(proxy.Type().String()) != interfaces.ProxyInvalid {
		return proxy
	}

	return nil
}
func processProxyPayload(proxyPayload string) bool {
	if !strings.Contains(proxyPayload, "type: vless") {
		return false
	}
	var data map[string]interface{}
	_ = yaml.Unmarshal([]byte(proxyPayload), &data)
	if flow, ok := data["flow"].(string); ok {
		legacyFlows := []string{"xtls-rprx-direct", "xtls-rprx-origin", "xtls-rprx-splice"}
		for _, legacyFlow := range legacyFlows {
			if flow == legacyFlow {
				utils.DInfof("Detected legacy flow: %s. This protocol is deprecated.", flow)
				return true
			}
		}
	}
	return false
}
