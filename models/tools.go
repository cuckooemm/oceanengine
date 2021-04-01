package models

import "github.com/antihax/optional"

type ToolsIndustryGetOpts struct {
	Level optional.Int    `json:"level,omitempty"` // 只获取某级别数据，1:第一级,2:第二级,3:第三级，默认都返回
	Type  optional.String `json:"type,omitempty"`  // 可选值："ADVERTISER"，"AGENT"，"ADVERTISER"为原有广告3.0行业, "AGENT"为代理商行业获取，代理商行业level都为1
}

type ToolsIndustryGetRsp struct {
	Code      int                     `json:"code,omitempty"`
	Message   string                  `json:"message,omitempty"`
	Data      ToolsIndustryGetRspData `json:"data,omitempty"`
	RequestId string                  `json:"request_id,omitempty"`
}

type ToolsIndustryGetRspData struct {
	List []ToolsIndustryGetRspDataList `json:"list,omitempty"`
}

type ToolsIndustryGetRspDataList struct {
	IndustryId         int64  `json:"industry_id,omitempty"`          // 行业ID
	IndustryName       string `json:"industry_name,omitempty"`        // 行业名称
	Level              int    `json:"level,omitempty"`                // 所在级别，1：一级行业、2：二级行业、3：三级行业
	FirstIndustryId    int64  `json:"first_industry_id,omitempty"`    // 该行业的一级行业ID
	FirstIndustryName  string `json:"first_industry_name,omitempty"`  // 该行业的一级行业名称
	SecondIndustryId   int64  `json:"second_industry_id,omitempty"`   // 该行业的二级行业ID
	SecondIndustryName string `json:"second_industry_name,omitempty"` // 该行业的二级行业名称
	ThirdIndustryId    int64  `json:"third_industry_id,omitempty"`    // 该行业的三级行业ID
	ThirdIndustryName  string `json:"third_industry_name,omitempty"`  // 该行业的三级行业名称
}
type ToolsActionTextGetOpts struct {
	AdvancedCreativeType optional.String `json:"advanced_creative_type,omitempty"` // 附加创意类型
	Industry             optional.Int64  `json:"industry,omitempty"`               // 广告主行业id
}
type ToolsActionTextGetRsp struct {
	Code      int      `json:"code,omitempty"`
	Message   string   `json:"message,omitempty"`
	Data      []string `json:"data,omitempty"`
	RequestId string   `json:"request_id,omitempty"`
}

type ToolsADStatExtraInfoRsp struct {
	Code      int                           `json:"code,omitempty"`
	Message   string                        `json:"message,omitempty"`
	Data      []ToolsAdStatExtraInfoRspData `json:"data,omitempty"`
	RequestId string                        `json:"request_id,omitempty"`
}

type ToolsAdStatExtraInfoRspData struct {
	AdId          int64         `json:"ad_id,omitempty"`
	LearningPhase LearningPhase `json:"learning_phase,omitempty"`
}

type ToolsAdQualityRsp struct {
	Code      int                   `json:"code,omitempty"`
	Message   string                `json:"message,omitempty"`
	Data      ToolsAdQualityRspData `json:"data,omitempty"`
	RequestId string                `json:"request_id,omitempty"`
}

type ToolsAdQualityRspDataList struct {
	AppId        int64   `json:"app_id,omitempty"`        // app的id，可在应用下载广告中使用
	QualityScore float64 `json:"quality_score,omitempty"` // 计划综合质量得分
	CtrScore     float64 `json:"ctr_score,omitempty"`     // 创意质量得分
	WebScore     float64 `json:"web_score,omitempty"`     // 落地页响应得分
	CvrScore     float64 `json:"cvr_score,omitempty"`     // 落地页素材得分
}
type ToolsAdQualityRspData struct {
	List []ToolsAdQualityRspDataList `json:"list,omitempty"`
}

type ToolsAdDiagnosisRsp struct {
	Code      int                     `json:"code,omitempty"`
	Message   string                  `json:"message,omitempty"`
	Data      ToolsAdDiagnosisRspData `json:"data,omitempty"`
	RequestId string                  `json:"request_id,omitempty"`
}

