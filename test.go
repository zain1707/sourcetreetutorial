package main

import (
	"fmt"
	"os"
)

func main() {

	hostname, _ := os.Hostname()
	newhostname := "https://" + hostname
	path := "/google/authorize"
	port := ":5000"
	fmt.Println(newhostname + path + port)
	//ip := "0.0.0.0"
	// port := ":5000"
	// publickey := "/etc/letsencrypt/live/api1.dealbloom.com/fullchain.pem"
	// privatekey := "/etc/letsencrypt/live/api1.dealbloom.com/privkey.pem"
	// e := echo.New()

	// //e.Logger.Fatal(e.Start(hostname + port))
	// //fmt.Println(hostname + port)
	// e.Logger.Fatal(e.StartTLS(ip+port, publickey, privatekey))

}
