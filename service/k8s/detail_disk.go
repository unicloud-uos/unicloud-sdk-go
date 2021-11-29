package k8s

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DetailDiskRequest struct {
	*tchttp.BaseRequest
	DiskId string `name:"DiskId"`
}

type DetailDiskResponse struct {
	*tchttp.BaseResponse
	RequestId string `json:"RequestId"`
	Status    bool   `name:"status"`
	Msg       string `name:"msg"`
	Disk      Disk   `json:"Res"`
	Disk2     Disk2  `json:"Resd"`
}

type Disk struct {
	DiskId       string `json:"diskId"`
	DiskName     string `json:"diskName"`
	UserId       string `json:"userId"`
	Status       string `json:"status"`
	InstanceCode string `json:"instanceCode"`
	Wwn          string `json:"wwn"`
}

type Disk2 struct{}

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
