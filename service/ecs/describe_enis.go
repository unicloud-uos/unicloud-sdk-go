package ecs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DescribeEnisRequest struct {
	*tchttp.BaseRequest
	EniId      string `name:"EniId"`
	Size       string `name:"Size"`
	BindStatus string `name:"BindStatus"`
}

type DescribeEnisResponse struct {
	*tchttp.BaseResponse
	Enis       EniSet `json:"list"`
	Size       int    `json:"size"`
	Totalcount int    `json:"totalCount"`
	Totalpages int    `json:"totalPages"`
	Page       int    `json:"page"`
	Requestid  string `json:"RequestId"`
}

type EniSet []EniType

func (s *EniSet) UnmarshalJSON(data []byte) error {
	var enis = make([]EniType, 0)
	var err error
	if string(data) != "{}" {
		err = json.Unmarshal(data, &enis)
	}
	*s = enis
	return err
}

type EniType struct {
	Eipstatus    string `json:"eipStatus"`
	Azonename    string `json:"azoneName"`
	Subnetname   string `json:"subnetName"`
	Subnetcidr   string `json:"subnetCidr"`
	Ipv6Addr     string `json:"ipv6Addr"`
	Releasetype  int    `json:"releaseType"`
	Bandwidth    int    `json:"bandWidth"`
	Instanceid   string `json:"instanceId"`
	Instancename string `json:"instanceName"`
	Vpcname      string `json:"vpcName"`
	Subnetid     string `json:"subnetId"`
	Status       string `json:"status"`
	Ipv4Addr     string `json:"ipv4Addr"`
	Vpcid        string `json:"vpcId"`
	Eipcategory  string `json:"eipCategory"`
	Type         string `json:"type"`
	Eipid        string `json:"eipId"`
	Vpccidr      string `json:"vpcCidr"`
	Sglist       []Sg   `json:"sgList"`
	Eipip        string `json:"eipIp"`
	Eipname      string `json:"eipName"`
	Createtime   int64  `json:"createTime"`
	Azoneid      string `json:"azoneId"`
	Vmid         string `json:"vmId"`
	MAC          string `json:"mac"`
}

type Sg struct {
	Instanceid   string `json:"instanceId"`
	Instancename string `json:"instanceName"`
}

func (r *DescribeEnisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewDescribeEnisRequest() (request *DescribeEnisRequest) {
	request = &DescribeEnisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeEnis").SetHttpMethod(tchttp.GET)
	return
}

func NewDescribeEnisResponse() (response *DescribeEnisResponse) {
	response = &DescribeEnisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEnis(request *DescribeEnisRequest) (response *DescribeEnisResponse, err error) {
	response = NewDescribeEnisResponse()
	err = c.Send(request, response)
	return
}
