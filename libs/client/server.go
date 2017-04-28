package client

import (
	"mycap/agent"
	"mycap/libs/agrqueries"
	"mycap/libs/jsonrpc"
)

type ServerClient struct {
	jsonrpc.JsonRpcClient
}

type GetQueriesFromServerResponse struct {
	jsonrpc.JsonRpcResponse
	Result agrqueries.QueriesAgregated `json:"result"`
}

type GetAgentsFromServerResponse struct {
	jsonrpc.JsonRpcResponse
	Result agent.Agents `json:"result"`
}

func (self *ServerClient) GetQueries() (GetQueriesFromServerResponse, error) {
	res := GetQueriesFromServerResponse{}
	err := self.Call("GetQueries", nil, &res)

	return res, err
}

func (self *ServerClient) GetAgents() (GetAgentsFromServerResponse, error) {
	res := GetAgentsFromServerResponse{}
	err := self.Call("GetAgents", nil, &res)

	return res, err
}
