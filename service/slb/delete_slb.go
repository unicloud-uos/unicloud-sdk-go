package slb

import tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"

func NewDeleteLoadBalancerRequest(deleteLoadBalancerArgs DeleteLoadBalancerArgs) (request *DeleteLoadBalancerRequest) {
	request = &DeleteLoadBalancerRequest{
		BaseRequest:            &tchttp.BaseRequest{},
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
