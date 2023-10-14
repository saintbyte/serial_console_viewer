package structures

import (
	"flag"
	"fmt"
	"os"
)

func Help() {
	fmt.Println("serial_console_viewer - program to simple read from com ports")
	fmt.Println("Usage: serial_console_viewer [arguments]")
	fmt.Printf("Help:\n")
	flag.PrintDefaults()
	os.Exit(1)
}
