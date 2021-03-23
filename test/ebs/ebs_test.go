package ebs

import (
	"fmt"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/errors"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/profile"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/regions"
	"github.com/unicloud-uos/unicloud-sdk-go/service/ebs"
	"testing"
)

func TestListDisk(t *testing.T) {
	credential := common.NewCredential(
		"",
		"",
	)
	cpf := profile.NewClientProfile()
	client, _ := ebs.NewClient(credential, regions.HB1_BJMYB, cpf)

	request := ebs.NewDescribeDisksRequest()
	request.Page = 1
	request.Size = 10

	response, err := client.DescribeDisks(request)

	if _, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf(response.ToJsonString())
}

func TestCreateDisk(t *testing.T) {
	client := ebs.NewClientFromEnv()
	if client == nil{
		return
	}

	request := ebs.NewCreateDisksRequest()
	//request.AzId = []string{"HB1-BJMY2"}
	request.AzId = "HB1-BJMY2"
	request.PayType = "DAY_MONTH"
	request.Capacity = 20
	request.Quantity = 1
	request.SpecificationCode = "ebs.highIO.ssd"
	request.RentCount = 1
	request.RentUnit = "month"
   // ss := "{\"billingMethod\":\"DAY_MONTH\",\"conf\":{\"地域\":\"华北2-北京3 可用区1\",\"硬盘规格\":\"SSD云硬盘 20G\",\"计费模式\":\"按日月结\",\"资源组\":\"默认资源组\"},\"snapshotId\":\"\",\"vmId\":\"\",\"rgId\":\"rg-klepqvt6m4jk\"}"
	//request.Description = ss

	response, err := client.CreateDisks(request)

	if _, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf(response.ToJsonString())
}
