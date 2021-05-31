package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DescribeProductSpecRequest struct {
	*tchttp.BaseRequest
	SpecificationCode string `name:"specificationCode"`
}

type DescribeProductSpecResponse struct {
	*BaseResponse
	Status bool        `json:"status"`
	Auth   bool        `json:"auth"`
	Msg    string      `json:"msg"`
	Code   string      `json:"code"`
	Res    ProductSpec `json:"res"`
}

type ProductSpec struct {
	Componentcode          string `json:"componentCode"`
	Createtime             int64  `json:"createTime"`
	Describe               string `json:"describe"`
	GpuCount               string `json:"gpu_count"`
	GpuType                string `json:"gpu_type"`
	Memory                 string `json:"memory"`
	MemoryUnit             string `json:"memory_unit"`
	Specificationclasscode string `json:"specificationClassCode"`
	Specificationcode      string `json:"specificationCode"`
	Specificationid        string `json:"specificationId"`
	Specificationname      string `json:"specificationName"`
	Updatetime             int64  `json:"updateTime"`
	Vcpu                   string `json:"vCPU"`
	VcpuUnit               string `json:"vCPU_unit"`
	EniCount               string `json:"eni_count""`
	IpCountPerEni          string `json:"ip_count_per_eni""`
}

func (r *DescribeProductSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newDescribeProductSpecRequest() (request *DescribeProductSpecRequest) {
	request = &DescribeProductSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeProductSpecification").SetHttpMethod(tchttp.GET)
	return
}

func NewDescribeProductSpecResponse() (response *DescribeProductSpecResponse) {
	response = &DescribeProductSpecResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) DescribeProductSpec(specificationCode string) (res *ProductSpec, err error) {
	request := newDescribeProductSpecRequest()
	request.SpecificationCode = specificationCode
	response := NewDescribeProductSpecResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	} else {
		res = &response.Res
	}
	return
}
