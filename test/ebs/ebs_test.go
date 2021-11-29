package ebs

import (
	"fmt"
	"github.com/unicloud-uos/unicloud-sdk-go/service/ebs"
	"testing"
)

func TestListDisk(t *testing.T) {
	client := ebs.NewClientFromEnv()
	request := ebs.NewDescribeDisksRequest()
	request.Page = 1
	request.Size = 10
	response, err := client.DescribeDisks(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDetailDisk(t *testing.T) {
	client := ebs.NewClientFromEnv()

	request := ebs.NewDetailDiskRequest()
	request.VolumeId = "ebs-c8xc51m7ow6u"
	response, err := client.DetailDisk(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	if len(response.Volume[0].AttachInfos) != 0 {
		fmt.Println("InstanceId: " + response.Volume[0].AttachInfos[0].InstanceId)
	}

	t.Log(response.ToJsonString())
}

func TestCreateDisk(t *testing.T) {
	client := ebs.NewClientFromEnv()
	request := ebs.NewCreateDiskRequest()
	request.AzId = "cn-tianjin-yfb1"
	request.PayType = "CHARGING_HOURS"
	request.SpecificationCode = "ebs.highIO.ssd"
	request.Capacity = 20
	request.Quantity = 1

	//request := ebs.NewCreateDiskRequest()
	//request.AzId = "HB1-BJMY2"
	//request.PayType = "CHARGING_HOURS"
	//request.SpecificationCode = "ebs.highIO.ssd"
	//request.Capacity = 20
	//request.Quantity = 1
	//request.RentCount = 1

	response, err := client.CreateDisks(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDeleteDisk(t *testing.T) {
	client := ebs.NewClientFromEnv()
	request := ebs.NewDeleteDiskRequest()
	request.VolumeIds = "ebs-c8xc51m7ow6u"
	response, err := client.DeleteDisk(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}
