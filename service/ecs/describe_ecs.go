package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DescribeEcsRequest struct {
	*tchttp.BaseRequest
	InstanceId   string `name:"InstanceId"`
	Page         int    `name:"Page"`
	Size         int    `name:"Size"`
	InstanceName string `name:"InstanceName"`
	IpAddr       string `name:"IpAddr"`
}

type Detail struct {
	InstanceId       string   `name:"instanceId"`
	InstanceName     string   `name:"instanceName"`
	InstanceCode     string   `name:"instanceCode"`
	InstanceCodeName string   `name:"instanceCodeName"`
	Status           string   `name:"status"`
	ImageId          string   `name:"imageId"`
	ImageType        string   `name:"imageType"`
	ImageCode        string   `name:"imageCode"`
	ImageParentCode  string   `name:"imageParentCode"`
	Description      string   `name:"description"`
	Ip               string   `name:"ip"`
	EipId            string   `name:"eipId"`
	EipAddr          string   `name:"eipAddr"`
	EipSize          int      `name:"eipSize"`
	EipCode          string   `name:"eipCode"`
	SysDiskCode      string   `name:"sysDiskCode"`
	SysDiskSize      int      `name:"sysDiskSize"`
	StartTime        int64    `name:"startTime"`
	EndTime          int64    `name:"endTime"`
	BindDiskCount    int      `name:"bindDiskCount"`
	EniId            string   `name:"eniId"`
	SecondaryEni     []string `name:"secondaryEni"`
	AzoneId          string   `name:"azoneId"`
	PayType          string   `name:"payType"`
}

type DescribeEcsResponse struct {
	*tchttp.BaseResponse
	RequestId  string   `name:"RequestId"`
	Page       int      `name:"page"`
	Size       int      `name:"size"`
	TotalCount int      `name:"totalCount"`
	TotalPages int      `name:"totalPages"`
	List       []Detail `name:"list"`
}

func (r *DescribeEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDescribeEcsRequest() (request *DescribeEcsRequest) {
	request = &DescribeEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeEcs")
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
