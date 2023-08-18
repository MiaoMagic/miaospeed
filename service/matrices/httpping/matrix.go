package httpping

import (
	"github.com/moshaoli688/miaospeed/interfaces"
	"github.com/moshaoli688/miaospeed/service/macros/ping"
)

type HTTPPing struct {
	interfaces.HTTPPingDS
}

func (m *HTTPPing) Type() interfaces.SlaveRequestMatrixType {
	return interfaces.MatrixHTTPPing
}

func (m *HTTPPing) MacroJob() interfaces.SlaveRequestMacroType {
	return interfaces.MacroPing
}

func (m *HTTPPing) Extract(entry interfaces.SlaveRequestMatrixEntry, macro interfaces.SlaveRequestMacro) {
	if mac, ok := macro.(*ping.Ping); ok {
		m.Value = mac.Request
	}
}
