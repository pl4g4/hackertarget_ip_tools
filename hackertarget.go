package main

import (
	"flag"
	"fmt"
	"hackertarget_ip_tools/hackertargetapi"
)

func main() {
	ipAddress := flag.String("ip", "127.0.0.1", "The ip address. Could be a IP or Domain")
	hackerTargetOption := flag.String("option", "whois", "MTR Traceroute:mtr \nPing:nping DNS \nLookup:dnslookup \nReverse DNS Lookup:reversedns \nWhois Lookup:whois \nGeoIP:geoip \nReverse IP:reverseiplookup \nHTTP Headers: httpheaders \nPage Link:pagelinks")
	flag.Parse()
	response := hackertargetapi.GetHackerTargetResponse(*hackerTargetOption, *ipAddress)
	fmt.Println(response)

}
