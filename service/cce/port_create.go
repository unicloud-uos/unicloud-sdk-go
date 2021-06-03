package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
	"log"
)

type PortCreateRequest struct {
	*tchttp.BaseRequest
	PortCreateArgs
}

type PortCreateArgs struct {
	AssistIpAddress []string `json:"AssistIpAddress,omitempty"`
	AssistIpCount   int      `json:"AssistIpCount,omitempty"`
	AvailableZone   string   `json:"AvailableZone"`
	Name            string   `json:"Name"`
	SecurityGroups  []string `json:"Security_groups"`
	SubnetId        string   `json:"SubnetId"`
}

type PortCreateResponse struct {
	*BaseResponse
	Res PortCreateRes `json:"res"`
}

type PortCreateRes struct {
	Port Port `json:"Port"`
}

func (r *PortCreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newPortCreateRequest(args PortCreateArgs) (request *PortCreateRequest) {
	request = &PortCreateRequest{
		BaseRequest:    &tchttp.BaseRequest{},
		PortCreateArgs: args,
	}
	request.Init().WithApiInfo(Service, APIVersion, "PortCreate").SetHttpMethod(tchttp.POST)
	request.SetBodyData(args)
	return
}

func newPortCreateResponse() (response *PortCreateResponse) {
	response = &PortCreateResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) PortCreate(args PortCreateArgs) (port Port, err error) {
	argJson, err := json.Marshal(args)
	if err != nil {
		return
	}
	log.Println(string(argJson))
	request := newPortCreateRequest(args)
	response := newPortCreateResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	}
	port = response.Res.Port
	return
}
