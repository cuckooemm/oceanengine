package models

import "github.com/antihax/optional"

type AdAdAddReq struct {
	AdvertiserId int64  `json:"advertiser_id"`       // 广告主ID
	CampaignId   string `json:"campaign_id"`         // 广告组ID。注意：广告组ID要求属于广告主ID，且是非删除广告组ID
	Name         string `json:"name"`                // 广告计划名称，长度为1-100个字符，其中1个中文字符算2位。名称不可重复，否则会报错
	Operation    string `json:"operation,omitempty"` // 计划状态 默认值： "enable"开启状态 允许值: "enable"开启,"disable"关闭
	// 投放范围
	DeliveryRange  AdDeliveryRange `json:"delivery_range"`             // 投放范围。 默认值: "DEFAULT"  允许值: "DEFAULT"默认, "UNION"穿山甲
	UnionVideoType string          `json:"union_video_type,omitempty"` // 投放形式（穿山甲视频创意类型），当delivery_range为"UNION"时必填
	// 推广目标
	// 推广目的为应用推广（landing_type=APP）时投放目标参数
	DownloadType         DownloadType        `json:"download_type"`                    // 下载方式 默认值：DOWNLOAD_URL下载链接
	DownloadUrl          string              `json:"download_url,omitempty"`           // 下载链接，当download_type为DOWNLOAD_URL或者QUICK_APP_URL时必填
	QuickAppUrl          string              `json:"quick_app_url,omitempty"`          // 快应用链接，当 download_type 为QUICK_APP_URL时必填
	ExternalUrl          string              `json:"external_url,omitempty"`           // 落地页链接（支持橙子建站落地页和第三方落地页） 当推广目的为APP类型，且download_type为EXTERNAL_URL时必填
	AppType              AppDownloadType     `json:"app_type,omitempty"`               // 下载的应用类型,当download_type为DOWNLOAD_URL时必填
	Package              string              `json:"package,omitempty"`                // 应用包名，当download_type为DOWNLOAD_URL时必填，需要与应用下载链接中包名一致
	DownloadMode         AppDownloadModeType `json:"download_mode,omitempty"`          // 优先从系统应用商店下载（下载模式）允许值：APP_STORE_DELIVERY（仅安卓应用下载支持）、 DEFAULT当应用下载时，默认default下载，可选用APP_STORE_DELIVERY（应用商店直投）
	ConvertId            int64               `json:"convert_id,omitempty"`             // 转化目标， 当出价方式为"OCPM"时必填，当出价方式为"CPC"和"CPM"时非必填。
	OpenUrl              string              `json:"open_url,omitempty"`               // 直达链接(点击唤起APP)
	AdvancedCreativeType string              `json:"advanced_creative_type,omitempty"` // 附加创意类型 (创建后不可修改), 允许值: 游戏礼包码 ATTACHED_CREATIVE_GAME_PACKAGE 推广目的为应用推广类型、下载方式选择下载链接且下载链接为安卓应用下载时才可以设置
	AdGamePackage

	// 推广目的为销售线索收集（landing_type=LINK）时投放目标参数  该推广方式目前不支持“小程序”投放内容，创建计划时不支持附加创意转化组件，可在创建创意时创建附加创意。
	//ExternalUrl string `json:"external_url,omitempty"` // 落地页链接（支持橙子建站落地页和第三方落地页） 当推广目的为APP类型，且download_type为EXTERNAL_URL时必填
	//ConvertId   int64  `json:"convert_id,omitempty"`   // 转化目标， 当出价方式为"OCPM"时必填，当出价方式为"CPC"和"CPM"时非必填。
	//OpenUrl     string `json:"open_url,omitempty"`     // 直达链接(点击唤起APP)

	// 推广目的为抖音号推广（landing_type=AWEME）时投放目标参数
	PromotionType string `json:"promotion_type,omitempty"` // 投放内容，允许值： LIVE：直播 AWEME_HOME_PAGE：抖音主页（默认） LANDING_PAGE_LINK：落地页
	AwemeAccount  string `json:"aweme_account,omitempty"`  // 抖音号，可从【获取绑定抖音号】接口获取，默认取绑定的第一个抖音号
	//ExternalUrl     string          `json:"external_url,omitempty"`     // 落地页链接 当推广目的为AWEME类型，且投放内容为落地页（LANDING_PAGE_LINK）时必填
	ExternalActions []AdConvertType `json:"external_actions,omitempty"` // 转化类型列表，目前仅允许填写一个转化目标 出价方式为OCPM、CPA的计划（投放目标为转化量）必填
	//OpenUrl         string          `json:"open_url,omitempty"`         // 直达链接，当传入落地页的时候可选择传入直达链接

	// 推广目的为门店推广（landing_type=STORE）时投放目标参数
	StoreproUnit       string  `json:"storepro_unit,omitempty"`        // 投放内容 允许值: "STORE"门店, "STORE_ACTIVITY"活动 目前暂时不支持线下商品类型
	StoreType          string  `json:"store_type,omitempty"`           // 门店类型，（storepro_unit 为 "STORE","STORE_ACTIVITY" 时必填）
	AdvertiserStoreIds []int64 `json:"advertiser_store_ids,omitempty"` // 门店ID列表 （storepro_unit 为 "STORE" 时必填） 最多可选择2000个
	StoreproPackId     int64   `json:"storepro_pack_id,omitempty"`     // 活动ID （storepro_unit 为 "STORE_ACTIVITY" 时必填）
	//ConvertId          int64   `json:"convert_id,omitempty"`           // 转化目标， 当出价方式为"OCPM"时必填，当出价方式为"CPC"和"CPM"时非必填。
	//OpenUrl            string  `json:"open_url,omitempty"`             // 直达链接

	// 推广目的为商品目录推广（landing_type=DPA）时投放目标参数
	ProductPlatformId int64                `json:"product_platform_id,omitempty"` // 商品目录ID(ID由【DPA商品广告-查询商品库】 得到)
	ProductId         int64                `json:"product_id,omitempty"`          // 商品ID，当广告组商品类型选择SDPA时必填(ID由【DPA商品广告-获取DPA商品库商品列表】 得到，创建后不可修改)
	CategoryType      string               `json:"category_type,omitempty"`       // 商品目录投放范围,当广告组商品类型选择 DPA 多商品时必填 允许值：NONE不限，"CATEGORY"选择分类，"PRODUCT"指定商品
	DpaCategories     []int64              `json:"dpa_categories,omitempty"`      // 分类列表，category_type为"CATEGORY"时必填，由【DPA商品广告-获取DPA分类】 得到 限制个数1~100
	DpaProducts       []int64              `json:"dpa_products,omitempty"`        // 商品列表，category_type为"PRODUCT"时必填，由【DPA商品广告-获取DPA商品库商品列表】 得到 限制个数1~100
	DpaProductTarget  []AdDpaProductTarget `json:"dpa_product_target,omitempty"`  // 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。

	// 商品目录推广投放目标相关参数：
	DpaAdType           string   `json:"dpa_adtype,omitempty"`             // DPA广告类型，允许值: "DPA_LINK"落地页, "DPA_APP"应用下载
	ParamsType          string   `json:"params_type,omitempty"`            // 链接类型(落地页)，当dpa_adtype为"DPA_LINK"时必填
	DpaExternalUrlField string   `json:"dpa_external_url_field,omitempty"` // 落地页链接字段选择，当params_type为"DPA"时必填
	DpaExternalUrls     []string `json:"dpa_external_urls,omitempty"`      // 落地页链接地址列表，当params_type为"CUSTOM"时必填，目前只能填一个
	//AppType             AppDownloadType `json:"app_type,omitempty"`               // 下载的应用类型,当download_type为DOWNLOAD_URL时必填
	//DownloadUrl         string          `json:"download_url,omitempty"`           // 下载链接，当dpa_adtype为DPA_APP时必填
	//Package             string          `json:"package,omitempty"`                // 应用包名，当download_type为DOWNLOAD_URL时必填，需要与应用下载链接中包名一致
	//ConvertId           int64           `json:"convert_id,omitempty"`             // 转化目标， 当出价方式为"OCPM"时必填，当出价方式为"CPC"和"CPM"时非必填。
	DpaOpenUrlType    string   `json:"dpa_open_url_type,omitempty"`   // 直达链接类型，允许值: "NONE"不启用, "DPA"商品库所含链接, "CUSTOM"自定义链接
	DpaOpenUrlField   string   `json:"dpa_open_url_field,omitempty"`  // 直达链接字段选择，当dpa_open_url_type为"DPA"必填
	DpaOpenUrls       []string `json:"dpa_open_urls,omitempty"`       // 直达链接地址列表，当dpa_open_url_type为"CUSTOM"必填，目前只能填一个
	ExternalUrlParams string   `json:"external_url_params,omitempty"` // 落地页检测参数(DPA推广目的特有,在填写的参数后面添加"=urlencode"(开放平台提供的h5链接地址），其中urlencode(开放平台提供的h5链接地址）替换为商品库中的h5地址encode的结果)

	// 推广目的为电商店铺推广（landing_type=SHOP）时投放目标参数
	//ExternalUrl string `json:"external_url,omitempty"` // 落地页链接（支持橙子建站落地页和第三方落地页） 当推广目的为SHOP类型时必填 对于转化量为目标的计划如OCPM、CPA计划不允许更改，非转化为目标的计划如CPC、CPM计划可更改
	//ConvertId   int64  `json:"convert_id,omitempty"`   // 对于电商店铺推广类型，可以选择预定义的转化id：50（调起店铺）和125（店铺停留）
	//OpenUrl     string `json:"open_url,omitempty"`     // 直达链接，直达链接仅支持部分App唤起，点击创意将优先跳转App，再根据投放内容跳转相关链接

	// 推广目的为商品推广（landing_type=GOODS）时投放目标参数
	//PromotionType string `json:"promotion_type,omitempty"` //  投放内容，允许值： GOODS：商品推广（默认） LIVE：直播
	//AwemeAccount  string `json:"aweme_account,omitempty"`  // 抖音号 当投放内容为直播（LIVE）时有效。抖音号可从【获取绑定抖音号】接口获取，默认取绑定的第一个抖音号
	//ExternalUrl   string `json:"external_url,omitempty"`   // 商品链接（支持橙子建站落地页和第三方落地页） 当投放内容为商品推广（GOODS）时必填
	//ConvertId     int64  `json:"convert_id,omitempty"`     //  转化目标  当pricing出价方式为"OCPM"时必填 目前不支持深度转化目标，默认深度转化目标为"无"
	//OpenUrl       string `json:"open_url,omitempty"`       // 直达链接（点击唤起APP）直达链接仅支持部分App唤起，点击创意将优先跳转App，再根据投放内容跳转相关链接

	// 用户定向
	InterestActionMode string `json:"interest_action_mode,omitempty"` // 行为兴趣 允许值："UNLIMITED"不限,"CUSTOM"自定义,"RECOMMEND"系统推荐。
	InterestCategories []int  `json:"interest_categories,omitempty"`  // 兴趣类目词，当interest_action_mode传CUSTOM时有效
	AdAudience

	// 商品定向 仅推广目的为DPA时可用
	CommodityAudience

	// 预算与出价
	BudgetAndBid
}

