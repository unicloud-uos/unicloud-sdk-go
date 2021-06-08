package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type PortBindRequest struct {
	*tchttp.BaseRequest
	PortBindArgs
}

type PortBindArgs struct {
	PortId         string   `json:"PortId"`
	SecurityGroups []string `json:"SecurityGroupIds"`
}

type PortBindResponse struct {
	*BaseResponse
	Res interface{} `json:"res",omitempty`
}

func (r *PortBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newPortBindRequest(args PortBindArgs) (request *PortBindRequest) {
	request = &PortBindRequest{
		BaseRequest:  &tchttp.BaseRequest{},
		PortBindArgs: args,
	}
	request.Init().WithApiInfo(Service, APIVersion, "PortBindSecurityGroup").SetHttpMethod(tchttp.POST)
	request.SetBodyData(args)
	return
}

func newPortBindResponse() (response *PortBindResponse) {
	response = &PortBindResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) PortBind(portId string, sgId string) (response *PortBindResponse, err error) {
	request := newPortBindRequest(PortBindArgs{PortId: portId, SecurityGroups: []string{sgId}})
	response = newPortBindResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	}
	return
}
