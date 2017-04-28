package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mycap/agent"
)

type Agents struct {
	Items agent.Agents
}

func (self *Agents) CreateFromJsonFile(agentsFile string) {
	self.Items = agent.Agents{}
	content, err := ioutil.ReadFile(agentsFile)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(content, &self.Items); err != nil {
		log.Fatal(err)
	}
}

func (self *Agents) GetAgents() agent.Agents {
	return self.Items
}
