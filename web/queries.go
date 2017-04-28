package web

import (
	"fmt"
	"mycap/libs/agrqueries"
	"mycap/libs/client"
	"time"
)

type QueriesCollector struct {
	server  *Server
	queries agrqueries.Queries
}

func (self *QueriesCollector) Collect() {
	cli := client.ServerClient{}
	cli.Host = self.server.HeadServerHost
	cli.Port = self.server.HeadServerPort

	self.queries = agrqueries.Queries{}

	for {
		queries, err := cli.GetQueries()

		if err == nil && queries.Error.Code == 0 {
			for _, query := range queries.Result {

				self.queries.Add(agrqueries.CreateQuery(
					fmt.Sprintf("%s", query.SrcIP),
					fmt.Sprintf("%s:%d", query.DstIP, query.DstPort),
					query.Query,
					query.Duration,
				))
			}
		}

		time.Sleep(1 * time.Second)
	}
}
