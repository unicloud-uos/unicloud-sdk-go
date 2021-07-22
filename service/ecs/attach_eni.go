package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type AttachEniRequest struct {
	*tchttp.BaseRequest
	EniId string `name:"EniId"`
	EcsId string `name:"EcsId"`
}

type AttachEniResponse struct {
	*tchttp.BaseResponse
	RequestId string `json:"RequestId"`
}

func (r *AttachEniResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newAttachEniRequest(eniId string, ecsId string) (request *AttachEniRequest) {
	request = &AttachEniRequest{
		BaseRequest: &tchttp.BaseRequest{},
		EniId:       eniId,
		EcsId:       ecsId,
	}
	request.Init().WithApiInfo(Service, APIVersion, "AttachEniToEcs").SetHttpMethod(tchttp.GET)
	return
}

func NewAttachEniResponse() (response *AttachEniResponse) {
	response = &AttachEniResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachEni(eniId string, ecsId string) (response *AttachEniResponse, err error) {
	request := newAttachEniRequest(eniId, ecsId)
	response = NewAttachEniResponse()
	err = c.Send(request, response)
	return
}
