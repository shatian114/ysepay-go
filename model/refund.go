package yse_model

type RefundReq struct { // 银盛退款请求
	OrigOrderId   string `json:"origOrderId,omitempty" v:""` // 是否必填: N;长度：32;商户原交易订单号,原订单号、原交易流水号不能同时为空。如果同时存在优先取原订单号。示例值：201805256843192280647118。
	OrigTradeSn   string `json:"origTradeSn,omitempty" v:""` // 是否必填: N;长度：30;原交易流水号,银盛原交易流水号，原订单号、原交易流水号不能同时为空。如果同时存在优先取原订单号。 示例值：311160414497667096。
	ShopDate      string `json:"shopDate" v:"required"`      // 是否必填: Y;长度：8;商户日期,该参数做交易与查询时需要一致。
	RefundAmount  string `json:"refundAmount" v:"required"`  // 是否必填: Y;长度：10;申请退款金额,需要退款的金额，该金额不能大于订单金额,单位为：分。
	RefundReason  string `json:"refundReason" v:"required"`  // 是否必填: Y;长度：50;退款原因,退款的原因说明。
	RefundOrderId string `json:"refundOrderId" v:"required"` // 是否必填: Y;长度：32;商户退款订单号,商户系统生成的退款订单号，示例值：RD2012061713107，要求如下：; 1、标识一次退款请求，同一笔交易多次退款需要保证唯一。; 2、该参数支持汉字，最大长度为16个。; 3、用同一订单号同一退款流水号继续退款时，则会返回第一次退款的结果不会继续退款 。
	NotifyUrl     string `json:"notifyUrl,omitempty" v:""`   // 是否必填: N;长度：265;通知回调地址
}
type RefundRes struct { // 银盛退款响应
	Data       RefundResData `json:"data"`
	SystemCode string        `json:"systemCode"`
	ReqMsgID   string        `json:"reqMsgId"`
}

type RefundResData struct {
	RefundOrderID string `json:"refundOrderId"`
	RefundTradeSn string `json:"refundTradeSn"`
	AccountDate   string `json:"accountDate"`
	OrigOrderID   string `json:"origOrderId"`
	OrigTradeSn   string `json:"origTradeSn"`
	RefundAmount  string `json:"refundAmount"`
}

type RefundQueryReq struct { // 银盛退款查询请求
	OrigOrderId     string `json:"origOrderId,omitempty" v:""`     // 是否必填: N;长度：32;商户原交易订单号,原订单号、原交易流水号不能同时为空。如果同时存在优先取原订单号。示例值：201805256843192280647118。
	OrigTradeSn     string `json:"origTradeSn,omitempty" v:""`     // 是否必填: N;长度：30;原交易流水号,银盛原交易流水号，原订单号、原交易流水号不能同时为空。如果同时存在优先取原订单号。 示例值：311160414497667096。
	RefundOrderId   string `json:"refundOrderId" v:"required"`     // 是否必填: Y;长度：32;商户退款订单号,请求退款接口时，传入的退款请求订单号， 示例值：RD2012061713107。
	RefundQueryType string `json:"refundQueryType,omitempty" v:""` // 是否必填: N;长度：2;退款查询类型,退款查询类型 1-普通退款查询 2-分账退款查询； 分账退款查询时，原交易流水号origTradeSn必输
}

type RefundQueryRes struct {
	Data       RefundQueryResData `json:"data"`
	SystemCode string             `json:"systemCode"`
	ReqMsgID   string             `json:"reqMsgId"`
}

type RefundQueryResData struct {
	PayeeFeeFlag               string `json:"payeeFeeFlag"`
	RefundState                string `json:"refundState"`
	AccountDate                string `json:"accountDate"`
	RefundChannelFundsDynamics string `json:"refundChannelFundsDynamics"`
	OrigOrderID                string `json:"origOrderId"`
	PayeeFee                   string `json:"payeeFee"`
	RefundReason               string `json:"refundReason"`
	PayerFeeFlag               string `json:"payerFeeFlag"`
	SrcFeeFlag                 string `json:"srcFeeFlag"`
	PayerFee                   string `json:"payerFee"`
	RefundOrderID              string `json:"refundOrderId"`
	TotalAmount                string `json:"totalAmount"`
	SrcFee                     string `json:"srcFee"`
	MarktingRefundDetail       string `json:"marktingRefundDetail"`
	OrigTradeSn                string `json:"origTradeSn"`
	FundsState                 string `json:"fundsState"`
	RefundAmount               string `json:"refundAmount"`
}

