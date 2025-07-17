package main

import (
	_ "embed"
	"flag"
	"log"
	"net"

	"github.com/burgerdev/evil-registry/registry"
)

var addr = flag.String("addr", ":80", "address to serve on")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("Could not listen on %s: %v", *addr, err)
	}
	log.Printf("serving a registry on %s ...", *addr)
	registry.Run(listener)
}
