package inter

import "net/http"

type Response interface {
	App() App
	SetApp(app App)
	GetContent() interface{}
	Content(content interface{})
	GetBody() string
	Body(body string) Response
	GetStatus() int
	Status(status int) Response
	GetHeader(key string) string
	GetHeaders() http.Header
	Header(key, value string) Response
	Headers(headers http.Header) Response
	Cookie(cookies ...http.Cookie) Response
	GetCookies() []http.Cookie
}
