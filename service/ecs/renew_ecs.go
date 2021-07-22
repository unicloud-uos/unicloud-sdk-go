package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type RenewEcsRequest struct {
	*tchttp.BaseRequest
	VmId        []string `name:"VmId"`
	BandWidthId []string `name:"BandWidthId"`
	Period      int      `name:"Period"`
}

type RenewEcsResponse struct {
	*tchttp.BaseResponse
	RequestId string `name:"RequestId" string`
}

func (r *RenewEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewRenewEcsRequest() (request *RenewEcsRequest) {
	request = &RenewEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "RenewEcs")
	return
}

func NewRenewEcsResponse() (response *RenewEcsResponse) {
	response = &RenewEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RenewEcs(request *RenewEcsRequest) (response *RenewEcsResponse, err error) {
	response = NewRenewEcsResponse()
	err = c.Send(request, response)
	return
}
