package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DescribeDisksRequest struct {
	*tchttp.BaseRequest
	Size       int    `name:"Size"`
	Page       int    `name:"Page"`
	InstanceId string `name:"InstanceId"`
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足过滤条件的负载均衡实例总数。此数值与入参中的Limit无关。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
	Current   int    `json:"current"`
	Size      int    `json:"size"`
	Pages     int    `json:"pages"`
	Total     int    `json:"total"`
	Records   []Disk `json:"records"`
	RequestId string `json:"RequestId"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeDisks")
	return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
	response = &DescribeDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	response = NewDescribeDisksResponse()
	err = c.Send(request, response)
	return
}
