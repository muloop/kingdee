package kingdee

import (
	"github.com/imroc/req/v3"
)

func NewClient(cookie *Cookie, baseUrl, url, body string) *req.Request {
	cookies := cookie.Cookie()
	cli := req.C().
		DevMode().
		SetCommonContentType("application/json").
		SetCommonCookies(cookies...).
		SetBaseURL(baseUrl)
	cli.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
		return nil
	})
	return cli.Post(url).SetBody(body)
}
