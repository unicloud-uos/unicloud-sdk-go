package cce

type BaseResponse struct {
	Status    bool   `json:"status"`
	Auth      bool   `json:"auth"`
	Msg       string `json:"msg"`
	Code      string `json:"code"`
	RequestId string `json:"RequestId"`
}

type Port struct {
	AdminStateUp   string    `json:"AdminStateUp"`
	AvailableZone  string    `json:"AvailableZone"`
	DeviceId       string    `json:"DeviceId"`
	DeviceOwner    string    `json:"DeviceOwner"`
	FixedIps       []FixedIp `json:"FixedIps"`
	ID             string    `json:"Id"`
	MacAddress     string    `json:"MacAddress"`
	OnlineType     string    `json:"OnlineType"`
	Status         string    `json:"Status"`
	SubnetCidr     string    `json:"SubnetCidr"`
	SubnetId       string    `json:"SubnetId"`
	TenantId       string    `json:"TenantId"`
	StandardAttrId string    `json:"standardAttrId"`
}

type FixedIp struct {
	Ipv4ID    string `json:"IPv4Id"`
	IpAddress string `json:"IpAddress"`
	Primary   string `json:"Primary"`
	SubnetId  string `json:"SubnetId"`
}
