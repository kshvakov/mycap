package agent

import (
	"mycap/libs"
	"mycap/libs/jsonrpc"
)

type Service struct {
	jsonrpc.JsonRpcServer
	collector *Collector
}

func (self *Service) InitHandler() {
	self.SetHandler(func(request jsonrpc.JsonRpcRequest) (interface{}, error) {
		switch request.Method {
		case "GetQueries":
			return self.collector.queries, nil
		case "ClearQueries":
			self.collector.queries = libs.Queries{}
			return true, nil
		default:
			return nil, nil
		}
	})
}
