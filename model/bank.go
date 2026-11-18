package yse_model

type BankReq struct { // bank查询请求
	CityCode     string `json:"cityCode,omitempty" v:""`     // 是否必填: N;长度：265;地区编码,自动查询下级区域 包括市、地区 1、支行行号不能与地区编码、行别同时为空；2、行别不为空时，地区编码必填
	BankType     string `json:"bankType,omitempty" v:""`     // 是否必填: N;长度：265;行别,支行行号不能与地区编码、行别同时为空
	OpenBankName string `json:"openBankName,omitempty" v:""` // 是否必填: N;长度：265;银行名称,模糊匹配
	BankCode     string `json:"bankCode,omitempty" v:""`     // 是否必填: N;长度：265;支行行号,行号不能与地区编码、行别同时为空
}

type BankRes struct {
	SystemCode   string     `json:"systemCode"`
	BankCodeList []BankCode `json:"bankCodeList"`
	ReqMsgID     string     `json:"reqMsgId"`
}

type BankCode struct {
	BankCode     string `json:"bankCode"`
	Superior     string `json:"superior"`
	CityCode     string `json:"cityCode"`
	BankType     string `json:"bankType"`
	ViewBankType string `json:"viewBankType"`
	DeadTime     string `json:"deadTime"`
	BankName     string `json:"bankName"`
	EffectTime   string `json:"effectTime"`
	OnePoint     string `json:"onePoint"`
	Remarks      string `json:"remarks"`
}
