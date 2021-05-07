package main

import (
	"fmt"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/errors"
	slb "github.com/unicloud-uos/unicloud-sdk-go/service/slb"
	"testing"
)

func TestListSLbEnv(t *testing.T) {

	client := slb.NewClientFromEnv()
	if client == nil {
		return
	}
	request := slb.NewDescribeLoadBalancersRequest()
	//request.Page = 1
	//request.Size = 10
	request.InstanceId = "slb-ks0ouqxkvk42"

	response, err := client.DescribeLoadBalancers(request)

	if _, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

func TestCreateSLb(t *testing.T) {

	client := slb.NewClientFromEnv()
	if client == nil {
		return
	}
	create := slb.CreateLoadBalancerArgs{
		AzoneId:           "HB1-BJMY2",
		ChargeType:        "postpaid",
		ComponentProperty: slb.ComponentProperty{SpecificationCode: "slb.s1.small"},
		InstanceCode:      "SLB",
		InstanceName:      "slbccm",
		OrderCategory:     "NEW",
		PayType:           "DAY_MONTH",
		ProductProperties: []slb.ProductProperties{
			{
				Address:  "",
				SubnetId: "subnet-h68ou3ylornpb",
				VpcId:    "vpc-g68ou3ylornpb",
			},
		},
		Quantity:  1,
		RenewType: "notrenew",
		RentCount: 1,
		RentUnit:  "month",
	}
	request := slb.NewCreateLoadBalancerRequest(create)

	response, err := client.CreateLoadBalancers(request)

	if _, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

func TestDeleteSLbEnv(t *testing.T) {

	client := slb.NewClientFromEnv()
	if client == nil {
		return
	}
	delete := slb.DeleteLoadBalancerArgs{
		Ids: []string{"slb-kpkbacv26uwb"},
	}
	request := slb.NewDeleteLoadBalancerRequest(delete)

	response, err := client.DeleteListener(request)

	if _, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())

}

func TestCreateListener(t *testing.T) {

	client := slb.NewClientFromEnv()
	if client == nil {
		return
	}
	create := slb.CreateListenerArgs{
		LoadbalancerId: "slb-ks0pbeecmqph",
		Description:    "create by go sdk",
		Name:           "k8s-listen",
		Protocol:       "TCP",
		ProtocolPort:   6443,
		UserId:         "0c3ade08-afb1-43bb-97b3-69ce7d7019c8",
	}
	request := slb.NewCreateListenerRequest(create)

	response, err := client.CreateListener(request)

	if _, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

func TestCreateServerGroup(t *testing.T) {

	client := slb.NewClientFromEnv()
	if client == nil {
		return
	}
	create := slb.CreateServerGroupArgs{
		LoadbalancerId: "slb-kui4j78dzu00",
		Description:    "create by go sdk",
		Name:           "k8s-servergroup",
		ListenerId:     "c5ecdb18735e4fccbff5bc3c2a62d026",
		Algorithm:      "ROUND_ROBIN",
		//SessionType: "",
		//CookieName: "",
		Servers: []slb.Servers{
			{
				PortId:     "eni-gSNp30763DSwJWFl4HkwHvwTnO11",
				ServerPort: 31515,
				ServerIp:   "172.16.0.55",
				Weight:     100,
				ServerId:   "ecs-DzLw8zbtSi4lc6VYMnZpNGhrVMKrScPR",
			},
			{
				PortId:     "eni-vZoMKXp1RHj5wRhp9jCCrGE2Atix",
				ServerPort: 31515,
				ServerIp:   "172.16.0.52",
				Weight:     100,
				ServerId:   "ecs-nOKicCPmlx3AfnrU3Q9sT4f6BHRRImEz",
			},
			{
				PortId:     "eni-PRIprXS2bhvbkGTkQGQA6d5mz6lM",
				ServerPort: 31515,
				ServerIp:   "172.16.0.53",
				Weight:     100,
				ServerId:   "ecs-q79N5lunjL9nK6bnh2x13yXf1uQ53Swu",
			},
		},
	}
	request := slb.NewCreateServerGroupRequest(create)

	response, err := client.CreateServerGroup(request)

	if _, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
