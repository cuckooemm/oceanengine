# oceanengine

巨量引擎 头条广告SDK

### Usage Examples
```golang
// 初始化
adClient := ads.Init(true)
adClient.UseProduction()
adClient.Cfg.SetAccessToken("token")
adClient.Cfg.SetUserAgent("golang sdk")
apiLog, _ := os.OpenFile("api.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
defer apiLog.Close()
adClient.AppendMiddleware(ads.NewLogMiddleware(adClient, apiLog))

param := models.AdAdGetOpts{}
param.PageSize = optional.NewInt64(100)
if rsp, header, err := adClient.Client.AdApi.Get(context.Background(), 123456789, param); err != nil {
    e := err.(api.SwaggerError)
    fmt.Printf("Code: \n",e.Code())
    fmt.Printf("RawBody: \n",e.Body())
    fmt.Printf("Message: \n",e.Message())
    fmt.Printf("RequestId: \n",e.RequestId())
    }
fmt.Printf("response: %+v\n",rsp)
fmt.Printf("headers: %+v\n",header)
```

### 已实现接口
- 广告投放
- [x] 广告账户预算
  - [x] 获取账户日预算
  - [x] 更新账户日预算
- [x] 广告组
  - [x] 获取广告组
  - [x] 创建广告组
  - [ ] 修改广告组
  - [ ] 广告组更新状态
- [x] 广告计划模块
  - [x] 获取广告计划
  - [x] 创建广告计划
  - [ ] 修改广告计划
  - [x] 更新计划状态
  - [x] 更新计划预算
  - [ ] 更新计划出价
  - [ ] 获取计划审核建议
- [x] 广告创意模块
    - [ ] 获取创意列表
    - [x] 创建广告创意
    - [ ] 创意详细信息
    - [ ] 修改创意信息
    - [ ] 更新创意状态
    - [ ] 创意素材信息
    - [ ] 获取创意审核建议