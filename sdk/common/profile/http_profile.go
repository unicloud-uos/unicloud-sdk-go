package profile

import tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"

type HttpProfile struct {
	ReqMethod  string
	ReqTimeout int
	Scheme     string
	RootDomain string
	// Deprecated, use Scheme instead
	Protocol string
}

func NewHttpProfile() *HttpProfile {
	return &HttpProfile{
		ReqMethod:  "GET",
		ReqTimeout: 60,
		Scheme:     "HTTPS",
		RootDomain: tchttp.RootDomain,
	}
}
