package vpc

import (
	"github.com/unicloud-uos/unicloud-sdk-go/service/vpc"
	"testing"
)

func TestDescribeSubnets(t *testing.T) {
	client := vpc.NewClientFromEnv()
	request := vpc.NewDescribeSubnetsRequest()
	request.VpcId = "vpc-7iamipow2cupb"
	response, err := client.DescribeSubnets(request)
	if err != nil {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	t.Log(response.ToJsonString())
}
