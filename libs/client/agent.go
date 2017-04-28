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
	Result map[string]libs.Query `json:"result"`
}

func (self *AgentClient) GetQueries() (GetQueriesFromAgentResponse, error) {
	res := GetQueriesFromAgentResponse{}
	err := self.Call("GetQueries", nil, &res)

	return res, err
}
