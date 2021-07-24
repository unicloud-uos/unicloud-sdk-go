package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type SystemDiskRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
	Size       int    `name:"size"`
}

type SystemDiskResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId" string`
}

func (r *SystemDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewSystemDiskRequest() (request *SystemDiskRequest) {
	request = &SystemDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "ResizeEcsSystemDisk")
	return
}

func NewSystemDiskResponse() (response *SystemDiskResponse) {
	response = &SystemDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) SystemDisk(request *SystemDiskRequest) (response *SystemDiskResponse, err error) {
	response = NewSystemDiskResponse()
	err = c.Send(request, response)
	return
}
