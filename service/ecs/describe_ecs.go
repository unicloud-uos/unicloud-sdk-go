package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DescribeEcsRequest struct {
	*tchttp.BaseRequest
	InstanceId   string `name:"InstanceId"`
	InstanceName string `name:"InstanceName"`
	IpAddr       string `name:"IpAddr"`
}

type DescribeEcsResponse struct {
	*tchttp.BaseResponse
	List       []EcsInstance `json:"list"`
	Size       int           `json:"size"`
	TotalCount int           `json:"totalCount"`
	TotalPages int           `json:"totalPages"`
	Page       int           `json:"page"`
	RequestId  string        `json:"RequestId"`
}

type EcsInstance struct {
	Instancecode     string `json:"instanceCode"`
	Starttime        int64  `json:"startTime"`
	Sysdiskid        string `json:"sysDiskId"`
	Status           string `json:"status"`
	IP               string `json:"ip"`
	Description      string `json:"description"`
	Instancecodename string `json:"instanceCodeName"`
	Sysdiskcode      string `json:"sysDiskCode"`
	Instancename     string `json:"instanceName"`
	Sysdisksize      int    `json:"sysDiskSize"`
	Instanceid       string `json:"instanceId"`
	Imageparentcode  string `json:"imageParentCode"`
	Eniid            string `json:"eniId"`
	Paytype          string `json:"payType"`
	Imagetype        string `json:"imageType"`
	Imageid          string `json:"imageId"`
	Imagecode        string `json:"imageCode"`
	Azoneid          string `json:"azoneId"`
	Binddiskcount    int    `json:"bindDiskCount"`
}

func (r *DescribeEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDescribeEcsRequest() (request *DescribeEcsRequest) {
	request = &DescribeEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeEcs").SetHttpMethod(tchttp.GET)
	return
}

func NewDescribeEcsResponse() (response *DescribeEcsResponse) {
	response = &DescribeEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEcs(request *DescribeEcsRequest) (response *DescribeEcsResponse, err error) {
	response = NewDescribeEcsResponse()
	err = c.Send(request, response)
	return
}
