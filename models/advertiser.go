package models

type AdvertiserBudgetUpdateReq struct {
	AdvertiserId int64      `json:"advertiser_id"`
	BudgetMode   BudgetMode `json:"budget_mode"`
	Budget       float64    `json:"budget,omitempty"`
}

type AdvertiserBudgetUpdateRsp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

type AdvertiserDailyBudgetRsp struct {
	Code      int                          `json:"code"`
	Message   string                       `json:"message"`
	Data      AdvertiserDailyBudgetRspData `json:"data"`
	RequestId string                       `json:"request_id"`
}
type AdvertiserDailyBudgetRspDataList struct {
	AdvertiserId int64      `json:"advertiser_id"`
	Budget       float64    `json:"budget"`
	BudgetMode   BudgetMode `json:"budget_mode"`
}
type AdvertiserDailyBudgetRspData struct {
	List []AdvertiserDailyBudgetRspDataList `json:"list"`
}
