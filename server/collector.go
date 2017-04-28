package server

import (
	"mycap/libs"
	"mycap/libs/client"
	"time"
)

type Collector struct {
	queries map[string]libs.Query
	server  *Server
}

func (self *Collector) Collect() {
	self.queries = make(map[string]libs.Query)

	for {
		for _, agent := range self.server.GetAgents() {
			if agent.LastCheckState && agent.LastCheckTime > time.Now().Unix()-3 {
				continue
			}
			cli := client.AgentClient{}
			cli.Host = agent.Service.Host
			cli.Port = agent.Service.Port

			queries, err := cli.GetQueries()

			agent.LastCheckState = err == nil && queries.Error.Code == 0
			agent.LastCheckTime = time.Now().Unix()

			self.server.Agents.SetAgent(agent)

			for _, query := range queries.Result {
				self.queries[query.ID] = query
			}
		}

		time.Sleep(time.Second)
	}
}
