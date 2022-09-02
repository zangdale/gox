package xhttp

import "net/http"

// ClientNoRedirect 不进行 30x 跳转的客户端
var ClientNoRedirect = &http.Client{
	CheckRedirect: CheckRedirectNoRedirectFunc,
}

// CheckRedirectNoRedirectFunc client 禁止跳转
var CheckRedirectNoRedirectFunc = func(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
