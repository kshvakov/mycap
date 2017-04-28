package server

type Server struct {
	Service   Service
	Collector Collector
	Agents    Agents
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
