package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type StartEcsRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
}

type StartEcsResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId" string`
}

func (r *StartEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewStartEcsRequest() (request *StartEcsRequest) {
	request = &StartEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "StartEcs")
	return
}

func NewStartEcsResponse() (response *StartEcsResponse) {
	response = &StartEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) StartEcs(request *StartEcsRequest) (response *StartEcsResponse, err error) {
	response = NewStartEcsResponse()
	err = c.Send(request, response)
	return
}
