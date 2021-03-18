package models

type AdCreativeAddReq struct {
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID
	AdId         int64 `json:"ad_id"`         // 广告计划ID，计划ID要属于广告主ID，且非删除计划

	// 投放位置
	InventoryType  []InventoryType `json:"inventory_type,omitempty"`  // 广告位置（按媒体指定位置），详见【附录-投放位置】，具体描述参考广告创意
	SmartInventory int             `json:"smart_inventory,omitempty"` // 优选广告位，0表示不使用优选，1表示使用，使用优选广告位的时候默认忽略inventory_type字段； 注意：该字段与scene_inventory不能同时传
	SceneInventory string          `json:"scene_inventory,omitempty"` // 场景广告位，详见【附录-场景广告位】，使用场景广告位时默认忽略inventory_type字段； 注意：该字段与smart_inventory不能同时传。

	// 制作创意
	// 程序化创意 - 非DPA
	// 程序化创意实际会按照传入的title_list和image_list进行组合，对于效果不好的组合无法通过审核，获取到的都是审核通过的创意
	CreativeMaterialMode string              `json:"creative_material_mode,omitempty"` // 创意方式，当值为"STATIC_ASSEMBLE"表示程序化创意，其他情况不传字段
	ProceduralPackageId  int64               `json:"procedural_package_id,omitempty"`  // 程序化创意包ID，可通过【查询程序化创意包】接口进行查询，程序化创意包与自定义素材（title_list和image_list）不能同时使用，否则会报错 仅支持程序化创意，头条文章、DPA推广类型暂不支持
	IsPresentedVideo     int                 `json:"is_presented_video,omitempty"`     // 自动生成视频素材，利用已上传的图片与视频生成更多优质的短视频素材：1（启用），0（不启用） 默认值: 0
	GenerateDerivedAd    int                 `json:"generate_derived_ad,omitempty"`    // 是否开启衍生计划，衍生计划详情介绍：1（启用），0（不启用） 默认值: 0
	ImageList            []CreativeImageInfo `json:"image_list,omitempty"`             // 素材信息，creative_material_mode为"STATIC_ASSEMBLE"时必填，字段说明见下表。最多包含12张图和10个视频。
	TitleList            []CreativeTitleInfo `json:"title_list,omitempty"`             // 标题信息，creative_material_mode为"STATIC_ASSEMBLE"时必填，字段说明见下表。最多包含10个标题。

	// 程序化创意 - DPA
	// CreativeMaterialMode string      `json:"creative_material_mode,omitempty"` // 创意方式，当值为"STATIC_ASSEMBLE"表示程序化创意，其他情况不传字段

	// 基础创意信息
	Source                     string `json:"source,omitempty"`                        // 广告来源，4-20个字符，当推广目的为非应用下载或者应用下载且download_type为"EXTERNAL_URL时"时必填
	IesCoreUserId              string `json:"ies_core_user_id,omitempty"`              // 品牌主页-推广抖音号，当传入此字段时表示开启抖音主页。广告视频将同步到您的主页下，在客户端点击广告头像将进入您的主页。创建后不可修改。
	IsFeedAndFavSee            int    `json:"is_feed_and_fav_see,omitempty"`           // 主页作品列表隐藏广告内容，默认值：0 允选值：0（不隐藏），1（隐藏）
	CreativeAutoGenerateSwitch int    `json:"creative_auto_generate_switch,omitempty"` // 是否开启自动生成素材，delivery_range为UNIVERSAL：通投智选时可填，默认值: 1 允许值: 0：不启用, 1：启用
	AppName                    string `json:"app_name,omitempty"`                      // 应用名，当广告计划的download_type为"DOWNLOAD_URL"时，4到20个字符，必填
	SubTitle                   string `json:"sub_title,omitempty"`                     // APP 副标题。仅推广目标为APP，4到24个字符，填写Android下载链接时可设置
	WebUrl                     string `json:"web_url,omitempty"`                       // Android应用下载详情页（用户点击广告中“立即下载”按钮以外的区域时所到达的页面），当广告计划app_type为"APP_ANDROID"时, 必填
	ActionText                 string `json:"action_text,omitempty"`                   // 行动号召（仅应用下载推广类型有效） 备注：应用下载的行动号召字段使用action_text，门店与销售线索行动号召使用button_text 请求值可从接口【行动号召字段内容获取】进行获取，如果不传参默认为立即下载
	PlayableUrl                string `json:"playable_url,omitempty"`                  // 试玩素材URL，可通过【获取试玩素材列表】进行获取。 只有穿山甲激励视频可以使用试玩素材，同时素材需要审核通过
	IsCommentDisable           int    `json:"is_comment_disable,omitempty"`            // 是否关闭评论，0为开启，1为关闭，默认值：0 允许值: 0, 1

	// 附加创意
	AdvancedCreativeType  AdvancedCreativeType   `json:"advanced_creative_type,omitempty"`  // 附加创意类型
	AdvancedCreativeTitle string                 `json:"advanced_creative_title,omitempty"` // 副标题，最多24个字符
	PhoneNumber           string                 `json:"phone_number,omitempty"`            // 电话号码。当附加创意类型为"ATTACHED_CREATIVE_PHONE"时必填
	ButtonText            string                 `json:"button_text,omitempty"`             // 按钮文本，即行动号召，当附加创意类型非"ATTACHED_CREATIVE_NONE"时填写，请求值可从接口【行动号召字段内容获取】进行获取
	FormUrl               string                 `json:"form_url,omitempty"`                // 表单提交链接。当附加创意类型为"ATTACHED_CREATIVE_FORM"时 必填，必须为今日头条建站地址：【查询已有表单列表】
	PromotionCard         *CreativePromotionCard `json:"promotion_card,omitempty"`

	// 创意分类
	ThirdIndustryId int64    `json:"third_industry_id,omitempty"` // 三级行业ID
	AdKeywords      []string `json:"ad_keywords,omitempty"`       // 创意标签。最多20个标签，且每个标签长度不超过10个字符

	// 监测链接
	TrackUrl                   string `json:"track_url,omitempty"`                      // 展示（监测链接）
	ActionTrackUrl             string `json:"action_track_url,omitempty"`               // 点击（监测链接）（当推广目的为应用下载且创建计划传递了convert_id，系统会自动获取转化中的点击监测链接，且不可修改）
	VideoPlayEffectiveTrackUrl string `json:"video_play_effective_track_url,omitempty"` // 视频有效播放（监测链接），投放范围为穿山甲时暂不支持设置此链接
	VideoPlayDoneTrackUrl      string `json:"video_play_done_track_url,omitempty"`      // 视频播完（监测链接），投放范围为穿山甲时暂不支持设置此链接
	VideoPlayTrackUrl          string `json:"video_play_track_url,omitempty"`           // 视频播放（监测链接），投放范围为穿山甲时暂不支持设置此链接
	TrackUrlSendType           string `json:"track_url_send_type,omitempty"`            // 数据发送方式，不可修改,默认值: SERVER_SEND 允许值: SERVER_SEND(服务器端上传), CLIENT_SEND(客户端上传)

}

