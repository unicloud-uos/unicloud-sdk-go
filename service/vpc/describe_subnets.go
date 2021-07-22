package vpc

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest
	VpcId string `name:"VpcId"`
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse
	Code      string   `json:"Code"`
	Msg       string   `json:"Msg"`
	Subnets   []Subnet `json:"Res"`
	RequestId string   `json:"RequestId"`
}

type Subnet struct {
	Name            string `json:"Name"`
	IsCrosszone     bool   `json:"IsCrossZone"`
	IpAssignedcount int    `json:"IpAssignedCount"`
	Cidr            string `json:"Cidr"`
	AzoneId         string `json:"AzoneId"`
	GatewayIp       string `json:"GatewayIp"`
	Iptotalcount    int    `json:"IpTotalCount"`
	Id              string `json:"Id"`
	Ipv6Cidr        string `json:"IpV6Cidr"`
	Ipv6GatewayIp   string `json:"IpV6GatewayIp"`
	Hasipv6         bool   `json:"HasIpV6"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
	request = &DescribeSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeSubnet").SetHttpMethod(tchttp.GET)
	return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
	response = &DescribeSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
	response = NewDescribeSubnetsResponse()
	err = c.Send(request, response)
	return
}
