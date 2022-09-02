package xhttp

const (
	// ConContent-Type
	ConContentType           string = "ConContent-Type"
	ConContentTypeJson       string = "application/json; charset=utf-8"
	ConContentTypeJavaScript string = "application/x-javascript;charset=UTF-8"

	// Connection
	Connection          string = "Connection"
	ConnectionKeepAlive string = "keep-alive"

	// Accept-Language
	AcceptLanguage     string = "Accept-Language"
	AcceptLanguageZhCN string = "zh-CN,zh;q=0.9"

	// Accept-Encoding
	AcceptEncoding        string = "Accept-Language"
	AcceptEncodingGzip    string = "gzip"
	AcceptEncodingDeflate string = "deflate"
	AcceptEncodingBr      string = "br"
	AcceptEncodingDefault string = "gzip, deflate, br"

	// Strict-Transport-Security
	StrictTransportSecurity                 string = "Strict-Transport-Security"
	StrictTransportSecurityMax              string = "max-age=63072000"
	StrictTransportSecurityIncludeSubDomain string = "includeSubDomains"

	// User-Agent
	UserAgent                          string = "User-Agent"
	UserAgentWindowsChrome93_0_4577_63 string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36"
	UserAgentWindowsFirefox92_0        string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0"
)
