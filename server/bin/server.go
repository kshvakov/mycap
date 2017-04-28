package main

import (
	"flag"
	"mycap/server"
	"time"
)

var (
	nodesFile   = flag.String("nodes_file", "", "")
	serviceHost = flag.String("service_host", "localhost", "")
	servicePort = flag.Int("service_port", 9600, "")
)

func main() {
	flag.Parse()

	srv := &server.Server{}
	srv.Service.Host = *serviceHost
	srv.Service.Port = *servicePort
	srv.Agents.CreateFromJsonFile(*nodesFile)

	go srv.StartJsonRpcServer()
	go srv.StartCollector()

	for {
		time.Sleep(1 * time.Second)
	}
}
