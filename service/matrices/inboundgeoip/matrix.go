package inboundgeoip

import (
	"github.com/moshaoli688/miaospeed/interfaces"
	"github.com/moshaoli688/miaospeed/service/macros/geo"
)

type InboundGeoIP struct {
	interfaces.InboundGeoIPDS
}

func (m *InboundGeoIP) Type() interfaces.SlaveRequestMatrixType {
	return interfaces.MatrixInboundGeoIP
}

func (m *InboundGeoIP) MacroJob() interfaces.SlaveRequestMacroType {
	return interfaces.MacroGeo
}

func (m *InboundGeoIP) Extract(entry interfaces.SlaveRequestMatrixEntry, macro interfaces.SlaveRequestMacro) {
	if mac, ok := macro.(*geo.Geo); ok {
		m.MultiStacks = mac.InStacks
	}
}
