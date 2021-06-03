package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type PortDetachRequest struct {
	*tchttp.BaseRequest
	PortDetachArgs
}

type PortDetachArgs struct {
	EniId  string `json:"EniId"`
	NodeId string `json:"NodeId"`
}

type PortDetachResponse struct {
	*BaseResponse
	Res PortDetachRes `json:"res"`
}

type PortDetachRes struct {
	State  string `json:"State"`
	TaskId string `json:"TaskId"`
}

func (r *PortDetachResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newPortDetachRequest(args PortDetachArgs) (request *PortDetachRequest) {
	request = &PortDetachRequest{
		BaseRequest:    &tchttp.BaseRequest{},
		PortDetachArgs: args,
	}
	request.Init().WithApiInfo(Service, APIVersion, "PortDetach").SetHttpMethod(tchttp.POST)
	request.SetBodyData(args)
	return
}

func newPortDetachResponse() (response *PortDetachResponse) {
	response = &PortDetachResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) PortDetach(portId string, nodeId string) (response *PortDetachResponse, err error) {
	request := newPortDetachRequest(PortDetachArgs{EniId: portId, NodeId: nodeId})
	response = newPortDetachResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	}
	return
}
