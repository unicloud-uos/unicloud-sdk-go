package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DeleteDiskRequest struct {
	*tchttp.BaseRequest
	VolumeIds string `name:"VolumeIds"`
}

type DeleteDiskResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId"`
}

func (r *DeleteDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDeleteDiskRequest() (request *DeleteDiskRequest) {
	request = &DeleteDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DeleteDisk")
	return
}

func NewDeleteDiskResponse() (response *DeleteDiskResponse) {
	response = &DeleteDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDisk(request *DeleteDiskRequest) (response *DeleteDiskResponse, err error) {
	response = NewDeleteDiskResponse()
	err = c.Send(request, response)
	return
}