type CreativeImageInfo struct {
	ImageMode CreativeImageMode `json:"image_mode,omitempty"` // 素材类型，必填，注意：程序化创意不支持组图 CREATIVE_IMAGE_MODE_GROUP，其他类型图片都支持，如横版/竖版大图、小图。详见【附录-素材类型】
	//ImageMode string   `json:"image_mode,omitempty"` //

	ImageId  string   `json:"image_id,omitempty"`  // 图片ID，视频封面，视频素材时填写。可通过【获取图片素材】接口获得
	VideoId  string   `json:"video_id,omitempty"`  // 视频ID，视频素材时填写。可通过【获取视频素材】接口获得
	ImageIds []string `json:"image_ids,omitempty"` // 图片ID列表，非视频素材时填写。图片ID和视频ID可通过【获取图片素材】接口获得。组图类型传3张图，其他图片类型传1张，否则会报错。图片大小不能超过1.5M
}

type CreativeTitleInfo struct {
	Title           string  `json:"title,omitempty"`             // 创意标题，如果要使用动态词包，格式：“XXX{词包名}XXX{词包名}XXX”。请注意当您使用动态词包时，需在creative_word_ids字段中按顺序传入词包ID，并且在一个标题中最多使用两个动态词包。长度为5-30个字, 两个英文字符占1位。
	CreativeWordIds []int64 `json:"creative_word_ids,omitempty"` // 动态词包ID，最多支持两个词包。可使用【查询动态创意词包】获得，结合标题中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与标题中词包名顺序不一致我们将以词包ID顺序为准
}

