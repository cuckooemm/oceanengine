package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"probe_material_plan/marketing/oceanengine/models"
)

type AdCreativeApiService service

// 每个计划下程序化创意和自定义创意为二选一，且无法修改；
// 程序化创意: 最多10个标题、12个图片素材和10个视频素材;如果创建的是程序化创意（程序化创意实际会按照传入的title_list和image_list进行组合，对于效果不好的组合无法通过审核，获取到的都是审核通过的创意），只有在审核之后才会获取到创意数据与创意id；
// 自定义创意API当前最多支持30个创意；
// 每日最多创建500个创意（自定义创意+程序化创意）；
// 素材类型：不同广告未要求素材类型不同,其中每一种素材类型都有自己的规格，请传入符合要求的素材，否则会报错！
// 其中视频的时长需要>=4s，否则会报错！
// 监测链接：当在计划纬度设置了转化id，如果在创建创意时不传监测链接，会自动获取转化id里监测链接；如果在创建（更新）创意时传入对应的监测链接，会取传入的监测链接，但是对于应用下载推广，即便主动传入点击监测链接，也会取转化id监测链接；
// 对于不打算传的字段，不要传“”或者null，传了会校验！
// 如果计划ID下已有创意信息，需要使用update_v2接口进行修改或者新增创意素材，否则会报错！
// API不支持功能：
// 高级创意：预计10月份支持；

// https://ad.oceanengine.com/openapi/doc/index.html?id=519
func (a *AdCreativeApiService) Add(ctx context.Context, params models.AdCreativeAddReq) (models.AdCreativeAddRspData, http.Header, error) {
	var (
		apiPath      = a.client.Cfg.BasePath + "/creative/create_v2/"
		headerParams = make(map[string]string)
		postBody     []byte
		rspBody      []byte
		req          *http.Request
		rsp          *http.Response
		result       models.AdCreativeAddRsp
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
