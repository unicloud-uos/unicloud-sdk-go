package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type StopEcsRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
}

type StopEcsResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId" string`
}

func (r *StopEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewStopEcsRequest() (request *StopEcsRequest) {
	request = &StopEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "StopEcs")
	return
}

func NewStopEcsResponse() (response *StopEcsResponse) {
	response = &StopEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) StopEcs(request *StopEcsRequest) (response *StopEcsResponse, err error) {
	response = NewStopEcsResponse()
	err = c.Send(request, response)
	return
}
