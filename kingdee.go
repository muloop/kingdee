package kingdee

import (
	"context"
	"github.com/imroc/req/v3"
	"github.com/muloop/kingdee/object"
)

type KingDee struct {
	ctx     context.Context
	debug   bool
	baseUrl string
	client  *req.Client
	cookie  *Cookie
	formId  string
	json    *SafeJSONPool
	err     error
}

func New(ctx context.Context, c Config) (*KingDee, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	cookie := NewCookie(c)
	kingDee := &KingDee{
		ctx:     ctx,
		debug:   c.Debug,
		baseUrl: c.LoginConfig.Host,
		cookie:  cookie,
		json:    NewSafeJSONPool(),
	}
	return kingDee, nil
}

func (k *KingDee) View(formId string, data object.ViewData, resp any) *req.Response {
	var params object.ViewRequest
	params.FormId = formId
	params.Data = data
	var body string
	body, k.err = k.json.Marshal(params)
	return k.call(VIEW_API, body, resp)
}
func (k *KingDee) Draft() {

}
func (k *KingDee) Save(formId string, data object.SaveData, resp any) *req.Response {
	var params object.SaveRequest
	params.FormId = formId
	params.Data = data
	var body string
	body, k.err = k.json.Marshal(params)
	return k.call(SAVE_API, body, resp)
}
func (k *KingDee) Submit(formId string, data object.SubmitData, resp any) *req.Response {
	var params object.SubmitRequest
	params.FormId = formId
	params.Data = data
	var body string
	body, k.err = k.json.Marshal(params)
	return k.call(SUBMIT_API, body, resp)
}
func (k *KingDee) Audit(formId string, data object.SubmitData, resp any) *req.Response {
	var params object.SubmitRequest
	params.FormId = formId
	params.Data = data
	var body string
	body, k.err = k.json.Marshal(params)
	return k.call(AUDIT_API, body, resp)
}

// Allocate 分配
func (k *KingDee) Allocate(formId string, data object.AllocateData, resp any) *req.Response {
	var params object.AllocateRequest
	params.FormId = formId
	params.Data = data
	var body string
	body, k.err = k.json.Marshal(params)
	return k.call(ALLOCATE_API, body, resp)
}
func (k *KingDee) CancelAllocate() {

}

func (k *KingDee) ExecuteBillQuery(data object.ExecuteBillQueryData, resp any) *req.Response {
	var params object.ExecuteBillQueryRequest
	params.Data = data
	var body string
	body, k.err = k.json.Marshal(params)
	return k.call(EXECUTEBILLQUERY_API, body, resp)
}

func (k *KingDee) call(url, body string, resp any) *req.Response {

	return NewClient(k.cookie, k.baseUrl, url, body, k.debug).SetSuccessResult(resp).Do(k.ctx)
}
