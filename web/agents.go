package web

import (
	"mycap/agent"
	"mycap/libs/client"
	"time"
)

type AgentsCollector struct {
	server *Server
	agents map[string]agent.Agent
}

func (self *AgentsCollector) Collect() {
	cli := client.ServerClient{}
	cli.Host = self.server.HeadServerHost
	cli.Port = self.server.HeadServerPort

	for {
		agents, err := cli.GetAgents()

		if err == nil && agents.Error.Code == 0 {
			self.agents = agents.Result
		}
		time.Sleep(1 * time.Second)
	}
}
