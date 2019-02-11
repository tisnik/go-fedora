package main

import (
	"fmt"
	"net"
)

func performLookup(address string) {
	ip, err := net.LookupIP(address)
	if err != nil {
		println("Lookup failure")
	} else {
		fmt.Printf("%v\n", ip)
	}
}

func main() {
	performLookup("localhost")
	performLookup("root.cz")
	performLookup("google.com")
}
