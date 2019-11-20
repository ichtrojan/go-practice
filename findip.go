package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//find ip address using ifify.org
func main() {
	url := "https://api.ipify.org?format=text"
	fmt.Printf("Please wait while I fetch your IP address from  ipify\n")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("My IP is:%s\n", ip)
}
