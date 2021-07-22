package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type CreateEniArgs struct {
	RegionId    string `json:"regionId"`
	AzoneId     string `json:"azoneId"`
	Name        string `json:"name"`
	VpcId       string `json:"vpcId"`
	SubnetId    string `json:"subnetId"`
	SgId        string `json:"sgId"`
	ReleaseType int    `json:"releaseType"`
	IpAddress   string `json:"ipAddress"`
}

type CreateEniRequest struct {
	*tchttp.BaseRequest
	CreateEniArgs
}

type CreateEniResponse struct {
	*tchttp.BaseResponse
	RequestId  string `json:"RequestId"`
	InstanceId string `json:"instanceId"`
	MacAddr    string `json:"macAddr"`
	Ipv4Addr   string `json:"ipv4Addr"`
	Ipv6Addr   string `json:"ipv6Addr"`
}

func (r *CreateEniResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewCreateEniRequest(createEniArgs CreateEniArgs) (request *CreateEniRequest) {
	request = &CreateEniRequest{
		BaseRequest:   &tchttp.BaseRequest{},
		CreateEniArgs: createEniArgs,
	}
	request.Init().WithApiInfo(Service, APIVersion, "CreateEni").SetHttpMethod(tchttp.POST)
	request.SetBodyData(createEniArgs)
	return
}

func NewCreateEniResponse() (response *CreateEniResponse) {
	response = &CreateEniResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateEni(request *CreateEniRequest) (response *CreateEniResponse, err error) {
	response = NewCreateEniResponse()
	err = c.Send(request, response)
	return
}
