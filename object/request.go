package object

type LoginRequest struct {
	AcctID   string `json:"acctID"`
	Username string `json:"username"`
	LcId     int64  `json:"lcid"`
	Password string `json:"password"`
}
type ViewRequest struct {
	FormId string   `json:"formId"`
	Data   ViewData `json:"data"`
}

type ViewData struct {
	CreateOrgId int64  `json:"createOrgId,omitempty"`
	Number      string `json:"Number,omitempty"`
	Id          string `json:"Id,omitempty"`
	IsSortBySeq bool   `json:"IsSortBySeq,default=false,omitempty"`
}

type AllocateRequest struct {
	FormId string       `json:"formId"`
	Data   AllocateData `json:"data"`
}

type AllocateData struct {
	PkIds   int64  `json:"PkIds"`
	TOrgIds string `json:"TOrgIds"`
}

type SaveRequest struct {
	FormId string   `json:"formId"`
	Data   SaveData `json:"data"`
}

type SaveData struct {
	Model ModelObject `json:"model"`
}

type ModelObject struct {
	FSupplierId  int           `json:"FSupplierId"`  //实体主建
	FCreateOrgId FNumberObject `json:"FCreateOrgId"` //创建组织
	FNumber      string        `json:"FNumber"`      //编码
	FUseOrgId    FNumberObject `json:"FUseOrgId"`    //使用组织
	FName        string        `json:"FName"`        //名称
}

type FNumberObject struct {
	FNumber string `json:"FNumber"` //编码
}

type SubmitRequest struct {
	FormId string     `json:"formId"`
	Data   SubmitData `json:"data"`
}
type SubmitData struct {
	Number   []string `json:"Numbers,omitempty"`
	UseOrgId int64    `json:"UseOrgId,omitempty"`
	Ids      string   `json:"Ids,omitempty"`
}

type AuditRequest struct {
	FormId string    `json:"formId"`
	Data   AuditData `json:"data"`
}
type AuditData struct {
	Number   []string `json:"Numbers,omitempty"`
	UseOrgId int64    `json:"UseOrgId,omitempty"`
	Ids      string   `json:"Ids,omitempty"`
}

type ExecuteBillQueryRequest struct {
	Data ExecuteBillQueryData `json:"data"`
}

type ExecuteBillQueryData struct {
	FormId       string   `json:"FormId"`
	FieldKeys    string   `json:"FieldKeys"`
	FilterString []string `json:"FilterString,omitempty"`
	OrderString  int64    `json:"OrderString,omitempty"`
	TopRowCount  int64    `json:"TopRowCount,omitempty"`
	StartRow     int64    `json:"StartRow,omitempty"`
	Limit        int64    `json:"Limit,omitempty"`
	SubSystemId  int64    `json:"SubSystemId,omitempty"`
}
