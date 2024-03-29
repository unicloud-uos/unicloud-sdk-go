package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DetachDiskRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
	DiskId     string `name:"DiskId"`
}

type DetachDiskResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId"`
}

func (r *DetachDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDetachDiskRequest() (request *DetachDiskRequest) {
	request = &DetachDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "EcsDetachDisk")
	return
}

func NewDetachDiskResponse() (response *DetachDiskResponse) {
	response = &DetachDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachDisk(request *DetachDiskRequest) (response *DetachDiskResponse, err error) {
	response = NewDetachDiskResponse()
	err = c.Send(request, response)
	return
}
