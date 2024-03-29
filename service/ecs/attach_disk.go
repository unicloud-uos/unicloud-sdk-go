package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type AttachDiskRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
	DiskId     string `name:"DiskId"`
}

type AttachDiskResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId"`
}

func (r *AttachDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewAttachDiskRequest() (request *AttachDiskRequest) {
	request = &AttachDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "EcsAttachDisk")
	return
}

func NewAttachDiskResponse() (response *AttachDiskResponse) {
	response = &AttachDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachDisk(request *AttachDiskRequest) (response *AttachDiskResponse, err error) {
	response = NewAttachDiskResponse()
	err = c.Send(request, response)
	return
}
