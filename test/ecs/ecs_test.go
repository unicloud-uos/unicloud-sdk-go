package ecs

import (
	"github.com/unicloud-uos/unicloud-sdk-go/service/ecs"
	"testing"
)

func TestAttachDisk(t *testing.T) {
	client := ecs.NewClientFromEnv()
	request := ecs.NewAttachDiskRequest()
	request.InstanceId = "ecs-5bnnsw91o5ee498p4n2j"
	request.DiskId = "ebs-kti4u2h8kqum"
	response, err := client.AttachDisk(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDetachDisk(t *testing.T) {
	client := ecs.NewClientFromEnv()
	request := ecs.NewDetachDiskRequest()
	request.InstanceId = "ecs-knn1jv1uic2r"
	request.DiskId = "ebs-kquyqu1ayv2h"
	response, err := client.DetachDisk(request)

	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestGetPassword(t *testing.T) {
	client := ebs.NewClientFromEnv()
	request := ebs.NewGetPasswordRequest()
	request.InstanceId = "ecs-kp2kky0k9z8g"
	response, err := client.GetPassword(request)

	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response)
}

func TestMonitor(t *testing.T) {
	client := ebs.NewClientFromEnv()
	request := ebs.NewMonitorRequest()
	request.InstanceId = "ecs-kp2kky0k9z8g"
	request.DateCategory = ""
	response, err := client.Monitor(request)

	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response)
}

func TestDescribeEcs(t *testing.T) {
	client := ebs.NewClientFromEnv()
	request := ebs.NewDescribeEcsRequest()
	request.InstanceId = "ecs-kp2kky0k9z8g"
	request.Page = 1
	response, err := client.DescribeEcs(request)

	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response)
}

func TestDetailEcs(t *testing.T) {
	client := ebs.NewClientFromEnv()
	request := ebs.NewDetailEcsRequest()
	request.InstanceId = "ecs-kp2kky0k9z8g"
	response, err := client.DetailEcs(request)

	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response)
}

func TestRunEcs(t *testing.T) {
	client := ebs.NewClientFromEnv()
	run := ebs.Body{
		AzoneId:         "HB1-BJMY2",
		Description:     "",
		HostName:        "i-F8D3QNsoC5",
		ImageId:         "CentOS_8_0_64bit_Minimal_std",
		InitialShell:    "",
		InstanceName:    "i-gbX287u92H",
		KeyPair:         "",
		PayType:         "DAY_MONTH",
		SecurityGroupId: "sg-a1xy0gvuvlspb",
		SysDiskSize:     40,
		VpcId:           "vpc-g68ou3ylornpb",
	}
	request := ebs.NewRunEcsRequest(run)
	response, err := client.RunEcs(request)

	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response)
}
