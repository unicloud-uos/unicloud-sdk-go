package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type CreateDiskRequest struct {
	*tchttp.BaseRequest

	// cn-beijing-a cn-shanghai-a
	AzId string `name:"AzId"`

	//计费类型。YEAR_MONTH（包年包月）；DAY_MONTH（按日月结）；CHARGING_HOURS（按小时实时付费）
	PayType string `name:"PayType"`

	// 硬盘类型。ebs.highIO.ssd（SSD云硬盘）；ebs.hybrid.hdd（高性能HDD云硬盘）
	SpecificationCode string `name:"SpecificationCode"`

	// 硬盘大小。可选范围：20~32768 GB
	Capacity int `name:"Capacity"`

	// 硬盘数量。可选范围：1~15
	Quantity int `name:"Quantity"`
	//
	//// 租期。预付费必传，可选范围：1个月、2个月、3个月、6个月、1年
	//RentCount int `name:"RentCount"`
	//
	//// 租期单位。预付费必传，可选范围：day, month, year
	//RentUnit string `name:"RentUnit"`

	// 硬盘描述。
	Description string `name:"Description"`

	// 用快照创建硬盘时快照ID。
	SnapshotId string `name:"SnapshotId"`

	// 主机ID。创建包年包月硬盘必传该参数。
	VmId string `name:"VmId"`
}

type CreateDiskResponse struct {
	*tchttp.BaseResponse
	RequestId string   `name:"RequestId"`
	Ids       []string `name:"Ids"`
	OrderId   string   `name:"OrderId"`
}

func (r *CreateDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewCreateDiskRequest() (request *CreateDiskRequest) {
	request = &CreateDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "CreateDisk")
	return
}

func NewCreateDiskResponse() (response *CreateDiskResponse) {
	response = &CreateDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDisks(request *CreateDiskRequest) (response *CreateDiskResponse, err error) {
	response = NewCreateDiskResponse()
	err = c.Send(request, response)
	return
}
