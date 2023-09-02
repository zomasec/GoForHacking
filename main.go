package main

import (
  "flag"
  "fmt"
  "net"
)

func main() {

  ipPtr := flag.String("ip", "", "IP address to lookup")
  domainPtr := flag.String("d", "", "Domain to lookup IP for")

  flag.Parse()

  if *ipPtr != "" {
    // Reverse lookup
    ip := net.ParseIP(*ipPtr)
    names, _ := net.LookupAddr(ip.String())

    fmt.Printf("Domains/subdomains for %s:\n", ip)
    for _, name := range names {
      fmt.Println(name)
    }

  } else if *domainPtr != "" {
    // Forward lookup   
    ips, _ := net.LookupIP(*domainPtr)

    fmt.Printf("IP addresses for %s:\n", *domainPtr)
    for _, ip := range ips {
      fmt.Println(ip)
    }

  } else {
    fmt.Println("Error: no IP or domain specified")
  }

}
