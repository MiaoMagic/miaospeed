package rttping

import (
	"github.com/moshaoli688/miaospeed/interfaces"
	"github.com/moshaoli688/miaospeed/service/macros/ping"
)

type RTTPing struct {
	interfaces.RTTPingDS
}

func (m *RTTPing) Type() interfaces.SlaveRequestMatrixType {
	return interfaces.MatrixRTTPing
}

func (m *RTTPing) MacroJob() interfaces.SlaveRequestMacroType {
	return interfaces.MacroPing
}

func (m *RTTPing) Extract(entry interfaces.SlaveRequestMatrixEntry, macro interfaces.SlaveRequestMacro) {
	if mac, ok := macro.(*ping.Ping); ok {
		m.Value = mac.RTT
	}
}
