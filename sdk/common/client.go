package common

import (
	"fmt"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/errors"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/profile"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	RequestMethodGET  = "GET"
	RequestMethodPOST = "POST"

	SignatureMethodHMacSha256 = "HmacSHA256"
)

func (c *Client) WithCredential(cred *Credential) *Client {
	c.credential = cred
	return c
}

func (c *Client) WithProfile(clientProfile *profile.ClientProfile) *Client {
	c.profile = clientProfile
	c.signMethod = clientProfile.SignMethod
	c.unsignedPayload = clientProfile.UnsignedPayload
	c.httpProfile = clientProfile.HttpProfile
	c.debug = clientProfile.Debug
	c.httpClient.Timeout = time.Duration(c.httpProfile.ReqTimeout) * time.Second
	return c
}

func (c *Client) WithSecretId(secretId, secretKey string) *Client {
	c.credential = NewCredential(secretId, secretKey)
	return c
}

type Client struct {
	region          string
	httpClient      *http.Client
	httpProfile     *profile.HttpProfile
	profile         *profile.ClientProfile
	credential      *Credential
	signMethod      string
	unsignedPayload bool
	debug           bool
}

func (c *Client) Init(region string) *Client {
	c.httpClient = &http.Client{}
	c.region = region
	c.signMethod = "TC3-HMAC-SHA256"
	c.debug = false
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return c
}

func (c *Client) InitFromEnv() (cc *Client) {
	secretId := os.Getenv("UnicloudSecretId")
	secretKey := os.Getenv("UnicloudSecretKey")
	region := os.Getenv("UnicloudRegion")
	rootDomain := os.Getenv("UnicloudRootDomain")
	scheme := os.Getenv("Scheme")

	if region == "" {
		fmt.Printf("Fail to init client because env.UnicloudRegion is null ")
		return nil
	}

	if secretId == "" || secretKey == "" {
		fmt.Printf("Fail to init client because env.UnicloudSecretId = {%s}, env.UnicloudSecretKey= {%s} ", secretId, secretKey)
		return nil
	}
	credential := NewCredential(
		secretId,
		secretKey,
	)

	cpf := profile.NewClientProfile()
	if rootDomain != "" {
		cpf.HttpProfile.RootDomain = rootDomain
	}
	if scheme != "" {
		cpf.HttpProfile.Scheme = scheme
	}

	c.Init(region).WithCredential(credential).WithProfile(cpf)
	return c
}

func (c *Client) Send(request tchttp.Request, response tchttp.Response) (err error) {
	// 反射把request 中有tag的值 拼接到url
	tchttp.ConstructParams(request)

	if request.GetScheme() == "" {
		request.SetScheme(c.httpProfile.Scheme)
	}

	if request.GetRootDomain() == "" {
		request.SetRootDomain(c.httpProfile.RootDomain)
	}

	if request.GetDomain() == "" {
		domain := request.GetServiceDomain(request.GetService())
		request.SetDomain(domain)
	}

	if request.GetHttpMethod() == "" {
		request.SetHttpMethod(c.httpProfile.ReqMethod)
	}

	// 公共参数
	tchttp.CompleteCommonParams(request, c.GetRegion())

	return c.doSend(request, response)

}

func (c *Client) doSend(request tchttp.Request, response tchttp.Response) (err error) {

	request.GetParams()["AccessKeyId"] = c.credential.SecretId

	sign := GetSignature(request.GetHttpMethod(), request.GetRequestUrl(), c.credential.SecretKey)

	url := request.GetRequestUrl() + "&Signature=" + sign

	httpRequest, err := http.NewRequest(request.GetHttpMethod(), url, request.GetBodyData())
	if request.GetHttpMethod() == tchttp.POST || request.GetHttpMethod() == tchttp.PUT {
		httpRequest.Header.Set("Content-Type", "application/json;charset=UTF-8")
	}
	fmt.Printf("url: %s\n", url)
	fmt.Printf("Header: %+v\n", httpRequest.Header)
	if httpRequest.GetBody != nil {
		body, err := httpRequest.GetBody()
		if err == nil {
			buf, err := ioutil.ReadAll(body)
			if err == nil {
				fmt.Printf("Body: %v\n", string(buf))
			}
		}
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		msg := fmt.Sprintf("Fail to get response because %s", err)
		return errors.NewUnicloudCloudSDKError("ClientError.NetworkError", msg, "")
	}

	err = tchttp.ParseFromHttpResponse(httpResponse, response)
	return err
}

func (c *Client) GetRegion() string {
	return c.region
}
