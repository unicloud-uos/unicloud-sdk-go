package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DetachEniRequest struct {
	*tchttp.BaseRequest
	EniId string `name:"EniId"`
}

type DetachEniResponse struct {
	*tchttp.BaseResponse
	RequestId string `json:"RequestId"`
}

func (r *DetachEniResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newDetachEniRequest(eniId string) (request *DetachEniRequest) {
	request = &DetachEniRequest{
		BaseRequest: &tchttp.BaseRequest{},
		EniId:       eniId,
	}
	request.Init().WithApiInfo(Service, APIVersion, "detachEniFromEcs").SetHttpMethod(tchttp.GET)
	return
}

func NewDetachEniResponse() (response *DetachEniResponse) {
	response = &DetachEniResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachEni(eniId string) (response *DetachEniResponse, err error) {
	request := newDetachEniRequest(eniId)
	response = NewDetachEniResponse()
	err = c.Send(request, response)
	return
}
