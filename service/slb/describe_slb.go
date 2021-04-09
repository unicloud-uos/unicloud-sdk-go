package slb

import tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
	request = &DescribeLoadBalancersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("networks/slb", APIVersion, "DescribeSlb")
	return
}

// 查询一个地域的负载均衡实例列表
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
	if request == nil {
		request = NewDescribeLoadBalancersRequest()
	}
	response = NewDescribeLoadBalancersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
	response = &DescribeLoadBalancersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}