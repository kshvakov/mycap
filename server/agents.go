package server

import "mycap/agent"

type Agents struct {
	Items agent.Agents `json:"items"`
}

func (self *Agents) GetAgents() agent.Agents {
	return self.Items
}
