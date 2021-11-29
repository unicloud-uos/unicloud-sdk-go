package ecs

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
	RequestId       string   `name:"RequestId"`
	Id              string   `name:"id"`
	Name            string   `name:"name"`
	Status          string   `name:"status"`
	Description     string   `name:"description"`
	KeyPair         string   `name:"keyPair"`
	StartTime       int64    `name:"startTime"`
	EndTime         int64    `name:"endTime"`
	RegionId        string   `name:"RegionId"`
	RegionName      string   `name:"RegionName"`
	Cpu             int      `name:"cpu"`
	Memory          int      `name:"memory"`
	SysDiskCode     string   `name:"sysDiskCode"`
	SysDiskSize     int      `name:"sysDiskSize"`
	EipId           string   `name:"eipId"`
	EipIp           string   `name:"eipIp"`
	BandWidth       int      `name:"bandWidth"`
	Ip              string   `name:"ip"`
	ImageId         string   `name:"imageId"`
	ImageType       string   `name:"imageType"`
	ImageCode       string   `name:"imageCode"`
	ImageParentCode string   `name:"imageParentCode"`
	EniId           string   `name:"eniId"`
	DataDiskIds     []string `name:"dataDiskIds"`
	SgIds           []string `name:"sgIds"`
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
