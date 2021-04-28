package slb

import (
	"encoding/json"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

const (
	UserQuotanotEnough = "Network.Order.UserQuotanotEnough"
)

func NewCreateLoadBalancerRequest(createLoadBalancerArgs CreateLoadBalancerArgs) (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		BaseRequest:            &tchttp.BaseRequest{},
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

// 创建一个slb模拟配额超过
func (c *Client) CreateLoadBalancersMockQuotaOut(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = NewCreateLoadBalancerResponse()
	err = errors.NewUnicloudCloudSDKError("Network.Order.UserQuotanotEnough", "用户配额不足", "")
	return
}

// 创建一个slb模拟配额超过
func (c *Client) CreateLoadBalancersMockInit(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = NewCreateLoadBalancerResponse()
	response.Res.Resources = []Resources{
		{
			InstanceId: "slb-ks0ouqxkvk42",
		},
	}
	return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *CreateLoadBalancerResponse) Id() string {
	return c.Res.Resources[0].InstanceId
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 由负载均衡实例唯一 ID 组成的数组。
		LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
	RequestId string
	Code      string `json:"Code`
	Msg       string `json:"Msg`
	Res       struct {
		Resources []Resources `json:"Resources`
	} `json:"Res`
}

type Resources struct {
	Product    string `json:"Product"`
	InstanceId string `json:"InstanceId"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	CreateLoadBalancerArgs CreateLoadBalancerArgs
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
