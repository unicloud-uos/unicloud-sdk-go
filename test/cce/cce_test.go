package cce

import (
	"github.com/unicloud-uos/unicloud-sdk-go/service/cce"
	"testing"
)

func TestQueryPort(t *testing.T) {
	client := cce.NewClientFromEnv()
	response, err := client.QueryPortsByDeviceId("ecs-ku7ycunzaniu")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}

func TestDescribeProductSpec(t *testing.T) {
	client := cce.NewClientFromEnv()
	response, err := client.DescribeProductSpec("cce.p1.xlarge.8")
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}
