package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type RebootEcsRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
}

type RebootEcsResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId"`
}

func (r *RebootEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewRebootEcsRequest() (request *RebootEcsRequest) {
	request = &RebootEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "RebootEcs")
	return
}

func NewRebootEcsResponse() (response *RebootEcsResponse) {
	response = &RebootEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RebootEcs(request *RebootEcsRequest) (response *RebootEcsResponse, err error) {
	response = NewRebootEcsResponse()
	err = c.Send(request, response)
	return
}
