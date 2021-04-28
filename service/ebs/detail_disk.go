package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DetailDiskRequest struct {
	*tchttp.BaseRequest
	VolumeId string `name:"VolumeId" string`
}

type DetailDiskResponse struct {
	*tchttp.BaseResponse
	RequestId string `json:"RequestId"`
	Volume    []Disk `json:"Volume"`
}

type Disk struct {
	DiskId            string       `json:"diskId"`
	DiskName          string       `json:"diskName"`
	RegionId          string       `json:"regionId"`
	AzoneId           string       `json:"azoneId"`
	UserId            string       `json:"userId"`
	DiskType          string       `json:"diskType"`
	DiskSize          string       `json:"diskSize"`
	Status            string       `json:"status"`
	Description       string       `json:"description"`
	SpecificationCode string       `json:"specificationCode"`
	PayType           string       `json:"payType"`
	ChargeType        int          `json:"chargeType"`
	Expired           bool         `json:"expired"`
	AttachInfos       []AttachInfo `json:"attachInfos"`
}

type AttachInfo struct {
	InstanceId string `json:"instanceId"`
	Type       string `json:"type"`
}

func (r *DetailDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDetailDiskRequest() (request *DetailDiskRequest) {
	request = &DetailDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DetailDisk")
	return
}

func NewDetailDiskResponse() (response *DetailDiskResponse) {
	response = &DetailDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DetailDisk(request *DetailDiskRequest) (response *DetailDiskResponse, err error) {
	response = NewDetailDiskResponse()
	err = c.Send(request, response)
	return
}
