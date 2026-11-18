package yse_model

// 合同签约

type SmscSignReq struct { // 合同签约申请请求
	BusOpenType       string        `json:"busOpenType" v:"required"`         // 是否必填: Y;长度：100;到账方式,00-扫码工作日到账 01-扫码实时到账 10-刷卡工作日到账 11-刷卡实时到账 20-D1到账 21 境外卡到账 22-提现手续费 允许多选用“|”分隔 T1工作日到账必选
	ContractType      string        `json:"contractType" v:"required"`        // 是否必填: Y;长度：1;合同类型,1-纸质合同、2-电子合同，其他证明文件、基金会法人登记证书仅支持纸质合同，其他均支持电子合同、纸质合同。
	IsSendConMsg      string        `json:"isSendConMsg,omitempty" v:""`      // 是否必填: N;长度：1;是否自动发送签约通知,contractType=2时必填 0(短信+邮件) 1(短信) 2(邮件) 3(不通知)
	NotifyUrl         string        `json:"notifyUrl" v:"required"`           // 是否必填: Y;长度：100;异步通知url,银盛通过商户提供的地址下发通知，商户可根据异步通知业务参数bizContent中的notifyType确定通知类型，具体通知内容、场景以及接口要求详见如下:; 客户进件异步通知（notifyType = 1）; 合同签约异步通知（notifyType = 2）; 入网资料变更通知（notifyType = 3）; 费率变更异步通知（notifyType = 5）; 报备异步通知（notifyType = 6）
	CodeScanT1Fee     CodeScanT1Fee `json:"codeScanT1Fee,omitempty" v:""`     // 是否必填: N;长度：265;扫码工作日到账费率,busOpenType=00时必填
	CodeScanD0Fee     any           `json:"codeScanD0Fee,omitempty" v:""`     // 是否必填: N;长度：265;扫码实时到账垫资费,busOpenType=01时必填
	SwCardT1Fee       any           `json:"swCardT1Fee,omitempty" v:""`       // 是否必填: N;长度：265;刷卡工作日到账费率,busOpenType=10时必填
	SwCardD0Fee       any           `json:"swCardD0Fee,omitempty" v:""`       // 是否必填: N;长度：265;刷卡实时到账垫资费,busOpenType=11时必填
	D1Fee             any           `json:"d1Fee,omitempty" v:""`             // 是否必填: N;长度：265;天天到账垫资费,busOpenType=20时必填
	CustId            string        `json:"custId" v:"required"`              // 是否必填: Y;长度：20;客户号
	ThirdAuthId       string        `json:"thirdAuthId,omitempty" v:""`       // 是否必填: N;长度：32;业务方权限id,唯一（签约状态信息接口 可通过此字段查询）
	OutCardFeeList    any           `json:"outCardFeeList,omitempty" v:""`    // 是否必填: N;长度：265;境外卡费率
	CashwithdrawalFee any           `json:"CashwithdrawalFee,omitempty" v:""` // 是否必填: N;长度：265;提现手续费,从待结算户或者一般消费户提现时收取的手续费
}

type SmscSignRes struct { // 合同签约申请响应
	SignUrl string `json:"signUrl,omitempty" v:""` // 是否必填: N;长度：50;签约地址
	SignId  string `json:"signId,omitempty" v:""`  // 是否必填: N;长度：20;签约id
	AuthId  string `json:"authId,omitempty" v:""`  // 是否必填: N;长度：32;权限id
}

type CodeScanT1Fee struct { // 扫码工作日到账费率
	WxPayFee          WxPayFee          `json:"wxPayFee,omitempty" v:""`          // 是否必填: N;长度：265;微信扫码费率
	AliPayFee         AliPayFee         `json:"aliPayFee,omitempty" v:""`         // 是否必填: N;长度：265;支付宝扫码费率,不区分借贷记
	AliPayDebitFee    AliPayDebitFee    `json:"aliPayDebitFee,omitempty" v:""`    // 是否必填: N;长度：265;支付宝扫码借记卡费率,区分借贷记
	AliPayCreditFee   AliPayCreditFee   `json:"aliPayCreditFee,omitempty" v:""`   // 是否必填: N;长度：265;支付宝扫码贷记卡费率,区分借贷记
	Bank1DebitPayFee  Bank1DebitPayFee  `json:"bank1debitPayFee,omitempty" v:""`  // 是否必填: N;长度：265;银联一档借记卡扫码费率(>1000)
	Bank1CreditPayFee Bank1CreditPayFee `json:"bank1creditPayFee,omitempty" v:""` // 是否必填: N;长度：265;银联一档贷记卡扫码费率(>1000)
	Bank2DebitPayFee  Bank2DebitPayFee  `json:"bank2debitPayFee,omitempty" v:""`  // 是否必填: N;长度：265;银联二档借记卡扫码费率(<=1000)
	Bank2CreditPayFee Bank2CreditPayFee `json:"bank2creditPayFee,omitempty" v:""` // 是否必填: N;长度：265;银联二档贷记卡扫码费率(<=1000)
}

