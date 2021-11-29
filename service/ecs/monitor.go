package ecs

import tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"

type MonitorRequest struct {
	*tchttp.BaseRequest
	InstanceId   string `name:"InstanceId"`
	DateCategory string `name:"DateCategory"`
	StartTime    int64  `name:"StartTime"`
	EndTime      int64  `name:"EndTime"`
	DestUnit     int64  `name:"DestUnit"`
}
type MonitorResponse struct {
	*tchttp.BaseResponse
	RequestId string            `name:"RequestId"`
	data      map[int64]float64 `name:"RequestId"`
}

func NewMonitorRequest() (request *MonitorRequest) {
	request = &MonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "Monitor")
	return
}

func NewMonitorResponse() (response *MonitorResponse) {
	response = &MonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}
func (c *Client) Monitor(request *MonitorRequest) (response *MonitorResponse, err error) {
	response = NewMonitorResponse()
	err = c.Send(request, response)
	return
}
