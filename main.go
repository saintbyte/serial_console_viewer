package main

import (
	"flag"
	"fmt"
	internal "github.com/saintbyte/serial_console_viewer/internal"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
	"log"
	"os"
	"os/signal"
)

func listPorts() {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	fmt.Printf("Found port:\n")
	for _, port := range ports {
		fmt.Printf("\t %v\n", port.Name)
		if port.IsUSB {
			fmt.Printf("\t\tUSB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("\t\tUSB serial %s\n", port.SerialNumber)
		}
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
	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
		}
		fmt.Printf("%v", string(buff[:n]))
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println("Receive:", sig.String())
			os.Exit(0)
		}
	}()

	if actionArgs.ListAction {
		listPorts()
	}
	if actionArgs.ReadAction {
		readPort(portName, portConfig)
	}
	if !actionArgs.ListAction && !actionArgs.ReadAction {
		internal.Help()
	}
}
