package inter

import (
	"github.com/lanvard/support"
	"net/http"
)

const RequestBodyDecoder = "request_body_decoder"

type Request interface {
	HeaderHolder
	App() App
	SetApp(app App)
	Make(abstract interface{}) interface{}
	Content() string
	SetContent(content string) Request
	Source() http.Request
	Header(key string) string
	Headers() http.Header
	Cookie(key string) string
	CookieE(key string) (string, error)
	Method() string
	Path() string
	Url() string
	FullUrl() string
	Body(key ...string) support.Value
	BodyOr(keys string, defaultValue interface{}) support.Value
	Parameter(key string) support.Value
	ParameterOr(key string, defaultValue interface{}) support.Value
	SetUrlValues(vars map[string]string) Request
	Query(key string) support.Value
	QueryOr(key string, defaultValue interface{}) support.Value
	Route() Route
}
