package structures

import "go.bug.st/serial"

type CommandLineActions struct {
	ListAction bool
	ReadAction bool
}

type PortConfig struct {
	BaudRate int
	DataBits int
	Parity   int
	StopBits int
}

func NewCommandLineActions() CommandLineActions {
	return CommandLineActions{ListAction: false, ReadAction: false}
}

func NewPortConfig() PortConfig {
	return PortConfig{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   int(serial.NoParity),
		StopBits: int(serial.OneStopBit),
	}
}
