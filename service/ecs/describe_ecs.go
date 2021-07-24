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
	InstanceId       string   `name:"instanceId" string`
	InstanceName     string   `name:"instanceName" string`
	InstanceCode     string   `name:"instanceCode" string`
	InstanceCodeName string   `name:"instanceCodeName" string`
	Status           string   `name:"status" string`
	ImageId          string   `name:"imageId" string`
	ImageType        string   `name:"imageType" string`
	ImageCode        string   `name:"imageCode" string`
	ImageParentCode  string   `name:"imageParentCode" string`
	Description      string   `name:"description" string`
	Ip               string   `name:"ip" string`
	EipId            string   `name:"eipId" string`
	EipAddr          string   `name:"eipAddr" string`
	EipSize          int      `name:"eipSize" int`
	EipCode          string   `name:"eipCode" string`
	SysDiskCode      string   `name:"sysDiskCode" string`
	SysDiskSize      int      `name:"sysDiskSize" int`
	StartTime        int64    `name:"startTime" int64 `
	EndTime          int64    `name:"endTime" int64`
	BindDiskCount    int      `name:"bindDiskCount" int`
	EniId            string   `name:"eniId" string`
	SecondaryEni     []string `name:"secondaryEni" []string`
	AzoneId          string   `name:"azoneId" string`
	PayType          string   `name:"payType" string`
}

type DescribeEcsResponse struct {
	*tchttp.BaseResponse
	RequestId  string   `name:"RequestId" string`
	Page       int      `name:"page" int`
	Size       int      `name:"size" int`
	TotalCount int      `name:"totalCount" int`
	TotalPages int      `name:"totalPages" int`
	List       []Detail `name: "list"`
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