type AdCreativeAddRsp struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	Data      AdCreativeAddRspData `json:"data"`
	RequestId string               `json:"request_id"`
}

type AdCreativeAddRspData struct {
	AdvertiserId             int64           `json:"advertiser_id"`
	AdId                     int64           `json:"ad_id"`
	ModifyTime               string          `json:"modify_time"`
	InventoryType            []InventoryType `json:"inventory_type"`             // 创意投放位置
	SmartInventory           int             `json:"smart_inventory"`            // 是否使用优选广告位，
	SceneInventory           string          `json:"scene_inventory"`            // 场景广告位
	CreativeMaterialMode     string          `json:"creative_material_mode"`     // 创意类型，该字段为STATIC_ASSEMBLE表示程序化创意，其他情况无该字段
	ProceduralPackageId      int64           `json:"procedural_package_id"`      // 程序化创意包ID
	ProceduralPackageVersion int64           `json:"procedural_package_version"` // 程序化创意包版本
	IsPresentedVideo         int             `json:"is_presented_video"`         // 	启用图片生成视频，允许值：0（不启用），1（启用）
	GenerateDerivedAd        int             `json:"generate_derived_ad"`        // 是否开启衍生计划，1为开启，0为不开启
	ImageList                []struct {
		ImageMode   string   `json:"image_mode"`   // 素材类型
		ImageId     string   `json:"image_id"`     // 图片ID，视频封面
		VideoId     string   `json:"video_id"`     // 视频ID
		ImageIds    []string `json:"image_ids"`    // 图片ID列表
		TemplateIds []int64  `json:"template_ids"` // 模版ID列表
	} `json:"image_list"` // 素材信息，程序化创意素材列表。最多包含12张图和6个视频。

	TitleList []struct {
		Title           string  `json:"title"`             // 创意标题
		CreativeWordIds []int64 `json:"creative_word_ids"` // 动态词包ID，可使用动态词包查询接口查询数据
		DpaDictIds      []int64 `json:"dpa_dict_ids"`      // DPA词包ID列表，针对DPA广告
	} `json:"title_list"` // 标题信息，程序化创意标题列表。最多包含10个标题
	Creatives []struct {
		CreativeId           int64    `json:"creative_id"`             // 创意ID。当类型为程序化创意时，没有审核通过前，该字段为null
		ImageMode            string   `json:"image_mode"`              // 素材类型
		Title                string   `json:"title"`                   // 创意标题
		CreativeWordIds      []int64  `json:"creative_word_ids"`       // 动态词包ID
		ImageIds             []string `json:"image_ids"`               // 图片ID列表，图片素材返回
		ImageId              string   `json:"image_id"`                // 图片ID，视频封面，视频素材时返回
		VideoId              string   `json:"video_id"`                // 视频ID，视频素材时返回
		DerivePosterCid      int      `json:"derive_poster_cid"`       // 是否将视频的封面和标题同步到图片创意，1为开启，0为不开启。视频素材时返回
		ThirdPartyId         string   `json:"third_party_id"`          // 创意自定义参数，例如开发者可设定此参数为创意打标签，用于区分使用的素材
		DpaDictIds           []int64  `json:"dpa_dict_ids"`            // DPA词包ID列表，针对DPA广告
		TemplateId           int64    `json:"template_id"`             // DPA模板ID，针对DPA广告
		TemplateImageId      string   `json:"template_image_id"`       // DPA创意实际显示的图片ID，针对DPA广
		DpaTemplate          int64    `json:"dpa_template"`            // 是否使用商品库视频模板，针对DPA广告
		DpaVideoTemplateType string   `json:"dpa_video_template_type"` // 商品库视频模板生成类型，针对DPA广告
		DpaVideoTaskIds      []string `json:"dpa_video_task_ids"`      // 自定义商品库视频模板ID，针对DPA广告
	} `json:"creatives"` // 素材信息, 投放位置和创意类型决定素材规格。程序化创意只有在审核通过后才有值
	Source                     string `json:"source"`                        // 广告来源
	IesCoreUserId              string `json:"ies_core_user_id"`              // 广告主绑定的抖音ID
	IsFeedAndFavSee            int    `json:"is_feed_and_fav_see"`           // 是否隐藏抖音主页，0：不隐藏，1：隐藏
	CreativeAutoGenerateSwitch int    `json:"creative_auto_generate_switch"` // 是否开启自动生成素材，delivery_range为UNIVERSAL：通投智选时返回，0：不启用,1：启用
	AppName                    string `json:"app_name"`                      // 应用名
	SubTitle                   string `json:"sub_title"`                     // 	APP 副标题。
	WebUrl                     string `json:"web_url"`                       // 	Android应用下载详情页
	ActionText                 string `json:"action_text"`                   // 	行动号召
	PlayableUrl                string `json:"playable_url"`                  // 试玩素材URL
	IsCommentDisable           int    `json:"is_comment_disable"`            // 是否关闭评论 允许值: 0, 1
	CloseVideoDetail           int    `json:"close_video_detail"`            // 是否关闭视频详情页落地页(勾选该选项后,视频详情页中不默认弹出落地页,仅对视频广告生效) 允许值: 0, 1
	CreativeDisplayMode        string `json:"creative_display_mode"`         // 创意展现方式
	AdvancedCreativeType       string `json:"advanced_creative_type"`        // 附加创意类型
	AdvancedCreativeTitle      string `json:"advanced_creative_title"`       // 附加创意副标题
	PhoneNumber                string `json:"phone_number"`                  // 电话号码(当附加创意类型为ATTACHED_CREATIVE_PHONE时返回)
	ButtonText                 string `json:"button_text"`                   // 按钮文本(当附加创意类型不为ATTACHED_CREATIVE_NONE时返回)
	FormUrl                    string `json:"form_url"`                      // 表单提交链接(当附加创意类型为ATTACHED_CREATIVE_FORM时返回)

	CommerceCards              []CreativeCommerceCards `json:"commerce_cards"`                 // 电商卡片信息。如果没有启用，那么不返回相关字段。
	ThirdIndustryId            int64                   `json:"third_industry_id"`              // 三级行业ID
	AdKeywords                 []string                `json:"ad_keywords"`                    // 创意标签
	TrackUrl                   string                  `json:"track_url"`                      // 展示（监测链接）
	ActionTrackUrl             string                  `json:"action_track_url"`               // 点击（监测链接）
	VideoPlayEffectiveTrackUrl string                  `json:"video_play_effective_track_url"` // 视频有效播放（监测链接）
	VideoPlayDoneTrackUrl      string                  `json:"video_play_done_track_url"`      // 视频播完（监测链接）
	VideoPlayTrackUrl          string                  `json:"video_play_track_url"`           // 视频播放（监测链接）
	TrackUrlSendType           string                  `json:"track_url_send_type"`            // 点击监测链接上报方式,默认值: SERVER_SEND 允许值:SERVER_SEND(服务器端上传), CLIENT_SEND(客户端上传)
	PromotionCard              CreativePromotionCard   `json:"promotion_card"`                 // 商品推广卡片
}

