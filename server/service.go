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
			return self.server.GetQueries(), nil
		case "GetAgents":
			return self.server.GetAgents(), nil
		default:
			return nil, nil
		}
	})
}
