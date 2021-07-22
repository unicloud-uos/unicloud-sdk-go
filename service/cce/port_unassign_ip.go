package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type PortUnAssignIpAddressesRequest struct {
	*tchttp.BaseRequest
	PortUnAssignIpAddressesArgs
}

type PortUnAssignIpAddressesArgs struct {
	PortId          string            `json:"PortId"`
	AssistIpAddress []AssistIpAddress `json:"AssistIpAddress"`
}

type AssistIpAddress struct {
	Ipv4ID    string `json:"IPv4Id"`
	Ipaddress string `json:"IpAddress"`
}

type PortUnAssignIpAddressesResponse struct {
	*BaseResponse
	Res string `json:"res"`
}

func (r *PortUnAssignIpAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newPortUnAssignIpAddressesRequest(args PortUnAssignIpAddressesArgs) (request *PortUnAssignIpAddressesRequest) {
	request = &PortUnAssignIpAddressesRequest{
		BaseRequest:                 &tchttp.BaseRequest{},
		PortUnAssignIpAddressesArgs: args,
	}
	request.Init().WithApiInfo(Service, APIVersion, "PortUnAssignIpAddresses").SetHttpMethod(tchttp.POST)
	request.SetBodyData(args)
	return
}

func newPortUnAssignIpAddressesResponse() (response *PortUnAssignIpAddressesResponse) {
	response = &PortUnAssignIpAddressesResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) PortUnAssignIp(portId string, ips []string) (response *PortUnAssignIpAddressesResponse, err error) {
	var ipAddress []AssistIpAddress
	for _, ip := range ips {
		ipAddress = append(ipAddress, AssistIpAddress{Ipaddress: ip})
	}
	args := PortUnAssignIpAddressesArgs{PortId: portId, AssistIpAddress: ipAddress}
	request := newPortUnAssignIpAddressesRequest(args)
	response = newPortUnAssignIpAddressesResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	}
	return
}
