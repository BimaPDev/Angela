package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello world")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%s: %s\n", name, h)
		}
	}
}

func getLocalIPv4() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		// Skip down or loopback interfaces
		if iface.Flags&(net.FlagUp|net.FlagLoopback) != net.FlagUp {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && ipnet.IP.To4() != nil && !ipnet.IP.IsLoopback() {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no IPv4 address found")
}

func main() {
	ip, err := getLocalIPv4()
	if err != nil {
		log.Printf("could not detect LAN IP: %v", err)
	}

	port := "8090"
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)

	if ip != "" {
		log.Printf("listening on http://%s:%s", ip, port)
	}
	log.Printf("also on http://localhost:%s", port)

	addr := ":" + port // Listen on all interfaces
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