type CreativeCommerceCards struct {
	Title                string `json:"title"`                  // 标题
	StartTime            string `json:"start_time"`             // 投放开始时间，形式如：2017-01-01 00:00:00
	EndTime              string `json:"end_time"`               // 投放结束时间，形式如：2017-01-01 00:00:00
	ButtonText           string `json:"button_text"`            // 按钮文案
	Source               string `json:"source"`                 // 广告源（如店铺名称等）
	AdvancedDurationType string `json:"advanced_duration_type"` // 卡片出现时间类型
	AdvancedDuration     int    `json:"advanced_duration"`      // 卡片出现时间
	Type                 string `json:"type"`                   // 组件类型
	ExternalUrl          string `json:"external_url"`           // 落地页链接
	ImageInfo            struct {
		Width  string `json:"width"`   // 宽度
		Height string `json:"height"`  // 高度
		WebUri string `json:"web_uri"` // 链接
	} `json:"image_info"` // 素材信息
}
type CreativePromotionCard struct {
	EnableStorePack      bool     `json:"enable_store_pack"`                // 是否使用门店包，true为使用，false为不使用，推广目的非门店推广时会忽略该字段。若选择使用，则卡片标题为{最近门店名称}
	ProductSellingPoints []string `json:"product_selling_points,omitempty"` // 商品卖点
	ProductDescription   string   `json:"product_description,omitempty"`    // 商品描述
	CallToAction         string   `json:"call_to_action,omitempty"`         // 行动号召
	EnablePersonalAction bool     `json:"enable_personal_action,omitempty"` // 是否使用智能优选，true为使用，false为不使用
	ProductImageId       string   `json:"product_image_id,omitempty"`       // 商品图片ID
}
