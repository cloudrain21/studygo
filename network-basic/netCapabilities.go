package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Printf("Interface Flags: %v\n", i.Flags.String())
		fmt.Printf("Interface MTU: %v\n", i.MTU)
		fmt.Printf("Interface Hardware Address: %v\n", i.HardwareAddr)
		fmt.Println()
	}
}
