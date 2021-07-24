package profile

import tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"

type HttpProfile struct {
	ReqMethod    string
	ReqTimeout   int
	SchemeDomain string
	// Deprecated, use Scheme instead
}

func NewHttpProfile() *HttpProfile {
	return &HttpProfile{
		ReqMethod:    tchttp.GET,
		ReqTimeout:   60,
		SchemeDomain: tchttp.SchemeDomain,
	}
}
