package slb

import (
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/profile"
)

const (
	APIVersion = "20200730"
)

type Client struct {
	common.Client
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

/**
 * 从环境变量中读取
 */
func NewClientFromEnv() (client *Client)  {
	client = &Client{}
	client.InitFromEnv()
	return
}

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

func NewCreateLoadBalancerRequest(createLoadBalancerArgs CreateLoadBalancerArgs) (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
		CreateLoadBalancerArgs: createLoadBalancerArgs,
	}
	request.Init().WithApiInfo("networks/slb", APIVersion, "CreateSlb").SetHttpMethod(tchttp.POST)
	request.SetBodyData(createLoadBalancerArgs)
	return
}

// 创建一个slb
func (c *Client) CreateLoadBalancers(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = NewCreateLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func NewDeleteLoadBalancerRequest(deleteLoadBalancerArgs DeleteLoadBalancerArgs) (request *DeleteLoadBalancerRequest) {
	request = &DeleteLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
		DeleteLoadBalancerArgs: deleteLoadBalancerArgs,
	}
	request.Init().WithApiInfo("networks/slb", APIVersion, "DeleteSlb").SetHttpMethod(tchttp.PUT)
	request.SetBodyData(deleteLoadBalancerArgs)
	return
}

func (c *Client) DeleteListener(request *DeleteLoadBalancerRequest) (response *DeleteListenerResponse, err error) {
	response = NewDeleteListenerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
	response = &DeleteListenerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}