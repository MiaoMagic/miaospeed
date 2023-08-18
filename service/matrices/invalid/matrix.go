package invalid

import "github.com/moshaoli688/miaospeed/interfaces"

type Invalid struct {
	interfaces.InvalidDS
}

func (m *Invalid) Type() interfaces.SlaveRequestMatrixType {
	return interfaces.MatrixInvalid
}

func (m *Invalid) MacroJob() interfaces.SlaveRequestMacroType {
	return interfaces.MacroInvalid
}

func (m *Invalid) Extract(entry interfaces.SlaveRequestMatrixEntry, macro interfaces.SlaveRequestMacro) {
}
