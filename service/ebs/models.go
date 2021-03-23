package ebs

import (
	"encoding/json"
	tchttp "github.com/unicloud-uos/unicloud-sdk-go/sdk/common/http"
)

type DescribeDisksRequest struct {
	*tchttp.BaseRequest

	Size int `name:"Size" int`

	Page int `name:"Page" int`

	InstanceId string `name:"InstanceId" string`
}


type DescribeDiskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足过滤条件的负载均衡实例总数。此数值与入参中的Limit无关。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
	Current int `json:"current`
	Size int `json:"size`
	Pages int `json:"pages`
	Total int `json:"total`
	Records []DescribeDisk `json:"records`
	RequestId string `json:"RequestId"`
}

type DescribeDisk struct {
	ChargeType int `json:"chargeType"`
	DiskId string `json:"diskId"`
}


func (r *DescribeDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest

	// cn-beijing-a cn-shanghai-a
	//AzId []string `name:"Tags" list`
	AzId string `name:"AzId"`

	//计费类型。YEAR_MONTH（包年包月）；DAY_MONTH（按日月结）；CHARGING_HOURS（按小时实时付费）
	PayType string `name:"PayType"`

	// 硬盘类型。ebs.highIO.ssd（SSD云硬盘）；ebs.hybrid.hdd（高性能HDD云硬盘）
	SpecificationCode string `name:"SpecificationCode"`

	// 硬盘大小。可选范围：20~32768 GB
	Capacity int `name:"Capacity"`

	// 硬盘数量。可选范围：1~15
	Quantity int `name:"Quantity"`

	// 租期。预付费必传，可选范围：1个月、2个月、3个月、6个月、1年
	RentCount int `name:"RentCount"`

	// 租期单位。预付费必传，可选范围：day, month, year
	RentUnit string `name:"RentUnit"`

	// 硬盘描述。
	Description string `name:"Description"`

	// 用快照创建硬盘时快照ID。
	SnapshotId string `name:"SnapshotId"`

	// 主机ID。创建包年包月硬盘必传该参数。
	VmId string `name:"VmId"`

	//// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘<br><li>CLOUD_HSSD：表示增强型SSD云硬盘<br><li>CLOUD_TSSD：表示极速型SSD云硬盘。
	//DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	//
	//// 云硬盘计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDCPAID：独享集群付费<br>各类型价格请参考云硬盘[价格总览](/document/product/362/2413)。
	//DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	//
	//// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。
	//Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	//
	//// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。
	//DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	//
	//// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制，具体参见[云硬盘使用限制](https://cloud.tencent.com/doc/product/362/5145)。
	//DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
	//
	//// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。<br>创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	//DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`
	//
	//// 云硬盘大小，单位为GB。<br><li>如果传入`SnapshotId`则可不传`DiskSize`，此时新建云盘的大小为快照大小<br><li>如果传入`SnapshotId`同时传入`DiskSize`，则云盘大小必须大于或等于快照大小<br><li>云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	//DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	//
	//// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过[DescribeSnapshots](/document/product/362/15647)接口查询快照，见输出参数DiskUsage解释。
	//SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	//
	//// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	//ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	//
	//// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。
	//Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	//
	//// 云盘绑定的标签。
	//Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
	//
	//// 可选参数，默认为False。传入True时，云盘将创建为共享型云盘。
	//Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`
	//
	//// 可选参数。使用此参数可给云硬盘购买额外的性能。<br>当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）
	//ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}