package models

import "github.com/antihax/optional"

// 广告组添加请求结构
type AdCampaignAddReq struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id,omitempty"`
	// 广告组名称，长度为1-100个字符，其中1个中文字符算2位
	CampaignName string `json:"campaign_name,omitempty"`
	// 广告组状态  允许值: "enable","disable"默认值：enable开启状态
	Operation string `json:"operation,omitempty"`
	//  广告组预算类型, 允许值: "BUDGET_MODE_INFINITE","BUDGET_MODE_DAY"
	BudgetMode BudgetMode `json:"budget_mode,omitempty"`
	// 广告组预算，取值范围: ≥ 0 当budget_mode为"BUDGET_MODE_DAY"时,必填,且日预算不少于300元
	Budget int64 `json:"budget,omitempty"`
	// 广告组推广目的, (创建后不可修改)
	//    允许值: "LINK","APP","DPA","GOODS","STORE","SHOP","AWEME"
	LandingType LandingType `json:"landing_type,omitempty"`
	// 广告组商品类型 (创建后不可修改)
	// 允许值:
	//    CAMPAIGN_DPA_DEFAULT_NOT: 非 DPA
	//    CAMPAIGN_DPA_SINGLE_DELIVERY:SDPA 单商品推广
	//    CAMPAIGN_DPA_MULTI_DELIVERY: DPA 商品推广
	// 默认值：
	//    推广目的非 DPA 时默认 CAMPAIGN_DPA_DEFAULT_NOT
	//    推广目的为 DPA 时默认 CAMPAIGN_DPA_MULTI_DELIVERY
	DeliveryRelatedNum string `json:"delivery_related_num,omitempty"`
}
type AdCampaignGetOpts struct {
	Filtering optional.Interface `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	Fields    optional.Interface `json:"fields,omitempty"`    // 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典
	Page      optional.Int64     `json:"page,omitempty"`      // 当前页码: 1
	PageSize  optional.Int64     `json:"page_size,omitempty"` // 页面大小 默认值: 10， 取值范围：1-1000
}

// AdCampaignGetOpts 过滤条件
type AdCampaignGetOptsFiltering struct {
	Ids                []int64        `json:"ids,omitempty"`                  // 广告组ID过滤，数组，不超过100个
	CampaignName       string         `json:"campaign_name,omitempty"`        // 广告组name过滤，长度为1-30个字符，其中1个中文字符算2位
	LandingType        LandingType    `json:"landing_type,omitempty"`         // 广告组推广目的过滤
	Status             CampaignStatus `json:"status,omitempty"`               // 广告组状态过滤，默认为返回“所有不包含已删除”
	CampaignCreateTime string         `json:"campaign_create_time,omitempty"` // 广告组创建时间，格式yyyy-mm-dd，表示过滤出当天创建的广告组
}

type AdCampaignGetRspData struct {
	List     []AdCampaignGetRspDataList `json:"list,omitempty"`
	PageInfo PageInfo                   `json:"page_info,omitempty"`
}

type AdCampaignGetRspDataList struct {
	Id                 int64  `json:"id,omitempty"`                   // 广告组ID
	Name               string `json:"name,omitempty"`                 // 广告组名称
	Budget             string `json:"budget,omitempty"`               // 广告组预算
	BudgetMode         string `json:"budget_mode,omitempty"`          // 广告组预算类型
	LandingType        string `json:"landing_type,omitempty"`         // 广告组推广目的
	ModifyTime         string `json:"modify_time,omitempty"`          // 广告组时间戳,用于更新时提交,服务端判断是否基于最新信息修改
	Status             string `json:"status,omitempty"`               // 广告组状态,详见
	CampaignCreateTime string `json:"campaign_create_time,omitempty"` // 广告组创建时间, 格式：yyyy-mm-dd hh:MM:ss
	CampaignModifyTime string `json:"campaign_modify_time,omitempty"` // 广告组修改时间, 格式：yyyy-mm-dd hh:MM:ss
	DeliveryRelatedNum string `json:"delivery_related_num,omitempty"` // 广告组商品类型
	DeliveryMode       string `json:"delivery_mode,omitempty"`        // 投放类型，允许值：MANUAL（手动）、PROCEDURAL（自动，投放管家
}

type AdCampaignGetRsp struct {
	Code      int                  `json:"code,omitempty"`
	Message   string               `json:"message,omitempty"`
	Data      AdCampaignGetRspData `json:"data,omitempty"`
	RequestId string               `json:"request_id,omitempty"`
}

// 广告组添加响应结构
type AdCampaignAddRsp struct {
	Code      int                  `json:"code,omitempty"`
	Message   string               `json:"message,omitempty"`
	Data      AdCampaignAddRspData `json:"data,omitempty"`
	RequestId string               `json:"request_id,omitempty"`
}

// 广告组添加响应数据
type AdCampaignAddRspData struct {
	// 广告组id
	CampaignId int64 `json:"campaign_id,omitempty"`
}
