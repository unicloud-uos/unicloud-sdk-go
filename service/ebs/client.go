package ebs

import (
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/profile"

)

const APIVersion = "2"

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

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeDisks")
	return
}

func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDiskResponse, err error) {
	if request == nil {
		request = NewDescribeDisksRequest()
	}
	response = NewDescribeDiskResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) CreateDisks(request *CreateDisksRequest) (response *DescribeDiskResponse, err error) {
	if request == nil {
		request = NewCreateDisksRequest()
	}
	response = NewDescribeDiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskResponse() (response *DescribeDiskResponse) {
	response = &DescribeDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func NewCreateDisksRequest() (request *CreateDisksRequest) {
	request = &CreateDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "CreateDisk")
	return
}