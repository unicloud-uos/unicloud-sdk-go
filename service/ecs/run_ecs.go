package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type RunEcsRequest struct {
	*tchttp.BaseRequest
	Body Body
}
type Body struct {
	RegionId                    string `name:"RegionId"`
	AzoneId                     string `name:"azoneId"`
	PayType                     string `name:"payType"`
	Period                      int    `name:"period"`
	VmSpecificationCode         string `name:"vmSpecificationCode"`
	SysDiskSpecificationCode    string `name:"sysDiskSpecificationCode"`
	SysDiskSize                 int    `name:"sysDiskSize"`
	ImageId                     string `name:"imageId" `
	ImageSpecificationClassCode string `name:"imageId" `
	VpcId                       string `name:"vpcId"`
	SecurityGroupId             string `name:"securityGroupId"`
	BandWidthSpecificationCode  string `name:"bandWidthSpecificationCode"`
	BandWidthSize               int    `name:"bandWidthSize"`
	Password                    string `name:"password"`
	KeyPair                     string `name:"keyPair"`
	InstanceName                string `name:"instanceName"`
	HostName                    string `name:"hostName"`
	Description                 string `name:"description"`
	BaseQuantity                int    `name:"baseQuantity"`
	MasterEniSubNetId           string `name:"masterEniSubNetId"`
	MasterEniName               string `name:"masterEniName"`
	MasterEniIp                 string `name:"masterEniIp"`
	SecondaryEniSubNetId        string `name:"secondaryEniSubNetId"`
	SecondaryEniName            string `name:"secondaryEniName"`
	InitialShell                string `name:"initialShell"`

	DataDisks struct {
		DataDiskSpecificationCode string `name:"dataDiskSpecificationCode"`
		DataDiskSize              int    `name:"dataDiskSize"`
		TemplateId                string `name:"templateId"`
	} `name:"body"`
}

type RunEcsResponse struct {
	*tchttp.BaseResponse
	RequestId   string   `name:"RequestId" string`
	OrderId     string   `name:"orderId" string`
	InstanceIds []string `name:"instanceIds" []string`
}

func (r *RunEcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func NewRunEcsRequest(body Body) (request *RunEcsRequest) {
	request = &RunEcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
		Body:        body,
	}
	request.Init().WithApiInfo(Service, APIVersion, "RunEcs").SetHttpMethod(tchttp.POST)
	request.SetBodyData(body)
	return
}

func NewRunEcsResponse() (response *RunEcsResponse) {
	response = &RunEcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RunEcs(request *RunEcsRequest) (response *RunEcsResponse, err error) {
	response = NewRunEcsResponse()
	err = c.Send(request, response)
	return
}
