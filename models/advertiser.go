package models

type AdvertiserBudgetUpdateReq struct {
	AdvertiserId int64      `json:"advertiser_id,omitempty"`
	BudgetMode   BudgetMode `json:"budget_mode,omitempty"`
	Budget       float64    `json:"budget,omitempty"`
}

type AdvertiserBudgetUpdateRsp struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}

type AdvertiserDailyBudgetRsp struct {
	Code      int                          `json:"code,omitempty"`
	Message   string                       `json:"message,omitempty"`
	Data      AdvertiserDailyBudgetRspData `json:"data,omitempty"`
	RequestId string                       `json:"request_id,omitempty"`
}
type AdvertiserDailyBudgetRspDataList struct {
	AdvertiserId int64      `json:"advertiser_id,omitempty"`
	Budget       float64    `json:"budget,omitempty"`
	BudgetMode   BudgetMode `json:"budget_mode,omitempty"`
}
type AdvertiserDailyBudgetRspData struct {
	List []AdvertiserDailyBudgetRspDataList `json:"list,omitempty"`
}
