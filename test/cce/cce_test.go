package cce

import (
	"encoding/hex"
	"github.com/unicloud-uos/unicloud-sdk-go/service/cce"
	"math/rand"
	"testing"
	"time"
)

func TestQueryPort(t *testing.T) {
	client := cce.NewClientFromEnv()
	ports, err := client.QueryPortsByDeviceId("ecs-ku7ycunzaniu")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Logf("spec:%+v", ports)
}

func TestQueryPortByEniId(t *testing.T) {
	client := cce.NewClientFromEnv()
	ports, err := client.QueryPortByEniId("da574b5e6c92486f8b57263147e35739")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Logf("spec:%+v", ports)
}

func TestDescribeProductSpec(t *testing.T) {
	client := cce.NewClientFromEnv()
	response, err := client.DescribeProductSpec("cce.large.4")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Logf("spec:%+v", response)
}

func TestPortAttach(t *testing.T) {
	client := cce.NewClientFromEnv()
	response, err := client.PortAttach("62956ff48061404e8d4ebf95cf370a30", "cn-tianjin-yfb1", "cce-kwz8p9lv6oc3-worker-0")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestPortDetach(t *testing.T) {
	client := cce.NewClientFromEnv()
	response, err := client.PortDetach("0d107a590cce42daadabc035398ac590", "cce-kxrqiw0r7b2o-worker-0")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestAssignIp(t *testing.T) {
	client := cce.NewClientFromEnv()
	ips, err := client.PortAssignIp("472ad61320a54d429ac655c09324967a", 1)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Logf("ips:%+v", ips)
}

func TestUnAssiginIp(t *testing.T) {
	client := cce.NewClientFromEnv()
	response, err := client.PortUnAssignIp("472ad61320a54d429ac655c09324967a", []string{""})
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestPortCreate(t *testing.T) {
	client := cce.NewClientFromEnv()
	eniName := generateEniName()
	createArgs := cce.PortCreateArgs{
		AvailableZone: "cn-tianjin-yfb1", Name: eniName,
		SecurityGroups: []string{"5651bf77-210a-4aec-94a3-ece189712d63"},
		SubnetId:       "b934aea2538e49a0b003cc9bf7972d1a", AssistIpCount: 0}
	port, err := client.PortCreate(createArgs)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Logf("port:%+v", port)
}

func TestPortDelete(t *testing.T) {
	client := cce.NewClientFromEnv()
	response, err := client.PortDelete("64382fbea3d743a59fb04d580f57e3f5")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestPortBind(t *testing.T) {
	client := cce.NewClientFromEnv()
	eniName := generateEniName()
	createArgs := cce.PortCreateArgs{
		AvailableZone: "cn-tianjin-yfb1", Name: eniName,
		SecurityGroups: []string{"4505486d-ecc7-423c-a5e8-1fde438b7da3"},
		SubnetId:       "137d895cda7e4781b2f9b7159eb1aea2", AssistIpCount: 0}
	port, err := client.PortCreate(createArgs)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Logf("create port:%+v", port)
	_, err = client.PortBind(port.ID, "4505486d-ecc7-423c-a5e8-1fde438b7da3")
	if err != nil {
		t.Fatalf("An API error has returned: %s", err)
	}
}

func generateEniName() string {
	b := make([]byte, 3)
	rand.Seed(time.Now().UnixNano())
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return "eni-cni-" + hex.EncodeToString(b)
}

func TestAttachDisk(t *testing.T) {
	client := cce.NewClientFromEnv()
	request := cce.NewAttachDiskRequest()
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
	client := cce.NewClientFromEnv()
	request := cce.NewDetachDiskRequest()
	request.InstanceId = "cce-k6ebackvtcp0"
	request.NodeId = "cce-gb0k3oyb7vf-331-worker-3"
	request.DomainId = "ecs-90pck7m035h1z4n5wp36"
	request.DiskId = "ebs-k6evw8kdqaih"
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
	client := cce.NewClientFromEnv()

	request := cce.NewDetailDiskRequest()
	request.DiskId = "ebs-k14t61c5la6m"
	response, err := client.DetailDisk(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}

	t.Log(response.ToJsonString())
}
