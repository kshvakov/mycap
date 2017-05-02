package server

import (
	"mycap/libs/jsonrpc"
)

type Service struct {
	jsonrpc.JsonRpcServer
	server *Server
}

func (self *Service) InitHandler() {
	self.SetHandler(func(request jsonrpc.JsonRpcRequest) (interface{}, error) {
		switch request.Method {
		case "GetQueries":
			return self.server.Collector.Queries, nil
		case "GetCountPerTime":
			return self.server.Collector.CountPerTime, nil
		case "GetAgents":
			return self.server.Agents.Items, nil
		default:
			return nil, nil
		}
	})
}
