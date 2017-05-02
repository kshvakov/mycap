package agent

type Agents []Agent

type Agent struct {
	Id string `json:"id"`

	Collector Collector `json:"collector"`
	Service   Service   `json:"service"`

	LastCheckTime  int64 `json:"LastCheckTime"`
	LastCheckState bool  `json:"LastCheckState"`
}

func (self *Agent) StartJsonRpcServer() {
	self.Service.collector = &self.Collector
	self.Service.InitHandler()
	self.Service.ListenAndServe()
}

func (self *Agent) StartCollector() {
	self.Collector.Collect()
}
