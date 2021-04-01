package models

import (
	"github.com/antihax/optional"
	"os"
)

type PushFailReason string

const (
	VideoBindingExisted PushFailReason = "VIDEO_BINDING_EXISTED" // 视频已存在
	ImageBindingExisted PushFailReason = "IMAGE_BINDING_EXISTED" // 图片已存在
)

func (p PushFailReason) String() string {
	switch p {
	case VideoBindingExisted:
		return "视频已存在"
	case ImageBindingExisted:
		return "图片已存在"
	default:
		return "unknown"
	}
}

type MaterialImageAddReq struct {
	AdvertiserId   int64      `json:"advertiser_id,omitempty"`   // 广告主ID
	UploadType     UploadType `json:"upload_type,omitempty"`     // 图片上传方式
	ImageSignature string     `json:"image_signature,omitempty"` // 图片的md5值(用于服务端校验) upload_type为UPLOAD_BY_FILE必填
	ImageFile      *os.File   `json:"image_file,omitempty"`      // 图片文件 upload_type为UPLOAD_BY_FILE必填 格式: jpg、jpeg、png、bmp、gif, 大小1.5M内
	ImageUrl       string     `json:"image_url,omitempty"`       // 图片url地址，如http://xxx.xxx upload_type为UPLOAD_BY_URL必填
	Filename       string     `json:"filename,omitempty"`        // 素材的文件名，可自定义素材名，不传择默认取文件名，最长255个字符 注：若同一素材已进行上传，重新上传不会改名
}

type MaterialImageAddRsp struct {
	Code      int                     `json:"code,omitempty"`
	Message   string                  `json:"message,omitempty"`
	Data      MaterialImageAddRspData `json:"data,omitempty"`
	RequestId string                  `json:"request_id,omitempty"`
}

type MaterialImageAddRspData struct {
	Id         string `json:"id,omitempty"`          // 图片id
	Size       int64  `json:"size,omitempty"`        // 图片大小
	Width      int64  `json:"width,omitempty"`       // 图片宽度
	Height     int64  `json:"height,omitempty"`      // 图片高度
	Url        string `json:"url,omitempty"`         // 图片预览地址(1小时内有效)
	Format     string `json:"format,omitempty"`      // 图片格式
	Signature  string `json:"signature,omitempty"`   // 图片md5
	MaterialId int64  `json:"material_id,omitempty"` // 素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
}

type MaterialVideoAddReq struct {
	AdvertiserId   int64    `json:"advertiser_id,omitempty"`   // 广告主ID
	VideoSignature string   `json:"video_signature,omitempty"` // 视频的md5值(用于服务端校验)
	VideoFile      *os.File `json:"video_file,omitempty"`      // 视频文件 允许格式：mp4、mpeg、3gp、avi（10s超时限制）
	Filename       string   `json:"filename,omitempty"`        // 素材的文件名，可自定义素材名，不传择默认取文件名，最长255个字符 注：若同一素材已进行上传，重新上传不会改名
}

type MaterialVideoAddRsp struct {
	Code      int                     `json:"code,omitempty"`
	Message   string                  `json:"message,omitempty"`
	Data      MaterialVideoAddRspData `json:"data,omitempty"`
	RequestId string                  `json:"request_id,omitempty"`
}

type MaterialVideoAddRspData struct {
	VideoId    string  `json:"video_id,omitempty"`
	Size       int64   `json:"size,omitempty"`
	Width      int64   `json:"width,omitempty"`
	Height     int64   `json:"height,omitempty"`
	VideoUrl   string  `json:"video_url,omitempty"`
	Duration   float64 `json:"duration,omitempty"`
	MaterialId int64   `json:"material_id,omitempty"` // 素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
}

type MaterialPushReq struct {
	AdvertiserId        int64    `json:"advertiser_id,omitempty"`
	TargetAdvertiserIds []int64  `json:"target_advertiser_ids,omitempty"`
	VideoIds            []string `json:"video_ids,omitempty"`
	ImageIds            []string `json:"image_ids,omitempty"`
}

