package slb

import (
	"fmt"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/errors"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/profile"
)

const (
	APIVersion = "20200730"
	Service    = "networks/slb"
)

type Client struct {
	common.Client
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

/**
 * 从环境变量中读取
 */
func NewClientFromEnv() (client *Client) {
	client = &Client{}
	client.InitFromEnv()
	return
}

/**
 * 是否是需要重试的error
 * UserQuotanotEnough 否
 */
func IsNoLangerRetryError(err error) (bool, string) {
	if err, ok := err.(*errors.UnicloudCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		if err.Code == UserQuotanotEnough {
			return false, UserQuotanotEnough
		}
	}
	return true, ""
}
