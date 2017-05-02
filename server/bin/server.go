package main

import (
	"flag"
	"log"
	"mycap/libs/config"
	"mycap/server"
	"time"
)

var configFile = flag.String("config", "./../etc/server.json", "")

func main() {
	flag.Parse()

	srv := &server.Server{}

	if err := config.ReadConfig(*configFile, &srv); err != nil {
		log.Fatal(err)
	}

	go srv.StartJsonRpcServer()
	go srv.StartCollector()

	for {
		time.Sleep(1 * time.Second)
	}
}
