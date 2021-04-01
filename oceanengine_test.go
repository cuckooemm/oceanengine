package oceanengine

import (
	"context"
	"github.com/antihax/optional"
	"github.com/cuckooemm/oceanengine/ads"
	"github.com/cuckooemm/oceanengine/api"
	"github.com/cuckooemm/oceanengine/models"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	adClient := ads.Init(true)
	adClient.UseProduction()
	adClient.Cfg.SetAccessToken("token")
	adClient.Cfg.SetUserAgent("golang sdk")
	apiLog, _ := os.OpenFile("api.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	defer apiLog.Close()
	adClient.AppendMiddleware(ads.NewLogMiddleware(adClient, apiLog))

	param := models.AdAdGetOpts{}
	param.PageSize = optional.NewInt64(100)
	rsp, header, err := adClient.Client.AdApi.Get(context.Background(), 123456789, param)
	if err != nil {
		e := err.(api.SwaggerError)
		t.Log(e.Code())
		t.Log(e.Body())
		t.Log(e.Error())
		t.Log(e.Message())
		t.Log(e.RequestId())
	}
	t.Logf("response: %+v\n", rsp)
	t.Logf("headers: %+v\n", header)
}
