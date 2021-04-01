package ads

import (
	"fmt"
	"github.com/cuckooemm/oceanengine/api"
	"github.com/cuckooemm/oceanengine/conf"
	"net/http"
)

type AdClient struct {
	http.RoundTripper
	Cfg            *conf.AdConfig
	middlewareList []Middleware
	ApiVersion     string
	Client         *api.APIClient
}

func Init(debug bool) *AdClient {
	cfg := conf.NewAdConfig(debug)
	cli := &AdClient{
		Cfg:          cfg,
		RoundTripper: http.DefaultTransport,
		Client:       api.NewAPIClient(cfg),
	}
	cli.Client.Cfg.HTTPClient.Transport = cli
	cli.UseSandbox()
	return cli
}

func (c *AdClient) UseSandbox() {
	c.SetHost("https", "test-ad.toutiao.com", "open_api/2")
}
func (c *AdClient) UseProduction() {
	c.SetHost("https", "ad.oceanengine.com", "open_api/2")
}

func (c *AdClient) SetHost(schema, host, apiVersion string) {
	modified := false
	if len(host) != 0 {
		c.Client.Cfg.Host = host
		modified = true
	}
	if len(schema) != 0 {
		c.Client.Cfg.Scheme = schema
		modified = true
	}
	if len(apiVersion) != 0 {
		c.Client.Cfg.ApiVersion = apiVersion
		modified = true
	}
	if modified {
		c.Client.Cfg.BasePath = fmt.Sprintf("%s://%s/%s", c.Client.Cfg.Scheme, c.Client.Cfg.Host, c.Client.Cfg.ApiVersion)
	}
}

func (c *AdClient) AppendMiddleware(middleware Middleware) {
	c.middlewareList = append(c.middlewareList, middleware)
}

func (c *AdClient) RoundTrip(req *http.Request) (rsp *http.Response, err error) {
	beforeFunc := func(req *http.Request) (rsp *http.Response, err error) {
		return c.RoundTripper.RoundTrip(req)
	}
	// 逆序遍历
	for i := len(c.middlewareList) - 1; i >= 0; i-- {
		beforeFunc = c.GenMiddlewareHandleFunc(c.middlewareList[i], beforeFunc)
	}
	return beforeFunc(req)
}

// GenMiddlewareHandleFunc ...
func (c *AdClient) GenMiddlewareHandleFunc(middleware Middleware, beforeFunc func(req *http.Request) (rsp *http.Response, err error),
) func(req *http.Request) (rsp *http.Response, err error) {
	return func(req *http.Request) (rsp *http.Response, err error) {
		return middleware.Handle(req, beforeFunc)
	}
}

func (c *AdClient) IsDebug() bool {
	return c.Cfg.IsDebug
}

func (c *AdClient) SetDebug(b bool) {
	c.Cfg.IsDebug = b
}

func (c *AdClient) SetHeaders(header http.Header) {
	c.Cfg.GlobalConfig.HttpOption.Header = header
}

func (c *AdClient) SetHeader(key string, value string) {
	if c.Cfg.GlobalConfig.HttpOption.Header == nil {
		c.Cfg.GlobalConfig.HttpOption.Header = http.Header{}
	}
	c.Cfg.GlobalConfig.HttpOption.Header.Set(key, value)
}