type WxPayFee struct { // 微信费率
	RateType   string `json:"rateType,omitempty" v:""`   // 是否必填: N;长度：1;收费方式,只支持0按百分比
	RateFee    string `json:"rateFee,omitempty" v:""`    // 是否必填: N;长度：16;费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom,omitempty" v:""` // 是否必填: N;长度：16;最低收费,单位为分
}

type AliPayFee struct { // 支付宝费率
	RateType   string `json:"rateType,omitempty" v:""`   // 是否必填: N;长度：1;收费方式,只支持0 按百分比
	RateFee    string `json:"rateFee,omitempty" v:""`    // 是否必填: N;长度：16;费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom,omitempty" v:""` // 是否必填: N;长度：16;最低收费,单位为:分
}

type AliPayDebitFee struct { // 支付宝扫码借记卡费率
	RateType   string `json:"rateType,omitempty" v:""`   // 是否必填: N;长度：265;收费方式,只支持0 按百分比
	RateFee    string `json:"rateFee,omitempty" v:""`    // 是否必填: N;长度：265;费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom,omitempty" v:""` // 是否必填: N;长度：265;最低收费,单位为:分
}

type AliPayCreditFee struct { // 支付宝扫码贷记卡费率
	RateType   string `json:"rateType,omitempty" v:""`   // 是否必填: N;长度：265;收费方式,只支持0 按百分比
	RateFee    string `json:"rateFee,omitempty" v:""`    // 是否必填: N;长度：265;费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom,omitempty" v:""` // 是否必填: N;长度：265;最低收费,单位为:分
}
type Bank1DebitPayFee struct { // 银联一档借记卡扫码费率
	RateType   string `json:"rateType,omitempty" v:""`   // 是否必填: N;长度：1;收费方式,只支持0 按百分比
	RateFee    string `json:"rateFee,omitempty" v:""`    // 是否必填: N;长度：16;费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom,omitempty" v:""` // 是否必填: N;长度：16;最低收费,单位为:分
	RateTop    string `json:"rateTop,omitempty" v:""`    // 是否必填: N;长度：16;最高收费,单位为:分
}

type Bank1CreditPayFee struct { // 银联一档贷记卡扫码费率
	RateType   string `json:"rateType,omitempty" v:""`   // 是否必填: N;长度：1;收费方式,只支持0按百分比
	RateFee    string `json:"rateFee,omitempty" v:""`    // 是否必填: N;长度：16;费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom,omitempty" v:""` // 是否必填: N;长度：16;最低收费,单位为分
}

type Bank2DebitPayFee struct { // 银联二档借记卡扫码费率
	RateType   string `json:"rateType,omitempty" v:""`   // 是否必填: N;长度：1;收费方式,只支持0按百分比
	RateFee    string `json:"rateFee,omitempty" v:""`    // 是否必填: N;长度：16;费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom,omitempty" v:""` // 是否必填: N;长度：16;最低收费,单位为：分
	RateTop    string `json:"rateTop,omitempty" v:""`    // 是否必填: N;长度：16;最高收费,单位为：分
}
type Bank2CreditPayFee struct { // 银联二档贷记卡扫码费率
	RateType   string `json:"rateType,omitempty" v:""`   // 是否必填: N;长度：1;收费方式,只支持0 按百分比
	RateFee    string `json:"rateFee,omitempty" v:""`    // 是否必填: N;长度：16;费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom,omitempty" v:""` // 是否必填: N;长度：16;最低收费,单位为：分
}
