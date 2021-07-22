package cce

import (
	"encoding/json"
	"errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type QueryPortRequest struct {
	*tchttp.BaseRequest
	QueryPortArgs
}

type QueryPortArgs struct {
	DeviceIds []string `json:"DeviceIds,omitempty"`
	Ids       []string `json:"Ids,omitempty"`
}

type QueryPortResponse struct {
	*BaseResponse
	Res QueryPortRes `json:"res"`
}

type QueryPortRes struct {
	Count string  `json:"Count"`
	Ports PortSet `json:"Ports"`
}

type PortSet []Port

func (s *PortSet) UnmarshalJSON(data []byte) error {
	var port = make([]Port, 0)
	var err error
	if string(data) != "{}" {
		err = json.Unmarshal(data, &port)
	}
	*s = port
	return err
}

func (r *QueryPortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func newQueryPortRequest(args QueryPortArgs) (request *QueryPortRequest) {
	request = &QueryPortRequest{
		BaseRequest:   &tchttp.BaseRequest{},
		QueryPortArgs: args,
	}
	request.Init().WithApiInfo(Service, APIVersion, "QueryPort").SetHttpMethod(tchttp.POST)
	request.SetBodyData(args)
	return
}

func newQueryPortResponse() (response *QueryPortResponse) {
	response = &QueryPortResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (c *Client) QueryPortByEniId(eniId string) (port Port, err error) {
	request := newQueryPortRequest(QueryPortArgs{Ids: []string{eniId}})
	response := newQueryPortResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	}
	if len(response.Res.Ports) > 0 {
		port = response.Res.Ports[0]
	} else {
		err = errors.New("port not found, eniId:" + eniId)
	}
	return
}

func (c *Client) QueryPortsByDeviceId(deviceId string) (ports []Port, err error) {
	request := newQueryPortRequest(QueryPortArgs{DeviceIds: []string{deviceId}})
	response := newQueryPortResponse()
	err = c.Send(request, response)
	if !response.Status {
		err = errors.New(response.Msg)
	}
	ports = response.Res.Ports
	return
}
