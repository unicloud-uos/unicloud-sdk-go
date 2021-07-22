package ecs

import (
	"github.com/unicloud-uos/unicloud-sdk-go/service/k8s"
	"testing"
)

func TestAttachDisk(t *testing.T) {
	client := k8s.NewClientFromEnv()
	request := k8s.NewAttachDiskRequest()
	request.InstanceId = "cce-gb0k3oyb7vf-331"
	request.NodeId = "cce-gb0k3oyb7vf-331-worker-3"
	request.DomainId = "bms-1t1aq5508lx85i29s85k"
	request.DiskId = "ebs-kx20uku7x8yo"
	response, err := client.AttachDisk(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	if !response.Status {
		t.Errorf("An API error has returned: %s", response.Msg)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDetachDisk(t *testing.T) {
	client := k8s.NewClientFromEnv()
	request := k8s.NewDetachDiskRequest()
	request.InstanceId = "cce-gb0k3oyb7vf-331"
	request.NodeId = "cce-gb0k3oyb7vf-331-worker-3"
	request.DomainId = "bms-1t1aq5508lx85i29s85k"
	request.DiskId = "ebs-kx2skssuymlf"
	response, err := client.DetachDisk(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	if !response.Status {
		t.Errorf("An API error has returned: %s", response.Msg)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDetailDisk(t *testing.T) {
	client := k8s.NewClientFromEnv()

	request := k8s.NewDetailDiskRequest()
	request.DiskId = "ebs-k14t61c5la6m"
	response, err := client.DetailDisk(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}

	t.Log(response.ToJsonString())
}
