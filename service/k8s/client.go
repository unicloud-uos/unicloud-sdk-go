package k8s

import (
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/profile"
)

const (
	APIVersion = ""
	Service    = "cce"
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
