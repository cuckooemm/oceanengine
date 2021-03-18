package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"probe_material_plan/marketing/oceanengine/models"
)

type AdvertiserApiService service

func (a AdvertiserApiService) GetDailyBudget(ctx context.Context, advIds []int64) (models.AdvertiserDailyBudgetRspData, http.Header, error) {
	var (
		apiPath     = a.client.Cfg.BasePath + "/advertiser/budget/get/"
		queryParams = url.Values{}
		rspBody     []byte
		req         *http.Request
		rsp         *http.Response
		result      models.AdvertiserDailyBudgetRsp
		err         error
	)
	queryParams.Add("advertiser_ids", fmt.Sprintf("%v", advIds))
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
func (a *AdvertiserApiService) BudgetUpdate(ctx context.Context, params models.AdvertiserBudgetUpdateReq) (http.Header, error) {
	var (
		apiPath      = a.client.Cfg.BasePath + "/advertiser/update/budget/"
		headerParams = make(map[string]string)
		postBody     []byte
		rspBody      []byte
		req          *http.Request
		rsp          *http.Response
		result       models.AdvertiserBudgetUpdateRsp
		err          error
	)
	postBody, _ = json.Marshal(params)
	headerParams["Content-Type"] = ApplicationJson
	if req, err = a.client.prepareRequest(ctx, apiPath, http.MethodPost, postBody, headerParams, nil); err != nil {
		return nil, err
	}
	if rsp, err = a.client.callAPI(req); err != nil || rsp == nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rspBody, err = ioutil.ReadAll(rsp.Body); err != nil {
		return rsp.Header, err
	}
	if rsp.StatusCode < 300 {
		if err = a.client.decode(&result, rspBody, rsp.Header.Get("Content-Type")); err != nil {
			return rsp.Header, NewApiSwaggerError(0, rspBody, err.Error(), "")
		}
		if result.Code != 0 {
			return rsp.Header, NewApiSwaggerError(result.Code, rspBody, result.Message, result.RequestId)
		}
		return rsp.Header, nil
	}
	return rsp.Header, NewApiSwaggerError(rsp.StatusCode, rspBody, rsp.Status, "")
}