// 商品定向 仅推广目的为DPA时可用
type CommodityAudience struct {
	DpaLbs               int64       `json:"dpa_lbs,omitempty"`                // 地域匹配-LBS 开启时，根据用户的地理位置信息，给用户投放位于其附近的产品 允许值：0，1（0表示不启用，1表示启用）
	DpaCity              int64       `json:"dpa_city,omitempty"`               // 地域匹配-商品所在城市 开启时，仅将商品投放给位于该商品设置的可投城市的用户 允许值：0，1（0表示不启用，1表示启用）
	DpaProvince          int64       `json:"dpa_province,omitempty"`           //地域匹配-商品所在省份 开启时，将商品仅投放给位于该商品设置的可投省份的用户 允许值：0，1（0表示不启用，1表示启用）
	DpaLocalAudience     int64       `json:"dpa_local_audience,omitempty"`     // DPA行为重定向，0:不启用，1：启用
	IncludeCustomActions interface{} `json:"include_custom_actions,omitempty"` // 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},]， dpa_local_audience为1时必填; day可选范围:1, 7, 14, 28, 60, 90, 120, 180。
	ExcludeCustomActions interface{} `json:"exclude_custom_actions,omitempty"` // 排除人群包((DPA推广目的特有,格式举例{"days": 7, "code": 1002},]，day可选范围: 1, 7, 14, 28, 60, 90, 120, 180。
	DpaRecommendType     int64       `json:"dpa_recommend_type,omitempty"`     // dpa商品重定向推荐类型，dpa_local_audience为1时必填; 允许值：1（基于重定向推荐更多商品(根据重定向商品和行业特点，推荐更多相关商品投放，包含重定向商品），2仅重定向商品（仅根据重定向人群内定义的重定向行为商品进行投放）
}

