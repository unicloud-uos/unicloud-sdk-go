package slb

import tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"

const (
	UserQuotanotEnough = "Network.Order.UserQuotanotEnough"
)

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