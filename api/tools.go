package api

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"probe_material_plan/marketing/oceanengine/models"
	"strconv"
)

type ToolsApiService service

// 获取行业列表，通过接口可以获取到一级行业、二级行业、三级行业列表，其中代理商创建广告主时使用的是二级行业，而在创建创意填写创意分类时使用的是三级行业，请注意区分。
// https://ad.oceanengine.com/openapi/doc/index.html?id=370
func (t *ToolsApiService) GetIndustry(ctx context.Context, advId int64, opts models.ToolsIndustryGetOpts) ([]models.ToolsIndustryGetRspDataList, http.Header, error) {
	var (
		apiPath     = t.client.Cfg.BasePath + "/tools/industry/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.ToolsIndustryGetRsp
		err         error
	)
	if opts.Level.IsSet() {
		queryParams.Add("level", strconv.Itoa(opts.Level.Value()))
	}
	if opts.Type.IsSet() {
		queryParams.Add("type", opts.Type.Value())
	}
	if req, err = t.client.prepareRequest(ctx, apiPath, http.MethodGet, nil, nil, queryParams); err != nil {
		return result.Data.List, nil, err
	}
	if rsp, err = t.client.callAPI(req); err != nil || rsp == nil {
		return result.Data.List, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data.List, nil, err
	}
	if rsp.StatusCode < 300 {
		if err = t.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data.List, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data.List, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data.List, rsp.Header, nil
	}
	return result.Data.List, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

// 获取行动号召字段内容，注意：结合附加创意类型以及广告主行业参数可以查询出更多细纬度的行动号召内容。
// https://ad.oceanengine.com/openapi/doc/index.html?id=1366
func (t *ToolsApiService) GetActionText(ctx context.Context, advId int64, LandingType string, opts models.ToolsActionTextGetOpts) ([]string, http.Header, error) {
	var (
		apiPath     = t.client.Cfg.BasePath + "/tools/action_text/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.ToolsActionTextGetRsp
		err         error
	)
	queryParams.Add("advertiser_id", strconv.FormatInt(advId, 10))
	queryParams.Add("landing_type", LandingType)
	if opts.AdvancedCreativeType.IsSet() {
		queryParams.Add("advanced_creative_type", opts.AdvancedCreativeType.Value())
	}
	if opts.Industry.IsSet() {
		queryParams.Add("type", strconv.FormatInt(opts.Industry.Value(), 10))
	}
	if req, err = t.client.prepareRequest(ctx, apiPath, http.MethodGet, nil, nil, queryParams); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = t.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, nil, err
	}
	if rsp.StatusCode < 300 {
		if err = t.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

// 获取广告计划学习期状态。关于学习期
// 最多传100个广告计划id。
// https://ad.oceanengine.com/openapi/doc/index.html?id=1664566788147212
func (t *ToolsApiService) GetAdLearnStat(ctx context.Context, advId int64, adIds []int64) ([]models.ToolsAdStatExtraInfoRspData, http.Header, error) {
	var (
		apiPath     = t.client.Cfg.BasePath + "/tools/ad_stat_extra_info/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.ToolsADStatExtraInfoRsp
		err         error
	)
	queryParams.Add("advertiser_id", strconv.FormatInt(advId, 10))
	queryParams.Add("landing_type", parameterToString(adIds, ""))
	if req, err = t.client.prepareRequest(ctx, apiPath, http.MethodGet, nil, nil, queryParams); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = t.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, nil, err
	}
	if rsp.StatusCode < 300 {
		if err = t.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

// 查询计划的广告质量度，只有产生过投放消耗的计划才会有质量度数据。
func (t *ToolsApiService) GetAdQuality(ctx context.Context, advId int64, adIds []int64) (models.ToolsAdQualityRspData, http.Header, error) {
	var (
		apiPath     = t.client.Cfg.BasePath + "/tools/ad_quality/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.ToolsAdQualityRsp
		err         error
	)
	queryParams.Add("advertiser_id", strconv.FormatInt(advId, 10))
	queryParams.Add("ad_ids", parameterToString(adIds, "multi"))
	if req, err = t.client.prepareRequest(ctx, apiPath, http.MethodGet, nil, nil, queryParams); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = t.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, nil, err
	}
	if rsp.StatusCode < 300 {
		if err = t.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}

// 查询计划的广告质量度，只有产生过投放消耗的计划才会有质量度数据。
func (t *ToolsApiService) GetAdDiagnosis(ctx context.Context, advId int64, adIds []int64) (models.ToolsAdDiagnosisRspData, http.Header, error) {
	var (
		apiPath     = t.client.Cfg.BasePath + "/tools/diagnosis/ad/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.ToolsAdDiagnosisRsp
		err         error
	)
	queryParams.Add("advertiser_id", strconv.FormatInt(advId, 10))
	queryParams.Add("ad_ids", parameterToString(adIds, "multi"))
	if req, err = t.client.prepareRequest(ctx, apiPath, http.MethodGet, nil, nil, queryParams); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = t.client.callAPI(req); err != nil || rsp == nil {
		return result.Data, nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return result.Data, nil, err
	}
	if rsp.StatusCode < 300 {
		if err = t.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return result.Data, rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}
