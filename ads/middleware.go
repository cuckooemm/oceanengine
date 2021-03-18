package ads

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

type Middleware interface {
	Handle(req *http.Request, next func(req *http.Request) (rsp *http.Response, err error)) (rsp *http.Response, err error)
}

// LogMiddleware ...
type LogMiddleware struct {
	cli *AdClient
	fd  io.Writer
}

func NewLogMiddleware(c *AdClient, f io.Writer) *LogMiddleware {
	return &LogMiddleware{
		cli: c,
		fd:  f,
	}
}

// Handle ...
func (l *LogMiddleware) Handle(req *http.Request, next func(req *http.Request) (rsp *http.Response, err error)) (rsp *http.Response, err error) {
	cli := l.cli
	buf := bytes.Buffer{}
	buf.Reset()
	buf.Grow(512)
	if cli.IsDebug() {
		if l.fd == nil {
			panic("not set output file")
		}
		// 过滤上传类日志
		if !strings.Contains(req.Header.Get("Content-Type"), "multipart/form-data") || strings.Contains(req.RequestURI, "/file/video/ad") {
			if request, err := httputil.DumpRequestOut(req, true); err == nil {
				buf.Write(request)
			}
		}
	}
	rsp, err = next(req)
	if cli.IsDebug() {
		if rsp != nil {
			if response, err := httputil.DumpResponse(rsp, true); err == nil {
				buf.Write(response)
				buf.WriteByte('\n')
			}
		}
		if err != nil {
			errStr, _ := json.Marshal(err)
			buf.Write(errStr)
		}
		_, _ = l.fd.Write(buf.Bytes())
	}
	return rsp, err
}
