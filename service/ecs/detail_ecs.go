package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DetailEcsRequest struct {
	*tchttp.BaseRequest
	InstanceId string `name:"InstanceId"`
}

type DetailEcsResponse struct {
	*tchttp.BaseResponse
	RequestId       string   `name:"RequestId" string`
	Id              string   `name:"id" string`
	Name            string   `name:"name" string`
	Status          string   `name:"status" string`
	Description     string   `name:"description" string`
	KeyPair         string   `name:"keyPair" string`
	StartTime       int64    `name:"startTime" int64`
	EndTime         int64    `name:"endTime" int64`
	RegionId        string   `name:"RegionId" string`
	RegionName      string   `name:"RegionName" string`
	Cpu             int      `name:"cpu" int`
	Memory          int      `name:"memory" int`
	SysDiskCode     string   `name:"sysDiskCode" string`
	SysDiskSize     int      `name:"sysDiskSize" int`
	EipId           string   `name:"eipId" string`
	EipIp           string   `name:"eipIp" string`
	BandWidth       int      `name:bandWidth int`
	Ip              string   `name:"ip" string`
	ImageId         string   `name:"imageId" string`
	ImageType       string   `name:"imageType" string`
	ImageCode       string   `name:"imageCode" string`
	ImageParentCode string   `name:"imageParentCode" string`
	EniId           string   `name:"eniId" string`
	DataDiskIds     []string `name:"dataDiskIds" []string`
	SgIds           []string `name:"sgIds" []string`
}

func (r *DetailEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDetailEcsRequest() (request *DetailEcsRequest) {
	request = &DetailEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DetailEcs")
	return
}

func NewDetailEcsResponse() (response *DetailEcsResponse) {
	response = &DetailEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DetailEcs(request *DetailEcsRequest) (response *DetailEcsResponse, err error) {
	response = NewDetailEcsResponse()
	err = c.Send(request, response)
	return
}