type RefundSharingMerc struct { // 银盛分账退款请求里的分账商户
	RefundMerId  string `json:"refundMerId" v:"required"`  // 是否必填: Y;长度：265;分账商户
	RefundAmount string `json:"refundAmount" v:"required"` // 是否必填: Y;长度：265;分账金额（单位：分）
}

type RefundNotifyObj struct { // 退款回调数据
	SignType                   string              `json:"sign_type"`
	AlipayRefundCouponInfo     string              `json:"alipay_refund_coupon_info"`
	WxRefundCouponInfo         WxRefundCouponInfo  `json:"wxRefundCouponInfo"`
	WxpayRefundCouponInfo      string              `json:"wxpay_refund_coupon_info"`
	TradeSn                    string              `json:"tradeSn"`
	OrderID                    string              `json:"orderId"`
	RefundNotifyObjAccountDate string              `json:"account_date"`
	Sign                       string              `json:"sign"`
	State                      string              `json:"state"`
	ChannelRecvTime            string              `json:"channelRecvTime"`
	AccountDate                string              `json:"accountDate"`
	TotalAmount                string              `json:"totalAmount"`
	AliRefundCouponInfo        AliRefundCouponInfo `json:"aliRefundCouponInfo"`
	Version                    string              `json:"version"`
	NotifyType                 string              `json:"notifyType"`
	RefundNotifyObjTotalAmount string              `json:"total_amount"`
	TradeStatus                string              `json:"tradeStatus"`
	SettlementAmount           string              `json:"settlementAmount"`
	BgURL                      string              `json:"bgUrl"`
	MerCharset                 string              `json:"merCharset"`
	UserCode                   string              `json:"userCode"`
	RefundNotifyObjNotifyTime  string              `json:"notify_time"`
	NotifyTime                 string              `json:"notifyTime"`
	RefundNotifyObjTradeNo     string              `json:"trade_no"`
	Amount                     string              `json:"amount"`
	RefundNotifyObjTradeStatus string              `json:"trade_status"`
	TradeSource                string              `json:"tradeSource"`
	OutTradeNo                 string              `json:"out_trade_no"`
	RefundNotifyObjNotifyType  string              `json:"notify_type"`
	TradeNo                    string              `json:"tradeNo"`
	SrcMercID                  string              `json:"srcMercId"`
}

type AliRefundCouponInfo struct {
	PresentRefundBuyerAmount     float64        `json:"presentRefundBuyerAmount"`
	SendBackFee                  float64        `json:"sendBackFee"`
	RefundAmount                 float64        `json:"refundAmount"`
	PresentRefundDiscountAmount  float64        `json:"presentRefundDiscountAmount"`
	GmtRefundPay                 float64        `json:"gmtRefundPay"`
	RefundFee                    float64        `json:"refundFee"`
	PresentRefundMdiscountAmount float64        `json:"presentRefundMdiscountAmount"`
	RefundDetail                 []RefundDetail `json:"refundDetail"`
}

type RefundDetail struct {
	Amount         float64 `json:"amount"`
	FundChannelStr string  `json:"fundChannelStr"`
	RealAmount     float64 `json:"realAmount"`
}

type WxRefundCouponInfo struct {
	CashRefundFee       float64 `json:"cashRefundFee"`
	FeeType             string  `json:"feeType"`
	SettlementRefundFee float64 `json:"settlementRefundFee"`
	EXT                 string  `json:"ext"`
	CouponRefundFee     float64 `json:"couponRefundFee"`
	RefundFee           float64 `json:"refundFee"`
}
