package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type UpdateNameRequest struct {
	*tchttp.BaseRequest
	InstanceId   string `name:"InstanceId"`
	InstanceName string `name:"InstanceName"`
}

type UpdateNameResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId"`
}

func (r *UpdateNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewUpdateNameRequest() (request *UpdateNameRequest) {
	request = &UpdateNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "UpdateEcsName")
	return
}

func NewUpdateNameResponse() (response *UpdateNameResponse) {
	response = &UpdateNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateName(request *UpdateNameRequest) (response *UpdateNameResponse, err error) {
	response = NewUpdateNameResponse()
	err = c.Send(request, response)
	return
}
