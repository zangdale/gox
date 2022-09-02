package xhttp

import (
	"net/http"
	"net/url"
	"strconv"
)

var (
	DefaultPageKey        = "page"
	DefaultPerPageKey     = "per_page"
	DefaultPerPageValue   = 20
	DefaultBoolTrueValue  = "true"
	DefaultBoolFalseValue = "false"
)

type Request struct {
	*http.Request
	UrlQuery    URLValues
	UrlForm     URLValues
	UrlPostForm URLValues
}

type URLValues url.Values

// NewRequest 创建一个封装后的解析内容
func NewRequest(req *http.Request) *Request {
	if req == nil {
		return nil
	}
	res := &Request{Request: req, UrlForm: URLValues(req.Form), UrlPostForm: URLValues(req.PostForm)}
	if req.URL != nil {
		res.UrlQuery = URLValues(req.URL.Query())
	}

	return res
}

// GetPageMsgs 获取页码和每页数量
func (req *Request) GetPageMsgs(pageKey, perPageKey string) (pageV int, perPageV int) {
	if pageKey == "" {
		pageKey = DefaultPageKey
	}
	pageV = req.UrlQuery.GetQueryInt(pageKey)
	if pageV < 0 {
		pageV = 1
	}

	if perPageKey == "" {
		perPageKey = DefaultPerPageKey
	}
	perPageV = req.UrlQuery.GetQueryInt(perPageKey)
	if perPageV <= 0 {
		perPageV = DefaultPerPageValue
	}

	return
}

// GetQuery 从 query 中获取数据
func (req *Request) GetQuery(key string) (string, bool) {
	return req.UrlQuery.get(key)
}

// GetForm 从 form 中获取数据
func (req *Request) GetForm(key string) (string, bool) {
	return req.UrlForm.get(key)
}

// GePostForm 从 postform 中获取数据
func (req *Request) GePostForm(key string) (string, bool) {
	return req.UrlPostForm.get(key)
}

func (vls URLValues) get(key string) (string, bool) {

	if vls == nil || key == "" {
		return "", false
	}

	vs, ok := vls[key]

	if !ok || len(vs) == 0 {
		return "", false
	}
	return vs[0], true
}

func (vls URLValues) GetQueryString_(key string) *string {
	if v, ok := vls.get(key); ok {
		return &v
	}

	return nil
}

func (vls URLValues) GetQueryString(key string) string {
	if v, ok := vls.get(key); ok {
		return v
	}
	return ""
}

func (vls URLValues) GetQueryInt_(key string) *int {
	if v, ok := vls.get(key); ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return &i
		}
	}

	return nil
}

func (vls URLValues) GetQueryInt(key string) int {
	if v, ok := vls.get(key); ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	return 0
}

func (vls URLValues) GetQueryBool_(key string) *bool {
	if v, ok := vls.get(key); ok {
		if v == DefaultBoolFalseValue {
			f := false
			return &f
		}
		if v == DefaultBoolTrueValue {
			t := true
			return &t
		}
	}
	return nil
}

func (vls URLValues) GetQueryBool(key string) bool {
	if v, ok := vls.get(key); ok {
		return v == DefaultBoolTrueValue
	}

	return false
}
