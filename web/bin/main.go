package main

import (
	"flag"
	"log"
	"mycap/libs/config"
	"mycap/web"
	"time"
)

var configFile = flag.String("config", "./../etc/web.json", "")

func main() {
	flag.Parse()

	s := web.Server{}

	if err := config.ReadConfig(*configFile, &s); err != nil {
		log.Fatal(err)
	}

	s.InitTemplates()

	go s.StartAgentsCollector()
	go s.StartQueriesCollector()
	go s.StartWebServer()

	for {
		time.Sleep(time.Second)
	}
}
