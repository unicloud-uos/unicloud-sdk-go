package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DeleteEcsRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
}

type DeleteEcsResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId" string`
}

func (r *DeleteEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDeleteEcsRequest() (request *DeleteEcsRequest) {
	request = &DeleteEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DeleteEcs")
	return
}

func NewDeleteEcsResponse() (response *DeleteEcsResponse) {
	response = &DeleteEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteEcs(request *DeleteEcsRequest) (response *DeleteEcsResponse, err error) {
	response = NewDeleteEcsResponse()
	err = c.Send(request, response)
	return
}
