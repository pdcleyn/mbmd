package rs485

import . "github.com/volkszaehler/mbmd/meters"

func init() {
	Register(NewSDMProducer)
}

const (
	METERTYPE_SDM = "SDM"
)

type SDMProducer struct {
	Opcodes
}

func NewSDMProducer() Producer {
	/**
	 * Opcodes as defined by Eastron.
	 * See http://bg-etech.de/download/manual/SDM630Register.pdf
	 * Please note that this is the superset of all SDM devices -
	 * some opcodes might not work on some devicep.
	 */
	ops := Opcodes{
		VoltageL1: 0x0000,
		VoltageL2: 0x0002,
		VoltageL3: 0x0004,
		CurrentL1: 0x0006,
		CurrentL2: 0x0008,
		CurrentL3: 0x000A,
		Power:     0x0034,
		PowerL1:   0x000C,
		PowerL2:   0x000E,
		PowerL3:   0x0010,
		ImportL1:  0x015a,
		ImportL2:  0x015c,
		ImportL3:  0x015e,
		Import:    0x0048,
		ExportL1:  0x0160,
		ExportL2:  0x0162,
		ExportL3:  0x0164,
		Export:    0x004a,
		CosphiL1:  0x001e,
		CosphiL2:  0x0020,
		CosphiL3:  0x0022,
		//L1THDCurrent: 0x00F0, // current
		//L2THDCurrent: 0x00F2, // current
		//L3THDCurrent: 0x00F4, // current
		//AvgTHDCurrent: 0x00Fa, // current
		THDL1:     0x00ea, // voltage
		THDL2:     0x00ec, // voltage
		THDL3:     0x00ee, // voltage
		THD:       0x00F8, // voltage
		Frequency: 0x0046,
	}
	return &SDMProducer{Opcodes: ops}
}

func (p *SDMProducer) Type() string {
	return METERTYPE_SDM
}

func (p *SDMProducer) Description() string {
	return "Eastron SDM meters"
}

func (p *SDMProducer) snip(iec Measurement) Operation {
	operation := Operation{
		FuncCode:  ReadInputReg,
		OpCode:    p.Opcode(iec),
		ReadLen:   2,
		IEC61850:  iec,
		Transform: RTUIeee754ToFloat64,
	}
	return operation
}

func (p *SDMProducer) Probe() Operation {
	return p.snip(VoltageL1)
}

func (p *SDMProducer) Produce() (res []Operation) {
	for op := range p.Opcodes {
		res = append(res, p.snip(op))
	}

	return res
}
