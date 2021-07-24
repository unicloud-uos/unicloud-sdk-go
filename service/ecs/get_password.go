package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type GetPasswordRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
}

type GetPasswordResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId" string`
	Password  string `name:"password" string`
}

func (r *GetPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewGetPasswordRequest() (request *GetPasswordRequest) {
	request = &GetPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "GetEcsPassword")
	return
}

func NewGetPasswordResponse() (response *GetPasswordResponse) {
	response = &GetPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPassword(request *GetPasswordRequest) (response *GetPasswordResponse, err error) {
	response = NewGetPasswordResponse()
	err = c.Send(request, response)
	return
}
