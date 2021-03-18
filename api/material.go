package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"probe_material_plan/marketing/oceanengine/models"
	"strconv"
)

type MaterialApiService service

// 通过此接口，用户可以上传和广告相关的素材图片，例如创意素材。
// 图片格式：jpg、jpeg、png、bmp、gif，大小1.5M内
// 注意
// 上传的图片一定要符合格式，才会在投放平台素材库显示！
// 若同一素材已进行上传，重新上传不会改名！
// https://ad.oceanengine.com/openapi/doc/index.html?id=331
func (m *MaterialApiService) UploadImage(ctx context.Context, params models.MaterialImageAddReq) (models.MaterialImageAddRspData, http.Header, error) {
	var (
		apiPath       = m.client.Cfg.BasePath + "/file/image/ad/"
		headerParams  = make(map[string]string)
		rspBody       []byte
		req           *http.Request
		rsp           *http.Response
		result        models.MaterialImageAddRsp
		body          = &bytes.Buffer{}
		w             = multipart.NewWriter(body)
		part          io.Writer
		isSetFilename = true
		err           error
	)
	if len(params.Filename) == 0 {
		if params.ImageFile != nil {
			params.Filename = params.ImageFile.Name()
			isSetFilename = false
		}
	}
	_ = w.WriteField("advertiser_id", strconv.FormatInt(params.AdvertiserId, 10))
	_ = w.WriteField("upload_type", string(params.UploadType))
	if isSetFilename && len(params.Filename) != 0 {
		_ = w.WriteField("filename", params.Filename)
	}
	w.Boundary()
	switch params.UploadType {
	case models.UploadByUrl:
		_ = w.WriteField("image_url", params.ImageUrl)
	case models.UploadByFile:
		if part, err = w.CreateFormFile("image_file", params.ImageFile.Name()); err != nil {
			return result.Data, nil, err
		}
		if _, err = io.Copy(part, params.ImageFile); err != nil {
			return result.Data, nil, err
		}
		_ = w.WriteField("image_signature", params.ImageSignature)
	}
	headerParams["Content-Type"] = w.FormDataContentType()
	headerParams["Content-Length"] = strconv.Itoa(body.Len())
	_ = w.Close()
	if req, err = m.client.prepareRequestAddFile(ctx, apiPath, body, headerParams); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = m.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, rsp.Header, err
	}
	if rsp.StatusCode < 300 {
		if err = m.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

// 通过此接口，用户可以上传和广告相关的素材视频。
// 视频格式：mp4、mpeg、3gp、avi
// 注意
// 视频设置了10s限制，超时会报错！
// 若同一素材已进行上传，重新上传不会改名！
// https://ad.oceanengine.com/openapi/doc/index.html?id=332
func (m *MaterialApiService) UploadVideo(ctx context.Context, params models.MaterialVideoAddReq) (models.MaterialVideoAddRspData, http.Header, error) {
	var (
		apiPath       = m.client.Cfg.BasePath + "/file/video/ad/"
		headerParams  = make(map[string]string)
		rspBody       []byte
		req           *http.Request
		rsp           *http.Response
		result        models.MaterialVideoAddRsp
		body          = &bytes.Buffer{}
		w             = multipart.NewWriter(body)
		part          io.Writer
		isSetFilename = true
		err           error
	)
	if len(params.Filename) == 0 {
		params.Filename = params.VideoFile.Name()
		isSetFilename = false
	}
	_ = w.WriteField("advertiser_id", strconv.FormatInt(params.AdvertiserId, 10))
	_ = w.WriteField("video_signature", params.VideoSignature)
	if isSetFilename && len(params.Filename) != 0 {
		_ = w.WriteField("filename", params.Filename)
	}
	w.Boundary()
	if part, err = w.CreateFormFile("video_file", params.VideoFile.Name()); err != nil {
		return result.Data, nil, err
	}
	if _, err = io.Copy(part, params.VideoFile); err != nil {
		return result.Data, nil, err
	}

	headerParams["Content-Type"] = w.FormDataContentType()
	headerParams["Content-Length"] = strconv.Itoa(body.Len())
	_ = w.Close()
	if req, err = m.client.prepareRequestAddFile(ctx, apiPath, body, headerParams); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = m.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, rsp.Header, err
	}
	if rsp.StatusCode < 300 {
		if err = m.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

// 通过此接口，用户可以进行同主体下不同广告主间的素材的推送。也就是说，将A广告主素材推送到，与A广告主主体（公司）相同的广告主。
// 注意
// 推送后素材的名称不会改变，将使用推送的原素材名！
// 新上传素材存在同步延迟情况，建议等待2-3分钟再尝试操作推送！
// 当素材已存在待推送的广告主的素材库内时，不会重复推送，推送失败的结果会在推送失败列表展示！
// https://ad.oceanengine.com/openapi/doc/index.html?id=1458
func (m *MaterialApiService) PushMaterial(ctx context.Context, params models.MaterialPushReq) (models.MaterialPushRspData, http.Header, error) {
	var (
		apiPath      = m.client.Cfg.BasePath + "/file/material/bind/"
		headerParams = make(map[string]string)
		postBody     []byte
		rspBody      []byte
		req          *http.Request
		rsp          *http.Response
		result       models.MaterialPushRsp
		err          error
	)
	postBody, _ = json.Marshal(params)
	headerParams["Content-Type"] = ApplicationJson
	if req, err = m.client.prepareRequest(ctx, apiPath, http.MethodPost, postBody, headerParams, nil); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = m.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, rsp.Header, err
	}
	if rsp.StatusCode < 300 {
		if err = m.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

// 通过此接口，用户可以获取针对素材视频推荐的智能封面。智能封面是通过提取视频关键帧筛选出推荐封面，帮助发现视频内优质封面素材。
// 推荐封面图片的数量是1-13个，对于相似度极高的封面图片会进行去重等处理，由实际的视频内容和时长决定。
// 注意
// 智能封面不是实时获取，而需要先根据status判断封面获取的状态，然后再进行获取视频封面！
// 新上传素材存在同步延迟情况，建议等待2-3分钟再尝试操作获取视频智能封面！
// https://ad.oceanengine.com/open_api/2/tools/video_cover/suggest/
func (m *MaterialApiService) GetVideoCover(ctx context.Context, advId int64, videoId string) (models.MaterialVideoCoverGetRspData, http.Header, error) {
	var (
		apiPath     = m.client.Cfg.BasePath + "/tools/video_cover/suggest/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.MaterialVideoCoverGetRsp
		err         error
	)
	queryParams.Add("advertiser_id", strconv.FormatInt(advId, 10))
	queryParams.Add("video_id", videoId)
	if req, err = m.client.prepareRequest(ctx, apiPath, http.MethodGet, nil, nil, queryParams); err != nil {
		return result.Data, nil, err
	}

	if rsp, err = m.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, rsp.Header, err
	}
	if rsp.StatusCode < 300 {
		if err = m.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

// 通过此接口，用户可以获取经过一定条件过滤后的广告主下创意素材库对应的视频及视频信息。
// 注意
// 为保证接口使用的安全性，避免调取他人的图片信息，该接口针对素材URL的字段仅查询自己广告主下的素材才会返回，即需查询的广告主账号的主体需与APPID对应开发者的主体保持一致，才可获取到素材的预览URL的信息，否则会提示：“素材所属主体与开发者主体不一致无法获取URL”(第三方获取敏感物料信息可在授权时申请广告主授权敏感物料权限，可参考常见问题【敏感物料授权】）！
// 对素材视频进行过滤的时候，video_ids（视频ID）、material_ids（素材ID）、signatures（视频的md5值）只能选择一个进行过滤！
// 获取视频素材数据目前仅支持10000个！
// https://ad.oceanengine.com/openapi/doc/index.html?id=351
func (m *MaterialApiService) GetVideo(ctx context.Context, advId int64, opts models.MaterialVideoGetOpts) (models.MaterialVideoGetRspData, http.Header, error) {
	var (
		apiPath     = m.client.Cfg.BasePath + "/file/video/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.MaterialVideoGetRsp
		err         error
	)
	queryParams.Add("advertiser_id", strconv.FormatInt(advId, 10))
	if opts.Filtering.IsSet() {
		queryParams.Add("filtering", parameterToString(opts.Filtering.Value(), "multi"))
	}
	if opts.Page.IsSet() {
		queryParams.Add("page", strconv.FormatInt(opts.Page.Value(), 10))
	}
	if opts.PageSize.IsSet() {
		queryParams.Add("page_size", strconv.FormatInt(opts.PageSize.Value(), 10))
	}

	if req, err = m.client.prepareRequest(ctx, apiPath, http.MethodGet, nil, nil, queryParams); err != nil {
		return result.Data, nil, err
	}

	if rsp, err = m.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, rsp.Header, err
	}
	if rsp.StatusCode < 300 {
		if err = m.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}