// 预算与出价相关字段
type BudgetAndBid struct {
	SmartBidType    SmartBidType    `json:"smart_bid_type"`           // 投放场景(出价方式) 允许值: 常规投放"SMART_BID_CUSTOM", 放量投放"SMART_BID_CONSERVATIVE" 概念解释：常规投放：控制成本，尽量消耗完预算；放量投放：接受成本上浮，尽量消耗更多预算
	AdjustCpa       int64           `json:"adjust_cpa,omitempty"`     // 是否调整自动出价，意味如果预期成本不在范围内将在此基础上调整，仅OCPM支持。 当smart_bid_type=SMART_BID_CONSERVATIVE时选填  当smart_bid_type为"SMART_BID_CONSERVATIVE"且adjust_cpa=0时，cpa_bid由系统自动计算  当smart_bid_type为"SMART_BID_CONSERVATIVE" 且adjust_cpa=1时，cpa_bid必填  允许值: "0", "1"  默认值: "0"
	FlowControlMode FlowControlMode `json:"flow_control_mode"`        // 竞价策略(投放方式) 允许值: "FLOW_CONTROL_MODE_FAST"优先跑量（对应CPC的加速投放）, "FLOW_CONTROL_MODE_SMOOTH"优先低成本（对应CPC的标准投放）, "FLOW_CONTROL_MODE_BALANCE"均衡投放（新增字段）
	BudgetMode      BudgetMode      `json:"budget_mode"`              // 预算类型(创建后不可修改) 允许值: "BUDGET_MODE_DAY"日预算, "BUDGET_MODE_TOTAL"总预算
	Budget          int64           `json:"budget"`                   // 预算(出价方式为CPC、CPM、CPV时，不少于100元；出价方式为OCPM、OCPC时，不少于300元；24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；单次修改预算幅度不能低于100元（增加或者减少）；修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；取值范围: ≥ 0
	ScheduleType    ScheduleType    `json:"schedule_type"`            // 投放时间类型 允许值: "SCHEDULE_FROM_NOW"从今天起长期投放, "SCHEDULE_START_END"设置开始和结束日期
	StartTime       string          `json:"start_time,omitempty"`     // 投放起始时间，当schedule_type为"SCHEDULE_START_END"时必填，形式如：2017-01-01 00:00 广告投放起始时间不允许修改
	EndTime         string          `json:"end_time,omitempty"`       // 投放结束时间，当schedule_type为"SCHEDULE_START_END"时必填，形式如：2017-01-01 00:00
	ScheduleTime    string          `json:"schedule_time,omitempty"`  // 投放时段，默认全时段投放，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。 例如：填写"000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000"，则投放时段为周一到周日的11:30~13:30
	Pricing         AdPricingType   `json:"pricing"`                  // 付费方式（计划出价类型（目前仅穿山甲类型支持OCPC(具体方式：出价类型传OCPC类型，cpa_bid传值 )） 决定投放目标的类型，比如CPC表示点击量，OCPM表示转化量
	Bid             float64         `json:"bid,omitempty"`            // 点击出价/展示出价，当pricing为"CPC"、"CPM"、"CPA"出价方式时必填 pricing为"CPC"时取值范围：0.2-100元； pricing为"CPM"时取值范围：4-100元; pricing为"CPA"时取值范围：1-1500元; 出价不能大于预算否则会报错
	CpaBid          float64         `json:"cpa_bid,omitempty"`        // 目标转化出价/预期成本， 当pricing为"OCPM"、"OCPC"出价方式时必填 pricing为"OCPC"时取值范围：0.1-10000元； pricing为"OCPM"时取值范围：0.1-10000元； 出价不能大于预算否则会报错
	DeepBidType     AdDeepBidType   `json:"deep_bid_type,omitempty"`  // 深度优化方式 对于每次付费的转化，深度优化类型需要设置为BID_PER_ACTION（每次付费出价） 具体概念见【深度优化方式】  当转化目标中含有深度转化时，该字段必填。 API后期会上线获取可用深度优化方式，请关注上线通知
	DeepCpaBid      float64         `json:"deep_cpabid,omitempty"`    // 深度优化出价，deep_bid_type为"DEEP_BID_MIN"时必填。当对应的转化convert_id，设定深度转化目标时才会有效。
	LubanRoiGoal    float64         `json:"luban_roi_goal,omitempty"` // 鲁班目标ROI出价策略系数。推广目的为商品推广(GOODS)时可填。当传入该参数时，表示启用鲁班ROI优化，支持范围(0,100]，精度：保留小数点后四位
	RoiGoal         float64         `json:"roi_goal,omitempty"`       // 深度转化ROI系数, 范围(0,5]，精度：保留小数点后四位, deep_bid_type为"ROI_COEFFICIENT"时必填
}

