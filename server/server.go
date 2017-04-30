package server

type Server struct {
	Service   Service   `json:"service"`
	Collector Collector `json:"collector"`
	Agents    Agents    `json:"agents"`
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
