package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/cuckooemm/oceanengine/conf"
	"net/http"
	"net/url"
	"strings"
)

const ApplicationJson = "application/json"

type service struct {
	client *APIClient
}

type APIClient struct {
	Cfg    *conf.AdConfig
	common service

	// api service
	CampaignApi   *CampaignApiService
	AdApi         *AdApiService
	CreativeApi   *CreativeApiService
	MaterialApi   *MaterialApiService
	ToolsApi      *ToolsApiService
	AdvertiserApi *AdvertiserApiService
}

func NewAPIClient(cfg *conf.AdConfig) *APIClient {
	cli := new(APIClient)
	cli.Cfg = cfg
	cli.common.client = cli
	cli.CampaignApi = (*CampaignApiService)(&cli.common)
	cli.AdApi = (*AdApiService)(&cli.common)
	cli.CreativeApi = (*CreativeApiService)(&cli.common)
	cli.MaterialApi = (*MaterialApiService)(&cli.common)
	cli.ToolsApi = (*ToolsApiService)(&cli.common)
	cli.AdvertiserApi = (*AdvertiserApiService)(&cli.common)
	return cli
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.Cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.Cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(ctx context.Context, path string, method string, postBody []byte, headerParams map[string]string, queryParams url.Values) (req *http.Request, err error) {
	var urlParams *url.URL
	if urlParams, err = url.Parse(path); err != nil {
		return nil, err
	}
	// Adding Query Param
	query := urlParams.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	// Encode the parameters.
	urlParams.RawQuery = query.Encode()

	// Generate a new request
	if postBody != nil {
		req, err = http.NewRequestWithContext(ctx, method, urlParams.String(), bytes.NewReader(postBody))
	} else {
		req, err = http.NewRequestWithContext(ctx, method, urlParams.String(), nil)
	}
	if err != nil {
		return nil, err
	}
	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		req.Header = headers
	}

	// Override request host, if applicable
	if c.Cfg.Host != "" {
		req.Host = c.Cfg.Host
	}

	// Add the user agent to the request.
	req.Header.Set("User-Agent", c.Cfg.UserAgent)

	req.Header.Add("Access-Token", c.Cfg.GetAccessToken())
	for header, value := range c.Cfg.DefaultHeader {
		req.Header.Add(header, value)
	}

	return req, nil
}

// prepareRequest build the request
func (c *APIClient) prepareRequestAddFile(ctx context.Context, path string, body *bytes.Buffer, headerParams map[string]string) (req *http.Request, err error) {
	// Generate a new request
	if req, err = http.NewRequestWithContext(ctx, http.MethodPost, path, body); err != nil {
		return nil, err
	}
	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		req.Header = headers
	}
	// Override request host, if applicable
	if c.Cfg.Host != "" {
		req.Host = c.Cfg.Host
	}
	// Add the user agent to the request.
	req.Header.Set("User-Agent", c.Cfg.UserAgent)
	req.Header.Add("Access-Token", c.Cfg.GetAccessToken())
	for header, value := range c.Cfg.DefaultHeader {
		req.Header.Add(header, value)
	}
	return req, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/html") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}
