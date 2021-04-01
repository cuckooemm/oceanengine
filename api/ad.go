package api

import (
	"context"
	"encoding/json"
	"github.com/cuckooemm/oceanengine/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type AdApiService service

// 此接口用于创建广告计划，对于搜索广告的创建可参照【搜索广告投放】
//
// 计划包括了计划名称、投放范围、投放目标、用户定向、预算与出价，对于其中的概念解释可以参考：【广告计划】
// 计划创建涉及了多个资产管理：应用管理、落地页管理、转化目标管理、定向包管理、DMP人群管理等；开发者需要提前根据开放接口构建这些资产功能，以免创建计划卡住！；
// 如果创建计划遇到问题，可通过 常见问题 来解决

// 注意：
// 1. 单广告组可创建计划上限500个，且单账号每天创建计划上限500个；
// 2. 投放目标与API对应的字段是：pricing;
// 3. 一旦设置日预算或总预算，预算模式不允许修改，但金额可以调整。仅CPC、CPM,最低计划日预算100；CPA、OCPC、OCPM最低计划日预算是300。
// 4. 24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；
// 5. 单次修改预算幅度不能低于100元（增加或者减少）；
// 6. 修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；
// 7. 计划中同时使用自定义定向和定向包，优先且只使用定向包中的定向；
// 8. 受众参数不填（不填不代表填null）或者填固定的不限格式均可以设置对应受众为不限，string类型受众设置无限的常量统一可使用"NONE"，list类型的受众设置无限统一使用空数组，即[]。
// 9. 对于不打算传的字段，不要传“”或者null，传了会校验!!!
// 10. API目前不支持设置多转化目标；

func (a *AdApiService) Add(ctx context.Context, params models.AdAdAddReq) (models.AdAdAddRspData, http.Header, error) {
	var (
		apiPath      = a.client.Cfg.BasePath + "/ad/create/"
		headerParams = make(map[string]string)
		postBody     []byte
		rspBody      []byte
		req          *http.Request
		rsp          *http.Response
		result       models.AdAdAddRsp
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
			return result.Data, nil, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return result.Data, rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return result.Data, rsp.Header, nil
	}
	return result.Data, rsp.Header, NewApiSwaggerError(50000, rspBody, rsp.Status, "")
}

// 此接口用于获取广告计划列表的信息；
// 如果创建计划时未设置对应的字段，返回的字段值会是null
// 支持filtering过滤，可按广告计划ID、出价方式、广计划状态进行过滤筛选
// 默认不获取删除的计划，如果要获取删除的计划，可在filtering中传入对应的status值；
// 对于搜索广告计划信息获取参见【搜索广告投放】

func (a *AdApiService) Get(ctx context.Context, advId int64, opts models.AdAdGetOpts) (models.AdAdGetRspData, http.Header, error) {
	var (
		apiPath     = a.client.Cfg.BasePath + "/ad/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.AdAdGetRsp
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
	return result.Data, rsp.Header, NewApiSwaggerError(50000, rspBody, rsp.Status, "")
}

// 通过此接口用于更新广告计划的预算；
// 一次可以处理100个计划
// 24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；
// 单次修改预算幅度不能低于100元（增加或者减少）；
// 修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；
func (a *AdApiService) BudgetUpdate(ctx context.Context, advId int64, data []models.AdBudgetData) (models.AdBudgetUpdateRspData, http.Header, error) {
	var (
		apiPath      = a.client.Cfg.BasePath + "/ad/update/budget/"
		headerParams = make(map[string]string)
		postBody     []byte
		rspBody      []byte
		req          *http.Request
		rsp          *http.Response
		result       models.AdBudgetUpdateRsp
		err          error
	)
	postBody, _ = json.Marshal(models.AdBudgetUpdateReq{AdvertiserId: advId, Data: data})
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
	return result.Data, rsp.Header, NewApiSwaggerError(50000, rspBody, rsp.Status, "")
}

// 通过此接口可对计划做启用/暂停/删除操作；
// 一次可以处理100个计划
// 对于删除的计划不能再进行状态操作，否则会报错！
// 如果有一个计划有问题，全部计划修改都不会成功！请确保传入的计划属于此广告主以及处于非删除状态。
func (a *AdApiService) UpdateStatus(ctx context.Context, advId int64, adIds []int64, status models.AdStatus) (models.AdStatusUpdateRspData, http.Header, error) {
	var (
		apiPath      = a.client.Cfg.BasePath + "/ad/update/status/"
		headerParams = make(map[string]string)
		postBody     []byte
		rspBody      []byte
		req          *http.Request
		rsp          *http.Response
		result       models.AdStatusUpdateRsp
		err          error
	)
	postBody, _ = json.Marshal(models.AdStatusUpdateReq{AdvertiserId: advId, AdIds: adIds, OptStatus: status})
	headerParams["Content-Type"] = ApplicationJson
	if req, err = a.client.prepareRequest(ctx, apiPath, http.MethodPost, postBody, headerParams, nil); err != nil {
		return result.Data, nil, err
	}
	if rsp, err = a.client.callAPI(req); err != nil {
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
	return result.Data, rsp.Header, NewApiSwaggerError(50000, rspBody, rsp.Status, "")
}
