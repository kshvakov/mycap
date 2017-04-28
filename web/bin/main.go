package main

import (
	"flag"
	"html/template"
	"mycap/web"
	"time"
)

var (
	tpl *template.Template

	webPath = flag.String("web_path", "localhost", "")

	serverHost = flag.String("server_host", "localhost", "")
	serverPort = flag.Int("server_port", 9600, "")

	serviceHost = flag.String("service_host", "localhost", "")
	servicePort = flag.Int("service_port", 9700, "")
)

func main() {
	flag.Parse()

	s := web.Server{
		Host: *serviceHost,
		Port: *servicePort,

		HeadServerHost: *serverHost,
		HeadServerPort: *serverPort,
	}

	s.PathTemplates = *webPath + "/templates/"
	s.PathStatic = *webPath + "/static/"
	s.InitTemplates()

	go s.StartAgentsCollector()
	go s.StartQueriesCollector()
	go s.StartWebServer()

	for {
		time.Sleep(time.Second)
	}
}
