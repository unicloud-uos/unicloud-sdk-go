package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type PortAssignIpAddressesRequest struct {
	*tchttp.BaseRequest
	PortAssignIpAddressesArgs
}

type PortAssignIpAddressesArgs struct {
	PortId        string `json:"PortId"`
	AssistIpCount int    `json:"AssistIpCount"`
}

type PortAssignIpAddressesResponse struct {
	*BaseResponse
	Res PortAssignIpAddressesRes `json:"res"`
}

type PortAssignIpAddressesRes struct {
	Fixedips []FixedIp `json:"FixedIps"`
	ID       string    `json:"Id"`
}

func (r *PortAssignIpAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newPortAssignIpAddressesRequest(args PortAssignIpAddressesArgs) (request *PortAssignIpAddressesRequest) {
	request = &PortAssignIpAddressesRequest{
		BaseRequest:               &tchttp.BaseRequest{},
		PortAssignIpAddressesArgs: args,
	}
	request.Init().WithApiInfo(Service, APIVersion, "PortAssignIpAddresses").SetHttpMethod(tchttp.POST)
	request.SetBodyData(args)
	return
}

func newPortAssignIpAddressesResponse() (response *PortAssignIpAddressesResponse) {
	response = &PortAssignIpAddressesResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) PortAssignIp(portId string, count int) (fixedips []FixedIp, err error) {
	request := newPortAssignIpAddressesRequest(PortAssignIpAddressesArgs{PortId: portId, AssistIpCount: count})
	response := newPortAssignIpAddressesResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	} else {
		fixedips = response.Res.Fixedips
	}
	return
}
