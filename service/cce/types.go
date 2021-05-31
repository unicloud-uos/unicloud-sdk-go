package cce

type BaseResponse struct {
	Status bool   `json:"status"`
	Auth   bool   `json:"auth"`
	Msg    string `json:"msg"`
	Code   string `json:"code"`
}

type FixedIp struct {
	Ipv4ID    string `json:"IPv4Id"`
	IpAddress string `json:"IpAddress"`
	Primary   string `json:"Primary"`
	SubnetId  string `json:"SubnetId"`
}
