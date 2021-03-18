package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"probe_material_plan/marketing/oceanengine/conf"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	ApplicationJson = "application/json"
	MultipartData   = "multipart/form-data"
)

var (
	jsonCheck = regexp.MustCompile("(?i:(?:application|text)/json)")
	xmlCheck  = regexp.MustCompile("(?i:(?:application|text)/xml)")
)

type service struct {
	client *APIClient
}

type APIClient struct {
	Cfg    *conf.AdConfig
	common service

	// api service
	AdCampaignApi *AdCampaignApiService
	AdAdApi       *AdAdApiService
	AdCreativeApi *AdCreativeApiService
	MaterialApi   *MaterialApiService
	ToolsApi      *ToolsApiService
	AdvertiserApi *AdvertiserApiService
}

func NewAPIClient(cfg *conf.AdConfig) *APIClient {
	cli := new(APIClient)
	cli.Cfg = cfg
	cli.common.client = cli
	cli.AdCampaignApi = (*AdCampaignApiService)(&cli.common)
	cli.AdAdApi = (*AdAdApiService)(&cli.common)
	cli.AdCreativeApi = (*AdCreativeApiService)(&cli.common)
	cli.MaterialApi = (*MaterialApiService)(&cli.common)
	cli.ToolsApi = (*ToolsApiService)(&cli.common)
	cli.AdvertiserApi = (*AdvertiserApiService)(&cli.common)
	return cli
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string
	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	case "multi":
		if jsonString, err := json.Marshal(obj); err == nil {
			return string(jsonString)
		}
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	if reflect.TypeOf(obj).Kind() == reflect.Struct {
		if jsonString, err := json.Marshal(obj); err == nil {
			return string(jsonString)
		}
	}
	return fmt.Sprintf("%v", obj)
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
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody []byte,
	headerParams map[string]string,
	queryParams url.Values) (req *http.Request, err error) {
	var (
		urlParams *url.URL
	)
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
func (c *APIClient) prepareRequestAddFile(ctx context.Context, path string, body *bytes.Buffer, headerParams map[string]string,
) (req *http.Request, err error) {
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

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}
	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}
	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyVal := strings.Split(part, "=")
			cc[strings.Trim(keyVal[0], " ")] = strings.Trim(keyVal[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type SwaggerError struct {
	body      []byte
	code      int
	message   string
	requestId string
}

func NewApiSwaggerError(code int, body []byte, msg string, requestId string) SwaggerError {
	return SwaggerError{code: code, body: body, message: msg, requestId: requestId}
}

// Error returns non-empty string if there was an error.
func (e SwaggerError) Error() string {
	b := bytes.Buffer{}
	b.WriteString("Code: ")
	b.WriteString(strconv.Itoa(e.code))
	if len(e.body) > 0 {
		b.WriteString(", response body: ")
		b.Write(e.body)
	}
	if len(e.message) > 0 {
		b.WriteString(", message: ")
		b.WriteString(e.message)
	}
	b.WriteString(", requestId: ")
	b.WriteString(e.requestId)
	return b.String()
}

// Body returns the raw bytes of the response
func (e SwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e SwaggerError) Message() interface{} {
	return e.message
}

func (e SwaggerError) RequestId() string {
	return e.requestId
}

func (e SwaggerError) Code() int {
	return e.code
}
