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

	response, err := client.DescribeLoadBalancers(request)

	if _, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Print(response.Res.Total)
	fmt.Printf("%s", response.ToJsonString())
}

func TestCreateSLb(t *testing.T) {

	client := slb.NewClientFromEnv()
	if client == nil{
		return
	}
	create := slb.CreateLoadBalancerArgs{
		AzoneId: "HB1-BJMY2",
		ChargeType: "postpaid",
		ComponentProperty: slb.ComponentProperty{SpecificationCode: "slb.s1.small"},
		InstanceCode: "SLB",
		InstanceName: "slbccm",
		OrderCategory: "NEW",
		PayType: "DAY_MONTH",
		ProductProperties: []slb.ProductProperties{
			{
				Address:"",
				SubnetId:"subnet-h68ou3ylornpb",
				VpcId:"vpc-g68ou3ylornpb",
			},
		},
		Quantity:1,
		RenewType: "notrenew",
		RentCount: 1,
		RentUnit: "month",
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
