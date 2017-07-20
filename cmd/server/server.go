package main

import (
	"../../pkg/network"
	"../../pkg/node"
	"../../pkg/routes"
	"log"
	"net/http"
	"os"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("=> %s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		h.ServeHTTP(w, r) //call handler
	})
}

type Network network.Network
type Node node.Node

const (
	NAME    = "octopus-server"
	VERSION = "a0.1"

	DEFAULT_PORT = "2048"
)

func main() {
	log.Println(NAME, VERSION)

	network.InitLocal()

	this := node.CreateIdentity()

	if len(os.Args) >= 2 {
		this.Port = os.Args[1]
	} else {
		this.Port = DEFAULT_PORT
	}

	log.Printf("Will be known as %s over %s:%s\n", this.Name, this.Ip, this.Port)

	log.Println("Will initialize routes")
	//initialize the router
	r := routes.Init()

	log.Println("Starting server on", this.Port)
	log.Fatal(http.ListenAndServe(":"+this.Port, Log(r)))
}
