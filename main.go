package main

import (
	"flag"
	"fmt"
	internal "github.com/saintbyte/serial_console_viewer/internal"
	"go.bug.st/serial"
	"log"
)

func listPorts() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	fmt.Printf("Found port:\n")
	for _, port := range ports {
		log.Println("\t %v\n", port)
	}
}

func readPort(portName string, config internal.PortConfig) {
	mode := &serial.Mode{
		BaudRate: config.BaudRate,
		DataBits: config.DataBits,
		Parity:   serial.Parity(config.Parity),
		StopBits: serial.StopBits(config.StopBits),
	}

	log.Println("Open port:", portName)
	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatal(err)
	}
	err = port.SetMode(mode)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	actionArgs := internal.NewCommandLineActions()
	portConfig := internal.NewPortConfig()
	var portName string
	flag.BoolVar(&actionArgs.ListAction, "list", false, "List available serial/com ports")
	flag.BoolVar(&actionArgs.ReadAction, "read", false, "Read data from serial/com ports")
	flag.StringVar(&portName, "port", "/dev/ttyUSB0", "Port to read")
	flag.IntVar(&portConfig.BaudRate, "baunrate", 9600, "BaudRate/Speed of serial port")
	flag.IntVar(&portConfig.DataBits, "databits", 8, "Data bits: 5,6,7,8")
	flag.IntVar(&portConfig.Parity, "Parity", 0, "Parity: 0,1,2")
	flag.IntVar(&portConfig.StopBits, "stopbits", 1, "Stop bits: 0,1,2")
	flag.Parse()

	if actionArgs.ListAction {
		listPorts()
	}
	if actionArgs.ReadAction {
		readPort(portName, portConfig)
	}
	if !actionArgs.ListAction || !actionArgs.ReadAction {
		internal.Help()
	}
}
