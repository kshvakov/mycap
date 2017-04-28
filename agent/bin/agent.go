package main

import (
	"flag"
	"fmt"
	"mycap/agent"
	"time"
)

var (
	Device = flag.String("device", "lo", "")

	BPFFilter = flag.String("bpf_filter", "tcp and port 3306", "")
	TXTFilter = flag.String("txt_filter", "", "")

	MaxQueryLen = flag.Int("max_query_len", 9192, "")

	ServiceHost = flag.String("service_host", "localhost", "")
	ServicePort = flag.Int("service_port", 9500, "")
)

func main() {

	flag.Parse()

	a := &agent.Agent{
		Id: fmt.Sprintf("%s:%d:%s", *ServiceHost, *ServicePort, *ServiceHost),
		Collector: agent.Collector{
			Device:      *Device,
			BPFFilter:   *BPFFilter,
			TXTFilter:   *TXTFilter,
			MaxQueryLen: *MaxQueryLen,
		},
	}

	a.Service.Host = *ServiceHost
	a.Service.Port = *ServicePort

	go a.StartJsonRpcServer()
	go a.Collector.Collect()

	// @TODO Here we'll be listen signals later
	for {
		time.Sleep(time.Second)
	}
}
