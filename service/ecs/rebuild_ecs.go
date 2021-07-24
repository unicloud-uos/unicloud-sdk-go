package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type RebuildEcsRequest struct {
	*tchttp.BaseRequest
	InstanceId                  string `name:"InstanceId"`
	ImageId                     string `name:"ImageId"`
	ImageSpecificationClassCode string `name:"ImageSpecificationClassCode"`
	Password                    string `name:"Password"`
	KeyPair                     string `name:"KeyPair"`
}

type RebuildEcsResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId" string`
}

func (r *RebuildEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewRebuildEcsRequest() (request *RebuildEcsRequest) {
	request = &RebuildEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "RebuildEcs")
	return
}

func NewRebuildEcsResponse() (response *RebuildEcsResponse) {
	response = &RebuildEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RebuildEcs(request *RebuildEcsRequest) (response *RebuildEcsResponse, err error) {
	response = NewRebuildEcsResponse()
	err = c.Send(request, response)
	return
}
