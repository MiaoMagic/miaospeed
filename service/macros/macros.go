package macros

import (
	"github.com/moshaoli688/miaospeed/interfaces"
	"github.com/moshaoli688/miaospeed/utils/structs"

	"github.com/moshaoli688/miaospeed/service/macros/geo"
	"github.com/moshaoli688/miaospeed/service/macros/ping"
	"github.com/moshaoli688/miaospeed/service/macros/script"
	"github.com/moshaoli688/miaospeed/service/macros/speed"
	"github.com/moshaoli688/miaospeed/service/macros/udp"

	"github.com/moshaoli688/miaospeed/service/macros/invalid"
)

var registeredList = map[interfaces.SlaveRequestMacroType]func() interfaces.SlaveRequestMacro{
	interfaces.MacroSpeed: func() interfaces.SlaveRequestMacro {
		return &speed.Speed{}
	},
	interfaces.MacroPing: func() interfaces.SlaveRequestMacro {
		return &ping.Ping{}
	},
	interfaces.MacroUDP: func() interfaces.SlaveRequestMacro {
		return &udp.Udp{}
	},
	interfaces.MacroGeo: func() interfaces.SlaveRequestMacro {
		return &geo.Geo{}
	},
	interfaces.MacroScript: func() interfaces.SlaveRequestMacro {
		return &script.Script{}
	},
}

func Find(macroType interfaces.SlaveRequestMacroType) interfaces.SlaveRequestMacro {
	if newFn, ok := registeredList[macroType]; ok && newFn != nil {
		return newFn()
	}

	return &invalid.Invalid{}
}

func FindBatch(macroTypes []interfaces.SlaveRequestMacroType) []interfaces.SlaveRequestMacro {
	return structs.Map(macroTypes, func(m interfaces.SlaveRequestMacroType) interfaces.SlaveRequestMacro {
		return Find(m)
	})
}
