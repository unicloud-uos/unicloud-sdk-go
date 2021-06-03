package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type PortDeleteRequest struct {
	*tchttp.BaseRequest
	PortId string `name:"portId"`
}

type PortDeleteArgs struct {
	AssistIpAddress []string `json:AssistIpAddress`
	AssistIpCount   int      `json:AssistIpCount`
	AvailableZone   string   `json:"AvailableZone"`
	Name            string   `json:"Name"`
	SecurityGroups  []string `json:"Security_groups"`
	SubnetId        string   `json:"SubnetId"`
	TenantId        string   `json:"TenantId"`
}

type PortDeleteResponse struct {
	*BaseResponse
	Res string `json:"res"`
}

func (r *PortDeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newPortDeleteRequest(portId string) (request *PortDeleteRequest) {
	request = &PortDeleteRequest{
		BaseRequest: &tchttp.BaseRequest{},
		PortId:      portId,
	}
	request.Init().WithApiInfo(Service, APIVersion, "PortDelete").SetHttpMethod(tchttp.POST)
	return
}

func newPortDeleteResponse() (response *PortDeleteResponse) {
	response = &PortDeleteResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) PortDelete(portId string) (response *PortDeleteResponse, err error) {
	request := newPortDeleteRequest(portId)
	response = newPortDeleteResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	}
	return
}
