package models

type BudgetMode string
type LandingType string
type CampaignStatus string
type AdDeliveryRange string
type AdCreativeUnionVideoType string
type DownloadType string
type AppDownloadType string
type AppDownloadModeType string
type AdConvertType string
type AdDistrict string           // 定向 地域
type AdGender string             // 定向 性别
type AdAge string                // 定向 年龄区间
type AdDyActionType string       // 抖音用户行为类型
type AdPricingType string        // 广告出价类型
type AdDeepBidType string        // 深度优化方式
type SmartBidType string         // 投放场景(出价方式)
type FlowControlMode string      // 计划投放速度类型
type ScheduleType string         // 投放时间类型
type InventoryType string        // 投放位置
type AdvancedCreativeType string // 附加创意类型
type VideoCoverStatus string     // 封面生成状态

type UploadType string        // 上传方式
type AdStatus string          // 广告状态
type CreativeImageMode string // 媒体素材类型
type LearningPhase string     // 计划学习期状态

const (
	BudgetModeInfinite BudgetMode = "BUDGET_MODE_INFINITE" // 不限
	BudgetModeDay      BudgetMode = "BUDGET_MODE_DAY"      // 日预算
	BudgetModeTotal    BudgetMode = "BUDGET_MODE_TOTAL"    // 总预算

	AdStatusEnable  AdStatus = "enable"  // 启用计划
	AdStatusDelete  AdStatus = "delete"  // 删除计划
	AdStatusDisable AdStatus = "disable" // 暂停计划

	LandingTypeLink  LandingType = "LINK"  // 	销售线索收集
	LandingTypeApp   LandingType = "APP"   // 	应用推广
	LandingTypeDpa   LandingType = "DPA"   // 	商品目录推广
	LandingTypeGoods LandingType = "GOODS" // 	商品推广（鲁班）
	LandingTypeStore LandingType = "STORE" // 	门店推广
	LandingTypeAweme LandingType = "AWEME" // 	抖音号推广
	LandingTypeShop  LandingType = "SHOP"  // 	电商店铺推广
	// LandingTypeArtical LandingType = "ARTICAL" // 	头条文章推广，目前API暂不支持该推广目的的设定，可在平台侧选择该推广目的类型

	CampaignStatusEnable                 CampaignStatus = "CAMPAIGN_STATUS_ENABLE"                   // 	启用
	CampaignStatusDisable                CampaignStatus = "CAMPAIGN_STATUS_DISABLE"                  // 	暂停
	CampaignStatusDelete                 CampaignStatus = "CAMPAIGN_STATUS_DELETE"                   // 	删除
	CampaignStatusAll                    CampaignStatus = "CAMPAIGN_STATUS_ALL"                      // 	所有包含已删除
	CampaignStatusNotDelete              CampaignStatus = "CAMPAIGN_STATUS_NOT_DELETE"               // 	所有不包含已删除（状态过滤默认值）
	CampaignStatusAdvertiserBudgetExceed CampaignStatus = "CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED" // 	超出广告主日预算

	AdDeliveryRangeDefault   AdDeliveryRange = "DEFAULT"   // 默认
	AdDeliveryRangeUnion     AdDeliveryRange = "UNION"     // 只投放到资讯联盟（穿山甲）
	AdDeliveryRangeUniversal AdDeliveryRange = "UNIVERSAL" // UNIVERSAL	通投智选

	AdCreativeUnionVideoOriginal AdCreativeUnionVideoType = "ORIGINAL_VIDEO" // 	原生视频
	AdCreativeUnionVideoRewarded AdCreativeUnionVideoType = "REWARDED_VIDEO" //	激励视频
	AdCreativeUnionVideoSplash   AdCreativeUnionVideoType = "SPLASH_VIDEO"   //	穿山甲开屏

	DownloadTypeUrl         DownloadType = "DOWNLOAD_URL"  // 下载链接
	DownloadTypeQuickAppUrl DownloadType = "QUICK_APP_URL" // 快应用+下载链接
	DownloadTypeExternalUrl DownloadType = "EXTERNAL_URL"  // 快应用+下载链接

	AppDownloadTypeAndroid AppDownloadType = "APP_ANDROID" // 	Android
	AppDownloadTypeIos     AppDownloadType = "APP_IOS"     // iOS

	AppDownloadModeTypeAppStore AppDownloadModeType = "APP_STORE_DELIVERY" // 仅安卓应用下载支持 (应用商店直投），选择后，将优先跳转目标应用对应手机系统应用商店安装详情页，跳转失败则使用下载链接下载。
	AppDownloadModeTypeDefault  AppDownloadModeType = "DEFAULT"            // 当应用下载时，默认default下载

	AdConvertTypeInter            AdConvertType = "AD_CONVERT_TYPE_INTERACTION"               // 当应用下载时，默认default下载
	AdConvertTypeFollow           AdConvertType = "AD_CONVERT_TYPE_FOLLOW_ACTION"             // 账号关注
	AdConvertTypeComment          AdConvertType = "AD_CONVERT_TYPE_COMMENT_ACTION"            // 视频评论
	AdConvertTypeShare            AdConvertType = "AD_CONVERT_TYPE_SHARE_ACTION"              // 分享
	AdConvertTypeClickLandingPage AdConvertType = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"        // 访问推广详情页
	AdConvertTypeMessage          AdConvertType = "AD_CONVERT_TYPE_MESSAGE_ACTION"            // 私信消息
	AdConvertTypeLiveEnter        AdConvertType = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"         // 直播间观看
	AdConvertTypeLiveStayTime     AdConvertType = "AD_CONVERT_TYPE_LIVE_STAY_TIME"            // 直播间停留
	AdConvertTypeClickProduct     AdConvertType = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION" // 直播间查看商品
	AdConvertTypeLiveGift         AdConvertType = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"          // 直播间打赏
	AdConvertTypeNewFollow        AdConvertType = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"         // 粉丝增长

	AdDistrictCity     AdDistrict = "CITY"              // 省市
	AdDistrictCounty   AdDistrict = "COUNTY"            // 区县
	AdDistrictBusiness AdDistrict = "BUSINESS_DISTRICT" // 商圈
	AdDistrictNone     AdDistrict = "NONE"              // 不限

	AdGenderNone   AdGender = "NONE"          // 不限
	AdGenderFemale AdGender = "GENDER_FEMALE" // 女
	AdGenderMale   AdGender = "GENDER_MALE"   // 男

	AdAgeBetween_18_23 AdAge = "AGE_BETWEEN_18_23" // 18-23 岁
	AdAgeBetween_24_30 AdAge = "AGE_BETWEEN_24_30"
	AdAgeBetween_31_40 AdAge = "AGE_BETWEEN_31_40"
	AdAgeBetween_41_49 AdAge = "AGE_BETWEEN_41_49"
	AdAgeAbove_50      AdAge = "AGE_ABOVE_50"

	DyActionTypeFollow  AdDyActionType = "FOLLOWED_USER"  // 已关注用户
	DyActionTypeComment AdDyActionType = "COMMENTED_USER" // 视频互动-已评论用户
	DyActionTypeLiked   AdDyActionType = "LIKED_USER"     // 视频互动-已点赞用户
	DyActionTypeShared  AdDyActionType = "SHARED_USER"    // 视频互动-已分享用户

	AdPricingTypeCPC  AdPricingType = "PRICING_CPC"  // CPC ，出价范围（单位元）:0.2-100，日预算范围（单位元）：大于100，总预算范围：大于最低日预算乘投放天数
	AdPricingTypeCPM  AdPricingType = "PRICING_CPM"  // CPM ，出价范围（单位元）: 4-100，日预算范围（单位元）：大于100，总预算范围：大于最低日预算乘投放天数
	AdPricingTypeOCPC AdPricingType = "PRICING_OCPC" // OCPC（已下线，仅投放范围为穿山甲可用）
	AdPricingTypeOCPM AdPricingType = "PRICING_OCPM" // OCPM，出价范围（单位元）:0.1-10000，日预算范围（单位元）：大于300，总预算范围：大于最低日预算乘投放天数
	AdPricingTypeCPV  AdPricingType = "PRICING_CPV"  // CPV ，出价范围（单位元）:0.07-100，日预算范围（单位元）：大于100，总预算范围：大于最低日预算乘投放天数（CPV广告只支持投放到头条系广告位，不支持投放到视频信息流如西瓜、火山、抖音）
	AdPricingTypeCPA  AdPricingType = "PRICING_CPA"  // CPA

	SmartBidTypeCustom       SmartBidType = "SMART_BID_CUSTOM"       // 	常规投放：控制成本，尽量消耗完预算
	SmartBidTypeConservative SmartBidType = "SMART_BID_CONSERVATIVE" // 	放量投放：接受成本上浮，尽量消耗更多预算

	FlowControlModeFast    FlowControlMode = "FLOW_CONTROL_MODE_FAST"    // 	优先跑量（对应CPC的加速投放）
	FlowControlModeSmooth  FlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"  // 	优先低成本（对应CPC的标准投放）
	FlowControlModeBalance FlowControlMode = "FLOW_CONTROL_MODE_BALANCE" // 	均衡投放（新增字段）

	ScheduleTypeNow      ScheduleType = "SCHEDULE_FROM_NOW"  // 	从今天起长期投放
	ScheduleTypeStartEnd ScheduleType = "SCHEDULE_START_END" // 	设置开始和结束日期

	AdDeepBidTypeDefault            AdDeepBidType = "DEEP_BID_DEFAULT"      // 	无深度优化
	AdDeepBidTypePacing             AdDeepBidType = "DEEP_BID_PACING"       // 自动优化
	AdDeepBidTypeMin                AdDeepBidType = "DEEP_BID_MIN"          // 自定义双出价
	AdDeepBidTypeROICoefficient     AdDeepBidType = "ROI_COEFFICIENT"       // ROI系数
	AdDeepBidTypeROIPacing          AdDeepBidType = "ROI_PACING"            // ROI自动出价
	AdDeepBidTypeMinSecondStage     AdDeepBidType = "MIN_SECOND_STAGE"      // 自定义两阶段
	AdDeepBidTypePacingSecondStage  AdDeepBidType = "PACING_SECOND_STAGE"   // 动态两阶段
	AdDeepBidTypeSmartBid           AdDeepBidType = "SMARTBID"              // 	自动出价
	AdDeepBidTypeAutoMinSecondStage AdDeepBidType = "AUTO_MIN_SECOND_STAGE" // 	自动出价两阶段
	AdDeepBidTypeBidPerAction       AdDeepBidType = "BID_PER_ACTION"        // 	每次付费出价

	InventoryFeed              InventoryType = "INVENTORY_FEED"              // 头条信息流（广告投放）
	InventoryVideoFeed         InventoryType = "INVENTORY_VIDEO_FEED"        // 西瓜信息流（广告投放）
	InventoryHotsoonFeed       InventoryType = "INVENTORY_HOTSOON_FEED"      // 火山信息流（广告投放）
	InventoryAwemeFeed         InventoryType = "INVENTORY_AWEME_FEED"        // 抖音信息流（广告投放）
	InventoryUnionSlot         InventoryType = "INVENTORY_UNION_SLOT"        // 穿山甲（广告投放）
	InventoryUnionBoutiqueGame InventoryType = "UNION_BOUTIQUE_GAME"         // 穿山甲精选休闲游戏（广告投放）
	InventoryUnionSplashSlot   InventoryType = "INVENTORY_UNION_SPLASH_SLOT" // 穿山甲开屏广告（广告投放）
	InventoryAwemeSearch       InventoryType = "INVENTORY_AWEME_SEARCH"      // 搜索广告——抖音位（广告投放）
	InventorySearch            InventoryType = "INVENTORY_SEARCH"            // 搜索广告——头条位（广告投放）
	InventoryUniversal         InventoryType = "INVENTORY_UNIVERSAL"         // 通投智选（广告投放）
	InventoryDeauty            InventoryType = "INVENTORY_BEAUTY"            // 轻颜相机
	InventoryPIPIXIA           InventoryType = "INVENTORY_PIPIXIA"           // 皮皮虾
	InventoryAutomobile        InventoryType = "INVENTORY_AUTOMOBILE"        // 懂车帝
	InventoryStudy             InventoryType = "INVENTORY_STUDY"             // 好好学习
	InventoryFaceU             InventoryType = "INVENTORY_FACE_U"            // faceu

	AttachedCreativeNone           AdvancedCreativeType = "ATTACHED_CREATIVE_NONE"            // 无
	AttachedCreativePhone          AdvancedCreativeType = "ATTACHED_CREATIVE_PHONE"           // 电话拨打
	AttachedCreativeForm           AdvancedCreativeType = "ATTACHED_CREATIVE_FORM"            // 表单收集
	AttachedCreativeCommerceCard   AdvancedCreativeType = "ATTACHED_CREATIVE_COMMERCE_CARD"   // 电商（即落地页）卡片
	AttachedCreativeDownloadCard   AdvancedCreativeType = "ATTACHED_CREATIVE_DOWNLOAD_CARD"   // 商品下载卡片
	AttachedCreativeConsultant     AdvancedCreativeType = "ATTACHED_CREATIVE_CONSULTANT"      // 	在线咨询
	AttachedCreativeCoupon         AdvancedCreativeType = "ATTACHED_CREATIVE_COUPON"          // 	卡券
	AttachedCreativeCard           AdvancedCreativeType = "ATTACHED_CREATIVE_CARD"            // 	图片磁贴
	AttachedCreativeInteract       AdvancedCreativeType = "ATTACHED_CREATIVE_INTERACT"        // 选择磁贴
	AttachedCreativeSmartPhone     AdvancedCreativeType = "ATTACHED_CREATIVE_SMART_PHONE"     // 智能电话
	AttachedCreativeVoteInteract   AdvancedCreativeType = "ATTACHED_CREATIVE_VOTE_INTERACT"   // 投票磁贴
	AttachedCreativeGamePackage    AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_PACKAGE"    // 游戏礼包
	AttachedCreativeCouponInteract AdvancedCreativeType = "ATTACHED_CREATIVE_COUPON_INTERACT" // 优惠券磁贴
	AttachedCreativeGameForm       AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_FORM"       // 游戏表单收集
	AttachedCreativeGameSubscribe  AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_SUBSCRIBE"  // 游戏预约按钮

	UploadByFile UploadType = "UPLOAD_BY_FILE" // 文件
	UploadByUrl  UploadType = "UPLOAD_BY_URL"  // url

	VideoCoverStatusRunning VideoCoverStatus = "RUNNING" // 生成中
	VideoCoverStatusSuccess VideoCoverStatus = "SUCCESS" // 成功
	VideoCoverStatusFailed  VideoCoverStatus = "FAILED"  // 失败

	CreativeImageModeSmall              CreativeImageMode = "CREATIVE_IMAGE_MODE_SMALL"              // 小图，宽高比1.52，大小1.5M以下，下限：456 & 300，上限：1368 & 900
	CreativeImageModeLarge              CreativeImageMode = "CREATIVE_IMAGE_MODE_LARGE"              // 大图，横版大图宽高比1.78，大小1.5M以下，下限：1280 & 720，上限：2560 & 1440
	CreativeImageModeGroup              CreativeImageMode = "CREATIVE_IMAGE_MODE_GROUP"              // 组图，宽高比1.52，大小1.5M以下，下限：456 & 300，上限：1368 & 900
	CreativeImageModeVideo              CreativeImageMode = "CREATIVE_IMAGE_MODE_VIDEO"              // 横版视频，封面图宽高比1.78（下限：1280 & 720，上限：2560 & 1440））,视频资源支持mp4、mpeg、3gp、avi文件格式，宽高比16:9，大小不超过1000M
	CreativeImageModeGif                CreativeImageMode = "CREATIVE_IMAGE_MODE_GIF"                // GIF图,宽高比(690, 388),大小不超过1.5M
	CreativeImageModeLargeVertical      CreativeImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"     // 大图竖图，宽高比0.56，大小1.5M以下，下限：720 & 1280，上限：1440 & 2560
	CreativeImageModeVideoVertical      CreativeImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"     // 竖版视频，封面图宽高比0.56（9:16），下限：720 & 1280，上限：1440 & 2560，视频资源支持mp4、mpeg、3gp、avi文件格式，大小不超过100M
	CreativeImageModeSearchAdImage      CreativeImageMode = "TOUTIAO_SEARCH_AD_IMAGE"                // 搜索大图 仅限搜索广告使用，下限：345 & 138，上限：690 & 276
	CreativeImageModeSearchAdSmallImage CreativeImageMode = "SEARCH_AD_SMALL_IMAGE"                  // 搜索小图 仅限搜索广告使用，下限：108 & 72，上限：432 & 288
	CreativeImageModeUnionSplash        CreativeImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"       // 	穿山甲开屏图片 仅限穿山甲开屏使用，下限：1080 & 1920，上限：2160 & 3840，比例0.56
	CreativeImageModeUnionSplashVideo   CreativeImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO" // 穿山甲开屏视频 仅限穿山甲开屏使用，宽高比（0.56，9:16）,视频码率≥516kbps,大小≤100M,分辨率≥720*1280,1s<时长<6s
	CreativeImageModeDisplayWindow      CreativeImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"     // 搜索橱窗，宽高比1，下限：109 & 109，上限：109 & 109
	CreativeImageModeTitle              CreativeImageMode = "MATERIAL_IMAGE_MODE_TITLE"              // 标题类型，非创意的素材类型，仅报表接口会区分此素材类型
	CreativeImageModeAwemeLive          CreativeImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"         // 直播画面类型

	LearningPhaseDefault     LearningPhase = "DEFAULT"      // 默认，不在学习期中
	LearningPhaseLearning    LearningPhase = "LEARNING"     // 学习中
	LearningPhaseLearned     LearningPhase = "LEARNED"      // 学习成功
	LearningPhaseLearnFailed LearningPhase = "LEARN_FAILED" // 学习失败

)

type PageInfo struct {
	Page        int `json:"page,omitempty"`         // 页数
	PageSize    int `json:"page_size,omitempty"`    // 页面大小
	TotalNumber int `json:"total_number,omitempty"` // 总数
	TotalPage   int `json:"total_page,omitempty"`   // 总页数
}
