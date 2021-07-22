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
