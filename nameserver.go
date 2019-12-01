package main

import (
	"fmt"
	"net"
)

//look up name servers of any domain
func main() {
	nameserver, _ := net.LookupNS("google.com")
	for _, ns := range nameserver {
		fmt.Println(ns)
	}
}
