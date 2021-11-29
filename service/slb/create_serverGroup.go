package slb

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type CreateServerGroupArgs struct {
	Description string

	// 负载均衡实例 ID
	LoadbalancerId string

	Name string

	Algorithm string

	//CookieName    string

	ListenerId string

	Servers []Servers

	//SessionType string

}

type Servers struct {
	PortId     string
	ServerPort int32
	Weight     int
	ServerIp   string
	ServerId   string
}

func NewCreateServerGroupRequest(args CreateServerGroupArgs) (request *CreateServerGroupRequest) {
	request = &CreateServerGroupRequest{
		BaseRequest:           &tchttp.BaseRequest{},
		CreateServerGroupArgs: args,
	}
	request.Init().WithApiInfo(Service, APIVersion, "CreateServerGroup").SetHttpMethod(tchttp.POST)
	request.SetBodyData(args)
	return
}

type CreateServerGroupRequest struct {
	*tchttp.BaseRequest
	CreateServerGroupArgs
}

func (r *CreateServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

func NewCreateServerGroupResponse() (response *CreateServerGroupResponse) {
	response = &CreateServerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建一个ServerGroup
func (c *Client) CreateServerGroup(request *CreateServerGroupRequest) (response *CreateServerGroupResponse, err error) {
	response = NewCreateServerGroupResponse()
	err = c.Send(request, response)
	return
}

type CreateServerGroupResponse struct {
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
