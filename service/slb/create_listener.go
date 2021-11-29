package slb

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type CreateListenerArgs struct {
	Description string

	// 负载均衡实例 ID
	LoadbalancerId string

	Name string
	// 监听器协议： TCP | UDP | HTTP | HTTPS | TCP_SSL（TCP_SSL 正在内测中，如需使用请通过工单申请）
	Protocol string `json:"Protocol,omitempty" name:"Protocol"`

	ProtocolPort int

	ConnectionLimit int `json:"ConnectionLimit,omitempty"`

	CertificateId string

	UserId string

	//LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	//
	//// 要将监听器创建到哪些端口，每个端口对应一个新的监听器
	//Ports []*int64 `json:"Ports,omitempty" name:"Ports"`
	//
	//// 要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数
	//ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`
	//
	//// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器
	//HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	//
	//// 证书相关信息，此参数仅适用于TCP_SSL监听器和未开启SNI特性的HTTPS监听器。
	//Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
	//
	//// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	//SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`
	//
	//// 监听器转发的方式。可选值：WRR、LEAST_CONN
	//// 分别表示按权重轮询、最小连接数， 默认为 WRR。此参数仅适用于TCP/UDP/TCP_SSL监听器。
	//Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
	//
	//// 是否开启SNI特性，此参数仅适用于HTTPS监听器。
	//SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`
	//
	//// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。
	//TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	//
	//// 会话保持类型。不传或传NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。
	//SessionType *string `json:"SessionType,omitempty" name:"SessionType"`
	//
	//// 是否开启长连接，此参数仅适用于HTTP/HTTPS监听器，0:关闭；1:开启， 默认关闭
	//KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`
	//
	//// 创建端口段监听器时必须传入此参数，用以标识结束端口。同时，入参Ports只允许传入一个成员，用以标识开始端口。【如果您需要体验端口段功能，请通过 [工单申请](https://console.cloud.tencent.com/workorder/category)】。
	//EndPort *uint64 `json:"EndPort,omitempty" name:"EndPort"`
}

func NewCreateListenerRequest(createListenerArgs CreateListenerArgs) (request *CreateListenerRequest) {
	request = &CreateListenerRequest{
		BaseRequest:        &tchttp.BaseRequest{},
		CreateListenerArgs: createListenerArgs,
	}
	request.Init().WithApiInfo(Service, APIVersion, "CreateListener").SetHttpMethod(tchttp.POST)
	request.SetBodyData(createListenerArgs)
	return
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	CreateListenerArgs
}

func (r *CreateListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
	response = &CreateListenerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建一个slb
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
	response = NewCreateListenerResponse()
	err = c.Send(request, response)
	return
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的监听器的唯一标识数组
		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
	Code string `json:"Code"`
	Msg  string `json:"Msg"`
	Res  struct {
		Ids []string `json:"Ids"`
	} `json:"Res"`
	RequestId string `json:"RequestId"`
}
