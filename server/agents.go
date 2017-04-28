package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mycap/agent"
)

type Agents struct {
	agents map[string]agent.Agent
}

func (self *Agents) CreateFromJsonFile(agentsFile string) {
	self.agents = make(map[string]agent.Agent)
	content, err := ioutil.ReadFile(agentsFile)

	if err != nil {
		log.Fatal(err)
	}

	agents := []agent.Agent{}

	if err := json.Unmarshal(content, &agents); err != nil {
		log.Fatal(err)
	}

	for _, agent := range agents {
		if len(agent.Id) == 0 {
			log.Fatal("Uninitialized agent ID")
		}
		self.agents[agent.Id] = agent
	}
}

func (self *Agents) GetAgents() map[string]agent.Agent {
	return self.agents
}

func (self *Agents) SetAgent(agent agent.Agent) {
	self.agents[agent.Id] = agent
}
