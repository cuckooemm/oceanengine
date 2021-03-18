package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"probe_material_plan/marketing/oceanengine/models"
	"strconv"
)

type AdCampaignApiService service

/*
  创建广告组
  doc https://ad.oceanengine.com/openapi/doc/index.html?id=295

  此接口用于创建信息流广告组，对于搜索广告的创建可参照【搜索广告投放】
  每个广告主账号下最多可允许创建500个广告组，如超出需要先删除一部分广告组后才可继续创建；
  当选择日预算类型时，日预算不少于300元；
  24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；
  单次修改预算幅度不能低于100元（增加或者减少）；
  修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；
*/
func (a *AdCampaignApiService) Add(ctx context.Context, params models.AdCampaignAddReq) (models.AdCampaignAddRspData, http.Header, error) {
	var (
		apiPath      = a.client.Cfg.BasePath + "/campaign/create/"
		headerParams = make(map[string]string)
		postBody     []byte
		rspBody      []byte
		req          *http.Request
		rsp          *http.Response
		result       models.AdCampaignAddRsp
		err          error
	)
	postBody, _ = json.Marshal(params)
	headerParams["Content-Type"] = ApplicationJson
	if req, err = a.client.prepareRequest(ctx, apiPath, http.MethodPost, postBody, headerParams, nil); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = a.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, rsp.Header, err
	}
	if rsp.StatusCode < 300 {
		if err = a.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

/*
  获取广告组
  doc https://ad.oceanengine.com/open_api/2/campaign/get/
	当预算类型为不限，返回的预算为0.0元
	支持filtering过滤，可按广告组ID、推广目的、广告组状态进行过滤筛选
	默认不获取删除的广告组，如果要获取删除的广告组，可在filtering中传入对应的status值；
	对于搜索广告组信息获取参见【搜索广告投放】
*/
func (a *AdCampaignApiService) Get(ctx context.Context, advId int64, opts models.AdCampaignGetOpts) (models.AdCampaignGetRspData, http.Header, error) {
	var (
		apiPath     = a.client.Cfg.BasePath + "/campaign/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.AdCampaignGetRsp
		err         error
	)
	queryParams.Add("advertiser_id", strconv.FormatInt(advId, 10))

	if opts.Filtering.IsSet() {
		queryParams.Add("filtering", parameterToString(opts.Filtering.Value(), "multi"))
	}
	if opts.Fields.IsSet() {
		queryParams.Add("fields", parameterToString(opts.Fields.Value(), "multi"))
	}
	if opts.Page.IsSet() {
		queryParams.Add("page", strconv.FormatInt(opts.Page.Value(), 10))
	}
	if opts.PageSize.IsSet() {
		queryParams.Add("page_size", strconv.FormatInt(opts.PageSize.Value(), 10))
	}
	if req, err = a.client.prepareRequest(ctx, apiPath, http.MethodGet, nil, nil, queryParams); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = a.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, rsp.Header, err
	}
	if rsp.StatusCode < 300 {
		if err = a.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}
