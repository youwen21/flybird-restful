package main

import (
	"fmt"
	"gofly/conf"
	"gofly/conf/flag_vars"
	"gofly/server"
	"log"
	"net"
	"net/http"
)

func main() {
	fmt.Println("starting")

	listener, err := net.Listen("tcp", conf.GetAddress())
	if err != nil {
		log.Fatal("init listener failure:", err)
	}
	log.Println("listen at:", conf.GetAddress())
	log.Println("local open:", fmt.Sprintf("http://127.0.0.1:%v", flag_vars.GetPort()))

	err = http.Serve(listener, server.Mux)
	if err != nil {
		log.Fatal("http.Serve failure:", err)
	}

}
