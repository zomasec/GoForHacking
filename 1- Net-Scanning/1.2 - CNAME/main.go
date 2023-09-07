package main

import (
	"flag"
	"fmt"
	"net"
)

//Function to output the CNAME Record for a host name
func CNAME(host string) {
	CNAME, err := net.LookupCNAME(host)
  // Handle the error if exist
	if err != nil {
		fmt.Println(err)
		return // End CNAME function if err is exist 
    }
  // Return the CNAME Record if no error exist
  fmt.Println("CNAME Record for", host, "is :", CNAME)
	}
	
//The main function
func main(){
  // Taking the domain from the user as argument with -d flag 
  domain := flag.String("d", "", "Domain name or host name ")
  flag.Parse()
  CNAME(*domain)
  
 
}
