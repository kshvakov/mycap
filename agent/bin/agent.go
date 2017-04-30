package main

import (
	"flag"
	"log"
	"mycap/agent"
	"mycap/libs/config"
	"time"
)

var configFile = flag.String("config", "./../etc/agent.json", "")

func main() {

	flag.Parse()

	a := &agent.Agent{}

	if err := config.ReadConfig(*configFile, &a); err != nil {
		log.Fatal(err)
	}

	go a.StartJsonRpcServer()
	go a.Collector.Collect()

	for {
		time.Sleep(time.Second)
	}
}
