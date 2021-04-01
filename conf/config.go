package conf

import (
	"net/http"
	"sync"
)

func (c *AdConfig) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
func (c *AdConfig) GetAccessToken() string {
	c.accessTokenMu.RLock()
	defer c.accessTokenMu.RUnlock()
	return c.accessToken
}

func (c *AdConfig) SetAccessToken(token string) {
	c.accessTokenMu.Lock()
	c.accessToken = token
	c.accessTokenMu.Unlock()
}

func (c *AdConfig) SetUserAgent(userAgent string) {
	c.UserAgent = userAgent
}

type AdConfig struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	ApiVersion    string            `json:"apiVersion,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	accessToken   string
	accessTokenMu *sync.RWMutex
	HTTPClient    *http.Client
	IsDebug       bool
	GlobalConfig  GlobalConfig
}

func NewAdConfig(debug bool) *AdConfig {
	return &AdConfig{
		IsDebug:       debug,
		DefaultHeader: make(map[string]string),
		UserAgent:     "Toutiao Ads Marketing API SDK",
		HTTPClient:    &http.Client{},
		accessTokenMu: &sync.RWMutex{},
	}
}

type GlobalConfig struct {
	ServiceName ServiceName
	HttpOption  HttpOption
}

type ServiceName struct {
	Name   string
	Schema string
}

type HttpOption struct {
	Header http.Header
}