type ToolsAdDiagnosisRspData struct {
	List []ToolsAdDiagnosisRspDataList `json:"list,omitempty"`
}
type ToolsAdDiagnosisRspDataList struct {
	Id            int64                      `json:"id,omitempty"`              // 兴趣词包id
	Severity      float64                    `json:"severity,omitempty"`        // attributions.severity的均值，0~1表示非常好到非常差
	Conclusion    string                     `json:"conclusion,omitempty"`      // 归因的主要结论。显示在列表页的“诊断结果”
	TimeRange     string                     `json:"time_range,omitempty"`      // 时间范围 eg. 2019-04-09 08:08 ~ 10:08
	TargetQuality string                     `json:"target_quality,omitempty"`  // 广告定向质量度（覆盖度，精准度，蓝海度），eg. {"aud_num": 0.311046511627907, "datetime": "2019-03-02 00:46:54.463356", "precision": 0.3136427566807314, "blueocean": 0.030942334739803123}
	FutureStarTag bool                       `json:"future_star_tag,omitempty"` // 标识是否为潜力广告 true为潜力广告
	Issue         string                     `json:"issue,omitempty"`           // 说明，eg. 该计划今天的展示量环比昨天下降73%
	Reason        string                     `json:"reason,omitempty"`          // 投放问题原因
	Suggest       string                     `json:"suggest,omitempty"`         // 投放建议
	ColdStart     float64                    `json:"cold_start,omitempty"`      // 冷启动预估分数
	SceneTag      string                     `json:"scene_tag,omitempty"`       // 投放问题场景标签 取值："future_star(潜力计划)","drop(掉量计划)"
	ModifyOption  string                     `json:"modify_option,omitempty"`   // 建议采纳按键的操作 取值："一键撤销","修改预算","快速优化","优化落地页","修改出价","修改定向"
	OriginValue   float64                    `json:"origin_value,omitempty"`    // 修改项的原始值
	SuggestValue  float64                    `json:"suggest_value,omitempty"`   // 修改项的建议修改值
	FunnelData    ToolsAdDiagnosisFunnelData `json:"funnel_data,omitempty"`     // 投放漏斗数据

}

type ToolsAdDiagnosisFunnelData struct {
	Target       ToolsAdDiagnosisFunnelDataRank         `json:"target,omitempty"` // 投放漏斗数据，定向阶段
	Rank         ToolsAdDiagnosisFunnelDataRank         `json:"rank,omitempty"`   // 投放漏斗数据，竞价阶段，字段同target
	Attributions ToolsAdDiagnosisFunnelDataAttributions `json:"attributions,omitempty"`
	RitResult    ToolsAdDiagnosisFunnelDataRitResult    `json:"rit_result,omitempty"` // 分投放位置的诊断结果集，键为广告位(INVENTORY_FEED/INVENTORY_HOTSOON_FEED/INVENTORY_AWEME_FEED/INVENTORY_VIDEO_FEED)，值见下表
}

type ToolsAdDiagnosisFunnelDataRank struct {
	SortRank float64 `json:"sort_rank,omitempty"` // 漏斗通过率排序
	Bench    float64 `json:"bench,omitempty"`     // 标杆漏斗通过率值
	Aim      float64 `json:"aim,omitempty"`       // 该广告计划漏斗通过率值
}

type ToolsAdDiagnosisFunnelDataAttributions struct {
	Name            string  `json:"name,omitempty"`              // 投放设置名称，eg. 落地页
	Severity        float64 `json:"severity,omitempty"`          // 0~1表示非常好到非常差
	Conclusion      string  `json:"conclusion,omitempty"`        // 诊断结论 eg. 转化率正常
	CardConclusion  string  `json:"card_conclusion,omitempty"`   // 卡片诊断结论 eg.一般
	CellValue       string  `json:"cell_value,omitempty"`        // 诊断计划的投放要素取值 eg. 2.4%
	BenchValue      string  `json:"bench_value,omitempty"`       // 比较对象的投放要素取值 eg. 3.5%
	LandingPageInfo string  `json:"landing_page_info,omitempty"` // 落地页洞察数据 eg. {"data":{"site":{"stats":{"loadTimeRank":0.632539,"loadTime":1691,"bouncedRateRank":1.000000,"bouncedRate":0.500000,"visitTimeRank":0.229752,"visitTime":10139,"exposureRateRank":0.272873,"exposureRate":51.000000}}}}
	FutureStarRates []struct {
		RateList []float64 `json:"rate_list,omitempty"` // 调价比例
		BidList  []float64 `json:"bid_list,omitempty"`  // 建议出价
		CostList []float64 `json:"cost_list,omitempty"` // cost_list为预估展示提升

	} `json:"future_star_rates,omitempty"` // 潜力广告的预估值，非潜力广告取值为空
	ModifyOption string  `json:"modify_option,omitempty"` // 建议采纳按键的操作
	OriginValue  float64 `json:"origin_value,omitempty"`  // 修改项的原始值
	SuggestValue float64 `json:"suggest_value,omitempty"` // 修改项的建议修改值
}

type ToolsAdDiagnosisFunnelDataRitResult struct {
	Severity     float64 `json:"severity,omitempty"`     // 0~1表示非常好到非常差
	Conclusion   string  `json:"conclusion,omitempty"`   // 诊断结论 eg. 转化率正常
	FunnelData   string  `json:"funnel_data,omitempty"`  // 定义同funnel_data
	Attributions string  `json:"attributions,omitempty"` // 定义同attributions
}
