package yse_model

type WithdrawReq struct { // 提现请求
	NotifyUrl     string `json:"notifyUrl" v:"required"`       // 是否必填: Y;长度：190;交易通知地址,必填；
	OrderId       string `json:"orderId" v:"required"`         // 是否必填: Y;长度：32;商户订单号,由大小写英文字母、数字、下划线及横杠组成
	ShopDate      string `json:"shopDate" v:"required"`        // 是否必填: Y;长度：8;交易日期,交易日期，格式：yyyyMMdd
	MercId        string `json:"mercId,omitempty" v:""`        // 是否必填: N;长度：32;商户号,1. 可空，商户号为空时，使用certId,作为商户号；; 2.商户号和发起方商户号不一致，会检查机构信息,或通过发起方的代理密码，检查代理关系
	Currency      string `json:"currency" v:"required"`        // 是否必填: Y;长度：265;币种,CNY：人民币; 大小写敏感
	Amount        string `json:"amount" v:"required"`          // 是否必填: Y;长度：10;订单金额,单位:分
	Remark        string `json:"remark" v:"required"`          // 是否必填: Y;长度：250;订单备注
	BankAccountNo string `json:"bankAccountNo,omitempty" v:""` // 是否必填: N;长度：32;到帐卡账号,cardType 为 提现卡（1），此参数必填；; cardType 为 结算卡（2），此参数可空
	CardType      string `json:"cardType,omitempty" v:""`      // 是否必填: N;长度：265;到账卡类型,1：提现卡; 2：结算卡; 为空则默认1-提现卡
	ProxyPassword string `json:"proxyPassword,omitempty" v:""` // 是否必填: N;长度：265;代理密码,商户号和发起方商户号不一致，通过发起方的代理密码，检查代理关系
	BusiCode      string `json:"busiCode,omitempty" v:""`      // 是否必填: N;长度：265;业务代码,可空，默认：00510003（支付账户提现）
}

type WithdrawRes struct { // 提现响应
	TradeStatus    string `json:"tradeStatus,omitempty" v:""`    // 是否必填: N;长度：265;交易状态,00: 支付成功; 99: 支付失败; 10: 处理中; 50: 超时; 其他: 均作为处理中处理，请稍后发起查询获取交易结果
	TradeStatusMsg string `json:"tradeStatusMsg,omitempty" v:""` // 是否必填: N;长度：265;状态描述
	OrderId        string `json:"orderId,omitempty" v:""`        // 是否必填: N;长度：265;商户订单号,同请求报文“商户订单号”
	TradeSn        string `json:"tradeSn,omitempty" v:""`        // 是否必填: N;长度：265;交易流水号,银盛内部订单号
	AccountDate    string `json:"accountDate,omitempty" v:""`    // 是否必填: N;长度：265;会计日期,格式:yyyyMMdd
	Amount         string `json:"amount,omitempty" v:""`         // 是否必填: N;长度：265;订单金额,单位:分
	Fee            string `json:"fee,omitempty" v:""`            // 是否必填: N;长度：265;手续费,单位:分
	PayeeFee       string `json:"payeeFee,omitempty" v:""`       // 是否必填: N;长度：265;收款方手续费,单位:分
	PartnerFee     string `json:"partnerFee,omitempty" v:""`     // 是否必填: N;长度：265;发起方手续费,单位:分
	PayerFee       string `json:"payerFee,omitempty" v:""`       // 是否必填: N;长度：265;付款方手续费,单位:分
}

type WithdrawQueryReq struct { // 提现查询请求
	OrderId  string `json:"orderId" v:"required"`  // 是否必填: Y;长度：32;商户订单号,商户系统生成的订单号。
	ShopDate string `json:"shopDate" v:"required"` // 是否必填: Y;长度：8;交易日期,商户系统的交易发生日期，格式yyyyMMdd， 示例值：20180525。
}

type WithdrawQueryRes struct { // 提现查询响应
	ReqMsgId     string                               `json:"reqMsgId,omitempty" v:""` // 是否必填: N;长度：265;请求流水号
	BusinessData `json:"businessData,omitempty" v:""` // 是否必填: N;长度：265;响应结果
}

type WithdrawQueryData struct { // 提现查询响应数据
	OrderId                string `json:"orderId" v:"required"`                // 是否必填: Y;长度：32;订单号,商户系统生成的订单号。
	TradeStatus            string `json:"tradeStatus" v:"required"`            // 是否必填: Y;长度：2;交易目前所处的状态,状态如下：; 00-支付成功/交易成功; 10-处理中/受理成功，请耐心等待异步通知结果; 11-已向渠道发起; 12-代付付款方支付; 11-处理中，已向渠道发起; 45-渠道参数转换完成; 46-交易待授权; 47-已获取授权码; 50-内部处理超时; 51-调用支付网关超时; 52-渠道超时; 53-订单超时处理中; 80-部分退款; 81-全部退款成功; 94-审核不通过; 99-支付失败/交易失败
	TradeStatusDescription string `json:"tradeStatusDescription" v:"required"` // 是否必填: Y;长度：256;状态描述
	TotalAmount            string `json:"totalAmount" v:"required"`            // 是否必填: Y;长度：10;订单总金额,单位：分。
	AccountDate            string `json:"accountDate,omitempty" v:""`          // 是否必填: N;长度：8;会计日期,格式为yyyyMMdd，示例值：20160413。
	TradeSn                string `json:"tradeSn,omitempty" v:""`              // 是否必填: N;长度：30;银盛支付交易流水号
	Fee                    string `json:"fee,omitempty" v:""`                  // 是否必填: N;长度：10;总手续费,单位：分。
	PartnerFee             string `json:"partnerFee" v:"required"`             // 是否必填: Y;长度：10;合作方手续费,单位：分。
	PayeeFee               string `json:"payeeFee" v:"required"`               // 是否必填: Y;长度：10;收款方手续费,单位：分。
	PayerFee               string `json:"payerFee" v:"required"`               // 是否必填: Y;长度：10;付款方手续费,单位：分。
	WithheldAmount         string `json:"withheldAmount,omitempty" v:""`       // 是否必填: N;长度：10;增值费,单位：分。
	ChargeRemark           string `json:"chargeRemark,omitempty" v:""`         // 是否必填: N;长度：265;增值费描述
}