type AdAdAddRsp struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      AdAdAddRspData `json:"data"`
	RequestId string         `json:"request_id"`
}
type AdAdAddRspData struct {
	AdId int64 `json:"ad_id"`
}

type AdAdGetOpts struct {
	Filtering optional.Interface `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	Fields    optional.Interface `json:"fields,omitempty"`    // 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典
	Page      optional.Int64     `json:"page,omitempty"`      // 当前页码: 1
	PageSize  optional.Int64     `json:"page_size,omitempty"` // 页面大小 默认值: 10， 取值范围：1-1000
}

type AdAdGetFiltering struct {
	Ids         []int64  `json:"ids,omitempty"`          // 按广告计划ID过滤，范围为1-100
	AdName      string   `json:"ad_name,omitempty"`      // 按广告计划name过滤，长度为1-30个字符
	PricingList []string `json:"pricing_list,omitempty"` // 按出价方式过滤
	Status      string   `json:"status,omitempty"`       // 按计划状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示
	CampaignId  int64    `json:"campaign_id,omitempty"`  //
}

type AdAdGetRsp struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      AdAdGetRspData `json:"data"`
	RequestId string         `json:"request_id"`
}

type AdAdGetRspData struct {
	List     []AdAdGetRspDataList `json:"list"`
	PageInfo PageInfo             `json:"page_info"`
}

type Geolocation struct {
	Radius       int64   `json:"radius,omitempty"`        // 半径
	Name         string  `json:"name,omitempty"`          // 地点名称
	Long         float64 `json:"long,omitempty"`          // 经度
	Lat          float64 `json:"lat,omitempty"`           // 纬度
	City         string  `json:"city,omitempty"`          // 城市名
	StreetNumber string  `json:"street_number,omitempty"` // 街道号
	Street       string  `json:"street,omitempty"`        // 街道名
	District     string  `json:"district,omitempty"`      // 区域
	Province     string  `json:"province,omitempty"`      // 省份名
}

type AdAdGetRspDataList struct {
	Id                   int64    `json:"id"`                     // 计划ID
	AdId                 int64    `json:"ad_id"`                  // 计划ID,返回值同id
	Name                 string   `json:"name"`                   // 计划名称
	AdvertiserId         int64    `json:"advertiser_id"`          // 广告主ID
	CampaignId           int64    `json:"campaign_id"`            // 广告组ID
	ModifyTime           string   `json:"modify_time"`            // 计划上次修改时间标识(用于更新计划时提交,服务端判断是否基于最新信息修改)
	AdModifyTime         string   `json:"ad_modify_time"`         // 计划上次修改时间
	AdCreateTime         string   `json:"ad_create_time"`         // 计划创建时间
	Status               string   `json:"status"`                 // 广告计划投放状态,详见【附录-广告计划投放状态】(进入投放之前,优先披露审核状态,此时优先于启用暂停,启用暂停信息以opt_status为准)
	OptStatus            string   `json:"opt_status"`             // 广告计划操作状态 允许值: "AD_STATUS_ENABLE","AD_STATUS_DISABLE"
	DeliveryRange        string   `json:"delivery_range"`         // 投放范围
	UnionVideoType       string   `json:"union_video_type"`       // 投放形式（穿山甲视频创意类型）默认值: ORIGINAL_VIDEO原生 允许值: "ORIGINAL_VIDEO"原生, "REWARDED_VIDEO"激励视频,"SPLASH_VIDEO"开屏
	DownloadType         string   `json:"download_type"`          // 应用下载方式，推广目的为APP时有值。返回值：DOWNLOAD_URL下载链接，QUICK_APP_URL快应用+下载链接，EXTERNAL_URL落地页链接
	DownloadUrl          string   `json:"download_url"`           // 下载链接，当推广类型为应用推广，且当download_type为DOWNLOAD_URL或者QUICK_APP_URL时有值
	QuickAppUrl          string   `json:"quick_app_url"`          // 快应用链接，当推广类型为应用推广，且download_type为QUICK_APP_URL时有值
	ExternalUrl          string   `json:"external_url"`           // 落地页链接，投放内容或下载方式为落地页时有值
	AppType              string   `json:"app_type"`               // 下载类型，当推广类型为应用推广且download_type为DOWNLOAD_URL或者QUICK_APP_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	DownloadMode         string   `json:"download_mode"`          // 优先从系统应用商店下载（下载模式） 允许值：APP_STORE_DELIVERY（仅安卓应用下载支持）、 DEFAULT当应用下载时，默认default下载，可选用APP_STORE_DELIVERY（应用商店直投），当为该值时，将优先跳转目标应用对应手机系统应用商店安装详情页，跳转失败则使用下载链接下载。 请确保投放的应用在应用商店内已上架
	ConvertId            string   `json:"convert_id"`             // 转化目标，其中convert_id数值较小时为预定义转化，具体枚举可查看【附录-预定义转化类型】
	ExternalActions      []string `json:"external_actions"`       // 转化类型，目前当推广类型为抖音时有值，允许值："AD_CONVERT_TYPE_FOLLOW_ACTION", "AD_CONVERT_TYPE_MESSAGE_ACTION", "AD_CONVERT_TYPE_INTERACTION"
	OpenUrl              string   `json:"open_url"`               // 直达链接(点击唤起APP)
	AdvancedCreativeType string   `json:"advanced_creative_type"` // 附加创意类型 允许值: ATTACHED_CREATIVE_GAME_PACKAGE游戏礼包码,ATTACHED_CREATIVE_GAME_FORM游戏表单收集,ATTACHED_CREATIVE_GAME_SUBSCRIBE游戏预约,ATTACHED_CREATIVE_NONE无 推广目的为应用推广类型、下载方式选择下载链接且下载链接为安卓应用下载时才可以设置

	StoreproUnit        string               `json:"storepro_unit"`          // 门店推广-投放内容，当推广目的为STORE(门店推广)时有值。 取值: "STORE"门店, "STORE_ACTIVITY"活动 目前暂时不支持线下商品类型
	StoreType           string               `json:"store_type"`             // 门店类型，（storepro_unit 为 "STORE" 时有值。 取值: "STORE_NORMAL"平台通用门店, "STORE_THIRT_PARTY"第三方门店, "STORE_DOUYIN"抖音POI门店
	AdvertiserStoreIds  []int64              `json:"advertiser_store_ids"`   // 门店ID列表 （storepro_unit 为 "STORE" 时有值
	StoreproPackId      int64                `json:"storepro_pack_id"`       // 活动ID （storepro_unit 为 "STORE_ACTIVITY" 时有值
	ProductLlatformId   int64                `json:"product_platform_id"`    // 产品目录ID(ID由查询产品目录接口得到), 当推广目的landing_type为DPA时有值
	ProductId           int64                `json:"product_id"`             // 商品id，当推广目的为 DPA 广告组商品类型为 SDPA 时有值
	CategoryType        string               `json:"category_type"`          // DPA投放范围，取值：NONE不限，"CATEGORY"选择分类，"PRODUCT"指定商品
	DpaCategories       []int64              `json:"dpa_categories"`         // 分类列表，category_type取值范围为CATEGORY时有值
	DpaProducts         []int64              `json:"dpa_products"`           // 商品列表，category_type为PRODUCT时有值
	DpaProductTarget    []AdDpaProductTarget `json:"dpa_product_target"`     // 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。
	DpaAdtype           []string             `json:"dpa_adtype"`             // dpa广告类型，取值范围："DPA_LINK"落地页, "DPA_APP"应用下载
	ParamsType          string               `json:"params_type"`            // 链接类型(落地页)，当dpa_adtype为"DPA_LINK"时有值，取值: "DPA"商品库所含链接, "CUSTOM"自定义链接
	DpaExternalUrlField string               `json:"dpa_external_url_field"` // 落地页链接字段选择，当params_type为"DPA"时有值
	DpaExternalUrls     []string             `json:"dpa_external_urls"`      // 落地页链接地址列表，当params_type为"CUSTOM"时有值
	Package             string               `json:"package"`                // 应用包名，当推广类型为应用推广且download_type为DOWNLOAD_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	InventoryType       []string             `json:"inventory_type"`         // 创意投放位置,详见【附录-投放位置】。创建选择优选广告位时，此字段回会返回对应的优选广告位
	PromotionType       string               `json:"promotion_type"`         // 投放内容 GOODS：商品推广 LIVE：直播 AWEME_HOME_PAGE：抖音主页 LANDING_PAGE_LINK：落地页
	AwemeAccount        string               `json:"aweme_account"`          // 抖音号
	SubscribeUrl        string               `json:"subscribe_url"`          // 	游戏营销场景-预约下载链接
	FormId              int64                `json:"form_id"`                // 	游戏营销场景-表单id
	FormIndex           int64                `json:"form_index"`             // 	游戏营销场景-表单位置索引
	AppDesc             string               `json:"app_desc"`               // 	游戏营销场景-应用描述
	AppIntroduction     string               `json:"app_introduction"`       // 	游戏营销场景-应用介绍
	AppThumbnails       []string             `json:"app_thumbnails"`         // 	游戏营销场景-应用图片集，返回图片集Id
	DpaOpenUrlType      string               `json:"dpa_open_url_type"`      // 直达链接类型，取值: "NONE"不启用, "DPA"商品库所含链接, "CUSTOM"自定义链接 商品库链接对应商品库内调起字段
	DpaOpenUrlField     string               `json:"dpa_open_url_field"`     // 直达链接字段选择，当dpa_open_url_type为"DPA"时有值
	DpaOpenUrls         []string             `json:"dpa_open_urls"`          // 达链接地址列表，当dpa_open_url_type为"CUSTOM"时有值
	ExternalUrlParams   string               `json:"external_url_params"`    // 落地页检测参数(DPA推广目的特有,在填写的参数后面添加"=urlencode(开放平台提供的h5链接地址）"，其中urlencode(开放平台提供的h5链接地址）替换为商品库中的h5地址encode的结果)
	OpenUrlParams       string               `json:"open_url_params"`        // 直达链接检测参数(DPA推广目的特有,在“产品库中提取的scheme地址"后面追加填写的参数)
	AdGamePackage
	AdAudience
	CommodityAudience
	BudgetAndBid
}

type AdDpaProductTarget struct {
	Title string `json:"title"` // 筛选字段
	Rule  string `json:"rule"`  // 定向规则，允许值：'=', '!=', '>', '<', '>=', '<=', 'contain', 'exclude', 'notEmpty'
	Type  string `json:"type"`  // 字段类型，允许值：'int', 'double', 'long', 'string'
	Value string `json:"value"` // 	规则值
}
type AdGamePackage struct {
	GamePackageDesc         string   `json:"game_package_desc,omitempty"`          // 应用描述,最少1字，最多15字
	GamePackageBatchId      int64    `json:"game_package_batch_id,omitempty"`      // 游戏礼包码id，目前仅支持直接发券类型
	GamePackageThumbnailIds []string `json:"game_package_thumbnail_ids,omitempty"` //应用图片集，图片image_id，有且仅有两个，游戏礼包码时有值,可以从 【素材管理-获取图片素材】 接口中获取，建议尺寸 16:9 与game_package_thumbnails同时传入时，以game_package_thumbnail_ids字段为准
}
type AdAudience struct {
	AudiencePackageId         int64            `json:"audience_package_id,omitempty"`          // 定向包ID，定向包ID由【工具-定向包管理-获取定向包】获取 如果同时传定向包ID和自定义用户定向参数时，仅定向包中的定向生效
	District                  AdDistrict       `json:"district,omitempty"`                     // 地域 取值: "CITY"省市, "COUNTY"区县, "BUSINESS_DISTRICT"商圈,"NONE"不限
	City                      []int64          `json:"city,omitempty"`                         // 	地域定向省市或者区县列表
	BusinessIds               []int64          `json:"business_ids,omitempty"`                 // 商圈ID数组
	Geolocation               []Geolocation    `json:"geolocation,omitempty"`                  //  从地图添加(地图位置)
	LocationType              string           `json:"location_type,omitempty"`                // 位置类型 允许值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户 当city和district有值时必填
	Gender                    AdGender         `json:"gender,omitempty"`                       // 性别
	Age                       []AdAge          `json:"age,omitempty"`                          // 年龄
	RetargetingTagsInclude    []int64          `json:"retargeting_tags_include,omitempty"`     // 定向人群包列表（自定义人群），内容为人群包id
	RetargetingTagsExclude    []int64          `json:"retargeting_tags_exclude,omitempty"`     // 排除人群包列表（自定义人群），内容为人群包id	InterestActionMode        string           `json:"interest_action_mode"`          // 行为兴趣 取值："UNLIMITED"不限,"CUSTOM"自定义,"RECOMMEND"系统推荐。若与自定义人群同时使用，系统推荐("RECOMMEND")不生效 仅推广范围为默认时可填，且不可与老版行为兴趣定向同时填写，否则会报错
	ActionScene               []string         `json:"action_scene,omitempty"`                 // string[]	行为场景，详见【附录-行为场景】，当interest_action_mode传CUSTOM时有效 允许值："E-COMMERCE"电商互动行为, "NEWS"资讯互动行为, "APP"APP推广互动行为
	ActionDays                int              `json:"action_days,omitempty"`                  // 	用户发生行为天数，当interest_action_mode传CUSTOM时有效 允许值：7, 15, 30, 60, 90, 180, 365
	ActionCategories          []int            `json:"action_categories,omitempty"`            // 行为类目词
	ActionWords               []int            `json:"action_words,omitempty"`                 // 行为关键词	InterestCategories        []int64          `json:"interest_categories"`           // 兴趣类目词，当interest_action_mode传CUSTOM时有效
	InterestWords             []int64          `json:"interest_words,omitempty"`               // 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传CUSTOM时有效
	AdTag                     []int64          `json:"ad_tag,omitempty"`                       // （老版行为兴趣）兴趣分类,如果传"空数组"表示不限，如果"数组传0"表示系统推荐,如果按兴趣类型传表示自定义，兴趣类型详见【附件-ad_tag.json】 仅推广范围为穿山甲时可填，且不可与新版行为兴趣定向同时填写，否则会报错
	InterestTags              []int64          `json:"interest_tags,omitempty"`                // （老版行为兴趣）兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。 仅推广范围为穿山甲时可填，且不可与新版行为兴趣定向同时填写，否则会报错
	AppBehaviorTarget         string           `json:"app_behavior_target,omitempty"`          // （老版行为兴趣）APP行为 取值：NONE不限，CATEGORY按分类，APP按APP 仅推广范围为穿山甲时可填，且不可与新版行为兴趣定向同时填写，否则会报错
	AppCategory               []int64          `json:"app_category,omitempty"`                 // （老版行为兴趣）APP行为定向——按分类, 详见【附件-app_category.json】（请注意如果投放的是"应用下载-IOS"不支持设置APP行为定向，请勿传值）当app_behavior_target为CATEGORY时返回 仅推广范围为穿山甲时可填，且不可与新版行为兴趣定向同时填写，否则会报错
	AppIds                    []int64          `json:"app_ids,omitempty"`                      // （老版行为兴趣）APP行为定向——按APP（请注意如果投放的是"应用下载-IOS"不支持设置APP行为定向，请勿传值。）可通过【工具-查询工具-查询应用信息】获取。当app_behavior_target为APP时有值 仅推广范围为穿山甲时可填，且不可与新版行为兴趣定向同时填写，否则会报错
	AwemeFanBehaviors         []AdDyActionType `json:"aweme_fan_behaviors,omitempty"`          // 抖音达人互动用户行为类型
	AwemeFanCategories        []int64          `json:"aweme_fan_categories,omitempty"`         // 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanAccounts          []int64          `json:"aweme_fan_accounts,omitempty"`           // 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向
	AwemeFansNumbers          []int64          `json:"aweme_fans_numbers,omitempty"`           // （抖音号推广特有）账号粉丝相似人群（添加抖音账号，会将广告投放给对应账号的相似人群粉丝）	FilterAwemeAbnormalActive int64            `json:"filter_aweme_abnormal_active"`  // （抖音号推广特有）过滤高活跃用户 取值：0表示不过滤，1表示过滤
	FilterAwemeAbnormalActive int64            `json:"filter_aweme_abnormal_active,omitempty"` // 抖音号推广特有）过滤高活跃用户 允许值：0表示不过滤，1表示过滤
	FilterAwemeFansCount      int64            `json:"filter_aweme_fans_count,omitempty"`      // 抖音号推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterOwnAwemeTans        int64            `json:"filter_own_aweme_fans,omitempty"`        // 抖音号推广特有）过滤自己的粉丝 允许值：0表示不过滤，1表示过滤
	SuperiorPopularityType    string           `json:"superior_popularity_type,omitempty"`     // 媒体定向  对于选择自定义媒体定向流量包，该字段不传，传flow_package或exclude_flow_package字段即可 媒体定向仅支持穿山甲、穿山甲-精选游戏广告位
	FlowPackage               []int64          `json:"flow_package,omitempty"`                 // 定向逻辑 定向逻辑和排除定向逻辑只能选其一
	ExcludeFlowPackage        []int64          `json:"exclude_flow_package,omitempty"`         // 	排除定向逻辑 定向逻辑和排除定向逻辑只能选其一
	Platform                  []string         `json:"platform,omitempty"`                     // 	平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载方式的时候，平台可多选。
	AndroidOsv                string           `json:"android_osv,omitempty"`                  // 最低安卓版本，当app_type为"APP_ANDROID"选填,其余情况不填
	IosOsv                    string           `json:"ios_osv,omitempty"`                      // 最低IOS版本，当app_type为"APP_IOS"选填,其余情况不填,
	DeviceType                []string         `json:"device_type,omitempty"`                  // 设备类型。 取值是："MOBILE", "PAD"。缺省表示不限设备类型。穿山甲已经全量，投放范围为默认时需要有白名单权限才可以

	Ac                    []string `json:"ac,omitempty"`                      // 网络类型,允许值: "WIFI", "2G", "3G", "4G"
	Carrier               []string `json:"carrier,omitempty"`                 // 运营商, 允许值: "MOBILE", "UNICOM", "TELCOM"
	HideIfExists          int      `json:"hide_if_exists,omitempty"`          // 过滤已安装，当推广目标为安卓应用下载时可填，0表示不限，1表示过滤，2表示定向。默认为不限
	ConvertedTimeDuration string   `json:"converted_time_duration,omitempty"` // 过滤时间范围
	HideIfConverted       string   `json:"hide_if_converted,omitempty"`       // 过滤已转化用户
	ActivateType          []string `json:"activate_type,omitempty"`           // 新用户(新用户使用头条的时间)
	ArticleCategory       []string `json:"article_category,omitempty"`        // 文章分类 文章分类：只针对投放在详情页位置的广告生效, 不支持人群预估
	DeviceBrand           []string `json:"device_brand,omitempty"`            // 手机品牌
	LaunchPrice           []int64  `json:"launch_price,omitempty"`            // 手机价格,传入价格区间，最高传入11000（表示1w以上）
	AutoExtendEnabled     int64    `json:"auto_extend_enabled,omitempty"`     // 是否启用智能放量。 允许值是：0、1。缺省为 0。 注意点：智能放量不支持受众预估
	AutoExtendTargets     []string `json:"auto_extend_targets,omitempty"`     // 可放开定向。当auto_extend_enabled=1 时选填。详见：【附录-可开放定向】。缺省为全不选。
}

type AdBudgetUpdateReq struct {
	AdvertiserId int64          `json:"advertiser_id"`
	Data         []AdBudgetData `json:"data,omitempty"` // 批量修改预算，包含计划ID和预算，list长度限制1～100.
}

type AdBudgetData struct {
	AdId int64 `json:"ad_id"` // 广告计划ID，广告计划id需要属于广告主，且ad_id不能重复，否则会报错！
	// 预算，单位：元。
	// 24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；
	// 单次修改预算幅度不能低于100元（增加或者减少）;
	// 修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；
	Budget int64 `json:"budget"`
}

type AdBudgetUpdateRsp struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Data      AdBudgetUpdateRspData `json:"data"`
	RequestId string                `json:"request_id"`
}

type AdBudgetUpdateRspData struct {
	AdIds []int64 `json:"ad_ids"`
}

type AdStatusUpdateReq struct {
	AdvertiserId int64    `json:"advertiser_id"`
	AdIds        []int64  `json:"ad_ids"`
	OptStatus    AdStatus `json:"opt_status"`
}

type AdStatusUpdateRsp struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Data      AdStatusUpdateRspData `json:"data"`
	RequestId string                `json:"request_id"`
}

type AdStatusUpdateRspData struct {
	AdIds []int64 `json:"ad_ids"`
}
