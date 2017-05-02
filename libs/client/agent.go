package client

import (
	"mycap/libs"
	"mycap/libs/jsonrpc"
)

type AgentClient struct {
	jsonrpc.JsonRpcClient
}

type GetQueriesFromAgentResponse struct {
	jsonrpc.JsonRpcResponse
	Result libs.Queries `json:"result"`
}

type ClearQueriesFromAgentResponse struct {
	jsonrpc.JsonRpcResponse
	Result bool `json:"result"`
}

func (self *AgentClient) GetQueries() (GetQueriesFromAgentResponse, error) {
	res := GetQueriesFromAgentResponse{}
	err := self.Call("GetQueries", nil, &res)

	return res, err
}

func (self *AgentClient) ClearQueries() (ClearQueriesFromAgentResponse, error) {
	res := ClearQueriesFromAgentResponse{}
	err := self.Call("ClearQueries", nil, &res)

	return res, err
}
