package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	ips := ""
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips += ipnet.IP.String() + ", "
			}
		}
	}
	ips = ips[:len(ips)-2]
	fmt.Fprintf(w,
		"<div style=\"color:black\" align=\"center\">\n"+
			"<h1>DemoApp v1</h1>\n"+
			"<h2>Hostname: %s</h2>\n"+
			"<h2>IP(s): %s</h2>\n"+
			"</div>",
		hostname, ips)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
