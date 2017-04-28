package agent

import (
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
		default:
			return nil, nil
		}
	})
}
