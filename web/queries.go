package web

import (
	"mycap/libs/agrqueries"
	"mycap/libs/client"
	"time"
)

type QueriesCollector struct {
	server  *Server
	queries agrqueries.QueriesAgregated
}

func (self *QueriesCollector) Collect() {
	cli := client.ServerClient{}
	cli.Host = self.server.HeadServerHost
	cli.Port = self.server.HeadServerPort

	for {
		queries, err := cli.GetQueries()

		if err == nil && queries.Error.Code == 0 {
			self.queries = queries.Result
		}

		time.Sleep(time.Second)
	}
}
