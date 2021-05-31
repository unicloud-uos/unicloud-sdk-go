package ecs

import (
	"encoding/hex"
	"encoding/json"
	"github.com/unicloud-uos/unicloud-sdk-go/service/ecs"
	"math/rand"
	"testing"
	"time"
)

func TestAttachDisk(t *testing.T) {
	client := ecs.NewClientFromEnv()
	request := ecs.NewAttachDiskRequest()
	request.InstanceId = "ecs-knn1jv1uic2r"
	request.DiskId = "ecs-kquyqu1ayv2h"
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
	request.DiskId = "ecs-kquyqu1ayv2h"
	response, err := client.DetachDisk(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDescribeEcs(t *testing.T) {
	client := ecs.NewClientFromEnv()
	request := ecs.NewDescribeEcsRequest()

	response, err := client.DescribeEcs(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDescribeEnis(t *testing.T) {
	client := ecs.NewClientFromEnv()
	request := ecs.NewDescribeEnisRequest()
	request.EniId = "eni-kwdym3jlfanb"
	response, err := client.DescribeEnis(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestUnmarshal(t *testing.T) {
	resp := []byte(`{"list":[{"eipStatus":null,"azoneName":"可用区Y","subnetName":"PodNet","ipv6Addr":"","releaseType":1,"bandWidth":null,"subnetCidr":"172.28.0.0\/14","instanceId":"eni-ku7xyjsf9e31","instanceName":"eni-cni-163935","vpcName":"vpc0409","subnetId":"caebe90ec4534f7280b79b93f5e32a7b","status":"RUNNING","ipv4Addr":"172.28.0.2","vpcId":"vpc-7iamipow2cupb","eipCategory":null,"type":"secondary","eipId":null,"vpcCidr":"172.16.0.0\/12","sgList":[{"instanceId":"sg-266p2ibq5w2ob","instanceName":"default"}],"eipIp":null,"eipName":null,"createTime":1620894471000,"azoneId":"cn-tianjin-yfb1","vmId":null}],"size":10,"totalCount":1,"totalPages":1,"page":1,"RequestId":"f6677c81-d115-4ff8-94d1-2a21e196488e"}`)
	result := ecs.NewDescribeEnisResponse()
	err := json.Unmarshal(resp, &result)
	if err != nil {
		t.Log("err:", err)
	}
	t.Log(result.ToJsonString())
	empty := ecs.NewDescribeEnisResponse()
	empty.Enis = ecs.EniSet{}
	t.Log(empty.ToJsonString())
}

func TestCreateEni(t *testing.T) {
	client := ecs.NewClientFromEnv()
	createArgs := ecs.CreateEniArgs{
		AzoneId:     "cn-tianjin-yfb1",
		Name:        generateEniName(),
		VpcId:       "vpc-7iamipow2cupb",
		SubnetId:    "caebe90ec4534f7280b79b93f5e32a7b",
		SgId:        "sg-266p2ibq5w2ob",
		ReleaseType: 1,
		RegionId:    client.GetRegion(),
	}
	request := ecs.NewCreateEniRequest(createArgs)
	response, err := client.CreateEni(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDeleteEni(t *testing.T) {
	client := ecs.NewClientFromEnv()
	response, err := client.DeleteEni("eni-kwdnpz24vmps")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestAttachEni(t *testing.T) {
	client := ecs.NewClientFromEnv()
	response, err := client.AttachEni("eni-ku7xyjsf9e31", "ecs-ku69cesc0a02")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDettachEni(t *testing.T) {
	client := ecs.NewClientFromEnv()
	response, err := client.DetachEni("eni-kv4apj7h91l7")
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
