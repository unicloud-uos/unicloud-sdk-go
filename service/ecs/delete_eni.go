package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DeleteEniRequest struct {
	*tchttp.BaseRequest
	EniId string `name:"EniId"`
}

type DeleteEniResponse struct {
	*tchttp.BaseResponse
	RequestId string `json:"RequestId"`
}

func (r *DeleteEniResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newDeleteEniRequest(eniId string) (request *DeleteEniRequest) {
	request = &DeleteEniRequest{
		BaseRequest: &tchttp.BaseRequest{},
		EniId:       eniId,
	}
	request.Init().WithApiInfo(Service, APIVersion, "DeleteEni").SetHttpMethod(tchttp.GET)
	return
}

func NewDeleteEniResponse() (response *DeleteEniResponse) {
	response = &DeleteEniResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteEni(eniId string) (response *DeleteEniResponse, err error) {
	request := newDeleteEniRequest(eniId)
	response = NewDeleteEniResponse()
	err = c.Send(request, response)
	return
}
