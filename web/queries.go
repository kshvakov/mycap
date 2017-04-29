package web

import (
	"mycap/libs/agrqueries"
	"mycap/libs/agrqueries/countpertime"
	"mycap/libs/client"
	"time"
)

type QueriesCollector struct {
	server       *Server
	queries      agrqueries.QueriesAgregated
	countPerTime countpertime.Counters
}

func (self *QueriesCollector) Collect() {
	cli := client.ServerClient{}
	cli.Host = self.server.HeadServerHost
	cli.Port = self.server.HeadServerPort

	for {
		func() {
			response, err := cli.GetQueries()

			if err == nil && response.Error.Code == 0 {
				self.queries = response.Result
			}
		}()

		func() {
			response, err := cli.GetCountPerTime()

			if err == nil && response.Error.Code == 0 {
				self.countPerTime = response.Result
			}
		}()

		time.Sleep(time.Second)
	}
}
