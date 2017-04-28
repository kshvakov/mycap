package server

import (
	"mycap/agent"
	"mycap/libs"
)

type Server struct {
	Service   Service
	Collector Collector
	Agents    Agents
}

func (self *Server) GetQueries() map[string]libs.Query {
	return self.Collector.queries
}

func (self *Server) GetAgents() map[string]agent.Agent {
	return self.Agents.GetAgents()
}

func (self *Server) StartJsonRpcServer() {
	self.Service.server = self
	self.Service.InitHandler()
	self.Service.ListenAndServe()
}

func (self *Server) StartCollector() {
	self.Collector.server = self
	self.Collector.Collect()
}
