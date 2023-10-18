package structures

import (
	"go.bug.st/serial"
	"strings"
)

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

func StringToNoParity(src_string string) serial.Parity {
	src_string = strings.Trim(src_string, " \n\r")
	src_string = strings.ToLower(src_string)
	if src_string == "none" || src_string == "0" || src_string == "noparity" {
		return serial.NoParity
	}
	if src_string == "odd" || src_string == "1" || src_string == "oddparity" {
		return serial.OddParity
	}
	if src_string == "even" || src_string == "2" || src_string == "evenparity" {
		return serial.EvenParity
	}
	if src_string == "mark" || src_string == "3" || src_string == "markparity" {
		return serial.MarkParity
	}
	if src_string == "space" || src_string == "4" || src_string == "spaceparity" {
		return serial.SpaceParity
	}
	return serial.NoParity
}

func NoParityToString(src serial.Parity) string {
	if src == serial.NoParity {
		return "NoParity"
	}
	if src == serial.OddParity {
		return "OddParity"
	}
	if src == serial.EvenParity {
		return "EvenParity"
	}
	if src == serial.MarkParity {
		return "MarkParity"
	}
	if src == serial.SpaceParity {
		return "SpaceParity"
	}
	return ""
}

func StringToStopBits(src_string string) serial.StopBits {
	src_string = strings.Trim(src_string, " \n\r")
	src_string = strings.ToLower(src_string)
	if src_string == "one" || src_string == "1" || src_string == "onestopbit" {
		return serial.OneStopBit
	}
	if src_string == "two" || src_string == "2" || src_string == "twostopbit" {
		return serial.TwoStopBits
	}
	if src_string == "five" || src_string == "5" || src_string == "onepointfivestopbits" {
		return serial.OnePointFiveStopBits
	}
	return serial.OneStopBit
}

func StopBitsToString(src serial.StopBits) string {
	if src == serial.OneStopBit {
		return "OneStopBit"
	}
	if src == serial.TwoStopBits {
		return "TwoStopBits"
	}
	if src == serial.OnePointFiveStopBits {
		return "OnePointFiveStopBits"
	}
	return ""
}
