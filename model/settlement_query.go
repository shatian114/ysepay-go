package yse_model

type SettlementQueryReq struct { // 查询绑定的银行卡请求
	Custid   string `json:"custid,omitempty" v:""`   // 是否必填: N;长度：265;客户编号,商户号和客户编号任选其一必传，同时传入取商户号
	UserCode string `json:"userCode,omitempty" v:""` // 是否必填: N;长度：265;商户号,商户号和客户编号任选其一必传，同时传入取商户号
}

type SettlementDetail struct { // 查询绑定的银行卡详情
	BankAccountNo    string `json:"bankAccountNo" v:"required"`    // 是否必填: Y;长度：265;卡号
	BankAccountName  string `json:"bankAccountName" v:"required"`  // 是否必填: Y;长度：265;户名
	BankAccountType  string `json:"bankAccountType" v:"required"`  // 是否必填: Y;长度：265;银行账户类型,11：个人借记卡12：个人贷记卡13：个人存折21：对公借记卡22：对公贷记卡23：对公存折 24：单位结算卡
	BankCardTypeList string `json:"bankCardTypeList" v:"required"` // 是否必填: Y;长度：265;卡类型,1：结算卡；2：提现卡；3：快捷绑卡；
	BindTime         string `json:"bindTime" v:"required"`         // 是否必填: Y;长度：265;绑定时间
	UpdateTime       string `json:"updateTime,omitempty" v:""`     // 是否必填: N;长度：265;更新时间
}

type SettlementQueryRes struct { // 查询绑定的银行卡响应
	State        string             `json:"state,omitempty" v:""`        // 是否必填: N;长度：265;交易状态
	BusinessData []SettlementDetail `json:"businessData,omitempty" v:""` // 是否必填: N;长度：265;业务响应数据
}
