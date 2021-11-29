package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type UpdateConfigRequest struct {
	*tchttp.BaseRequest
	InstanceId   string `name:"InstanceId"`
	InstanceCode string `name:"InstanceCode"`
}

type UpdateConfigResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId"`
}

func (r *UpdateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewUpdateConfigRequest() (request *UpdateConfigRequest) {
	request = &UpdateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "UpdateEcsConfig")
	return
}

func NewUpdateConfigResponse() (response *UpdateConfigResponse) {
	response = &UpdateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateConfig(request *UpdateConfigRequest) (response *UpdateConfigResponse, err error) {
	response = NewUpdateConfigResponse()
	err = c.Send(request, response)
	return
}
