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
	response, err := client.DescribeProductSpec("cce.p1.xlarge.8")
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
	response, err := client.PortDetach("16ab368d911945f3ab4a07badb452c11", "master-node-3301")
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
	response, err := client.PortDelete("c7fb8b66cb5547f7b0398bf7db4f17a0")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
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
