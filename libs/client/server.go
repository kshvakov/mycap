package client

import (
	"mycap/agent"
	"mycap/libs/agrqueries"
	"mycap/libs/agrqueries/countpertime"
	"mycap/libs/jsonrpc"
)

type ServerClient struct {
	jsonrpc.JsonRpcClient
}

type GetQueriesFromServerResponse struct {
	jsonrpc.JsonRpcResponse
	Result agrqueries.QueriesAgregated `json:"result"`
}

func (self *ServerClient) GetQueries() (GetQueriesFromServerResponse, error) {
	res := GetQueriesFromServerResponse{}
	err := self.Call("GetQueries", nil, &res)

	return res, err
}

type GetCountPerTimeFromServerResponse struct {
	jsonrpc.JsonRpcResponse
	Result countpertime.Counters `json:"result"`
}

func (self *ServerClient) GetCountPerTime() (GetCountPerTimeFromServerResponse, error) {
	res := GetCountPerTimeFromServerResponse{}
	err := self.Call("GetCountPerTime", nil, &res)

	return res, err
}

type GetAgentsFromServerResponse struct {
	jsonrpc.JsonRpcResponse
	Result agent.Agents `json:"result"`
}

func (self *ServerClient) GetAgents() (GetAgentsFromServerResponse, error) {
	res := GetAgentsFromServerResponse{}
	err := self.Call("GetAgents", nil, &res)

	return res, err
}
