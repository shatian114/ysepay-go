package yse_model

type BalanceQueryReq struct { // bank查询请求
	MercId string `json:"mercId,omitempty" v:""` // 商户号
}

type BalanceQueryRes struct { // 余额查询响应
	ReqMsgId   string       `json:"reqMsgId,omitempty" v:""` // 是否必填: N;长度：265;请求流水号
	Data       BusinessData `json:"data,omitempty" v:""`     // 是否必填: N;长度：265;响应结果
	SystemCode string       `json:"systemCode"`
}

type BusinessData struct { // 余额查询结果业务数据
	AccountTotalAmount         string          `json:"accountTotalAmount" v:"required"`         // 是否必填: Y;长度：10;账户总额,单位：分。
	AccountSettledUnpaidAmount string          `json:"accountSettledUnpaidAmount" v:"required"` // 是否必填: Y;长度：10;所有账户累计已结算未付款的金额,单位：分。
	AccountDetailList          []AccountDetail `json:"accountDetailList" v:"required"`          // 是否必填: Y;长度：265;账户明细信息列表
}

type AccountDetail struct { // 余额查询结果里的账号详情
	AccountId              string `json:"accountId" v:"required"`                // 是否必填: Y;长度：10;账户号
	AccountType            string `json:"accountType" v:"required"`              // 是否必填: Y;长度：10;账户类型,类型如下：; 一般消费类账户; 待结算账户; 代收账户; 代付账户; 订单账户; 海盟授信账户; 代付授信账户
	AccountAmount          string `json:"accountAmount" v:"required"`            // 是否必填: Y;长度：10;账户余额,单位：分。
	AccountWithdrawBalance string `json:"accountWithdrawBalance,omitempty" v:""` // 是否必填: N;长度：10;预可提现余额,单位：分，只针对一般消费类账户与待结算账户。
}