type MaterialPushRsp struct {
	Code      int                 `json:"code,omitempty"`
	Message   string              `json:"message,omitempty"`
	Data      MaterialPushRspData `json:"data,omitempty"`
	RequestId string              `json:"request_id,omitempty"`
}

type MaterialPushRspData struct {
	FailList []MaterialPushRspDataFailList `json:"fail_list,omitempty"`
}

type MaterialPushRspDataFailList struct {
	VideoId            string         `json:"video_id,omitempty"`
	ImageId            string         `json:"image_id,omitempty"`
	TargetAdvertiserId int64          `json:"target_advertiser_id,omitempty"`
	FailReason         PushFailReason `json:"fail_reason,omitempty"`
}

type MaterialVideoCoverGetRsp struct {
	Code      int                          `json:"code,omitempty"`
	Message   string                       `json:"message,omitempty"`
	Data      MaterialVideoCoverGetRspData `json:"data,omitempty"`
	RequestId string                       `json:"request_id,omitempty"`
}

type MaterialVideoCoverGetRspData struct {
	Status VideoCoverStatus         `json:"status,omitempty"`
	List   []MaterialVideoCoverList `json:"list,omitempty"`
}

type MaterialVideoCoverList struct {
	Id     string `json:"id,omitempty"`
	Width  int64  `json:"width,omitempty"`
	Height int64  `json:"height,omitempty"`
	Url    string `json:"url,omitempty"`
}

type MaterialVideoGetOpts struct {
	Filtering optional.Interface `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	Page      optional.Int64     `json:"page,omitempty"`      // 当前页码: 1
	PageSize  optional.Int64     `json:"page_size,omitempty"` // 页面大小 默认值: 10， 取值范围：1-1000
}

type MaterialVideoGetFilter struct {
	Width       int64    `json:"width,omitempty"`
	Height      int64    `json:"height,omitempty"`
	Ratio       float64  `json:"ratio,omitempty"`
	VideoIds    []string `json:"video_ids,omitempty"`
	MaterialIds []int64  `json:"material_ids,omitempty"`
	Signatures  []string `json:"signatures,omitempty"`
	StartTime   string   `json:"start_time,omitempty"` // 根据视频上传时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-mm-dd
	EndTime     string   `json:"end_time,omitempty"`
}

type MaterialVideoGetRsp struct {
	Code      int                     `json:"code,omitempty"`
	Message   string                  `json:"message,omitempty"`
	Data      MaterialVideoGetRspData `json:"data,omitempty"`
	RequestId string                  `json:"request_id,omitempty"`
}

type MaterialVideoGetRspData struct {
	List     []MaterialVideoGetRspDataList `json:"list,omitempty"`
	PageInfo PageInfo                      `json:"page_info,omitempty"`
}

type MaterialVideoGetRspDataList struct {
	Id         string  `json:"id,omitempty"`
	Size       int64   `json:"size,omitempty"`
	Width      int64   `json:"width,omitempty"`
	Height     int64   `json:"height,omitempty"`
	Url        string  `json:"url,omitempty"`
	Format     string  `json:"format,omitempty"`
	Signature  string  `json:"signature,omitempty"`
	PosterUrl  string  `json:"poster_url,omitempty"` // 视频首帧截图，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”，链接1小时过期
	BitRate    int64   `json:"bit_rate,omitempty"`   // 码率，单位bps
	Duration   float64 `json:"duration,omitempty"`   // 视频时长
	MaterialId int64   `json:"material_id,omitempty"`
	Source     string  `json:"source,omitempty"`      // 视频来源
	CreateTime string  `json:"create_time,omitempty"` // 素材的上传时间，格式："yyyy-mm-dd HH:MM:SS"
	Filename   string  `json:"filename,omitempty"`
}
