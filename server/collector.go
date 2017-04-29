package server

import (
	"fmt"
	"mycap/libs/agrqueries"
	"mycap/libs/agrqueries/countpertime"
	"mycap/libs/client"
	"time"
)

type Collector struct {
	Queries      agrqueries.QueriesAgregated
	CountPerTime countpertime.Counters

	server *Server
}

func (self *Collector) Collect() {
	self.CountPerTime.Init()
	for {
		for key, agent := range self.server.Agents.Items {
			if agent.LastCheckState && agent.LastCheckTime > time.Now().Unix()-3 {
				continue
			}
			cli := client.AgentClient{}
			cli.Host = agent.Service.Host
			cli.Port = agent.Service.Port

			queries, err := cli.GetQueries()

			agent.LastCheckState = err == nil && queries.Error.Code == 0
			agent.LastCheckTime = time.Now().Unix()

			self.server.Agents.Items[key] = agent

			if agent.LastCheckState {
				for _, query := range queries.Result {
					self.CountPerTime.Inc(query.Start.Unix())
					self.Queries.Add(agrqueries.CreateQuery(
						fmt.Sprintf("%s", query.SrcIP),
						fmt.Sprintf("%s:%d", query.DstIP, query.DstPort),
						query.Query,
						query.Duration,
					))
				}
				cli.ClearQueries()
			}
		}

		time.Sleep(time.Second)
	}
}
