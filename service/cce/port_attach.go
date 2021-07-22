package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type PortAttachRequest struct {
	*tchttp.BaseRequest
	PortAttachArgs
}

type PortAttachArgs struct {
	EniId  string `json:"EniId"`
	ZoneId string `json:"ZoneId"`
	NodeId string `json:"NodeId"`
}

type PortAttachResponse struct {
	*BaseResponse
	Res PortAttachRes `json:"res"`
}

type PortAttachRes struct {
	State  string `json:"State"`
	TaskId string `json:"TaskId"`
}

func (r *PortAttachResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newPortAttachRequest(args PortAttachArgs) (request *PortAttachRequest) {
	request = &PortAttachRequest{
		BaseRequest:    &tchttp.BaseRequest{},
		PortAttachArgs: args,
	}
	request.Init().WithApiInfo(Service, APIVersion, "PortAttach").SetHttpMethod(tchttp.POST)
	request.SetBodyData(args)
	return
}

func newPortAttachResponse() (response *PortAttachResponse) {
	response = &PortAttachResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) PortAttach(portId string, zoneId string, nodeId string) (response *PortAttachResponse, err error) {
	request := newPortAttachRequest(PortAttachArgs{EniId: portId, ZoneId: zoneId, NodeId: nodeId})
	response = newPortAttachResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	}
	return
}
