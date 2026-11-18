package yse_model

type PayNotifyObj struct { // 微信h5 支付回调里的bizContent
	CardType             string `json:"cardType"`
	PreferentialAmount   string `json:"preferentialAmount"`
	ChannelRecvSn        string `json:"channelRecvSn"`
	AliPlatformDisAmount string `json:"aliPlatformDisAmount"`
	MercID               string `json:"mercId"`
	TradeSource          string `json:"tradeSource"`
	Note                 string `json:"note"`
	TotalDiscount        string `json:"totalDiscount"`
	PayTime              string `json:"payTime"`
	PayerFee             string `json:"payerFee"`
	PayerBankaccountType string `json:"payerBankaccountType"`
	PayeeFee             string `json:"payeeFee"`
	Fee                  string `json:"fee"`
	TotalAmount          string `json:"totalAmount"`
	NotifyType           string `json:"notifyType"`
	BusiCode             string `json:"busiCode"`
	CurrencyCode         string `json:"currencyCode"`
	ChannelSendSn        string `json:"channelSendSn"`
	PartnerFee           string `json:"partnerFee"`
	NotifyTime           string `json:"notifyTime"`
	WxpayCouponInfo      string `json:"wxpayCouponInfo"`
	PayAmt               string `json:"payAmt"`
	OrderID              string `json:"orderId"`
	IsDiscount           string `json:"isDiscount"`
	TradeNo              string `json:"tradeNo"`
	TradeSn              string `json:"tradeSn"`
	SrcFee               string `json:"srcFee"`
	PayerBankType        string `json:"payerBankType"`
	SrcUsercode          string `json:"srcUsercode"`
	SettlementAmt        string `json:"settlementAmt"`
	SrcMercID            string `json:"srcMercId"`
	TotalDiscountAmt     string `json:"totalDiscountAmt"`
	Openid               string `json:"openid"`
	BuyerUserId          string `json:"buyerUserId"`
	TradeStatus          string `json:"tradeStatus"`
	Amount               string `json:"amount"`
	AccountDate          string `json:"accountDate"`
	AliMerchantAmount    string `json:"aliMerchantAmount"`
	OpenID               string `json:"openId"`
	PreferentialFee      string `json:"preferentialFee"`
	SumFee               string `json:"sumFee"`
	MerDiscountFee       string `json:"merDiscountFee"`
	DctGoodsInfo         string `json:"dctGoodsInfo"`
	PayMode              string `json:"payMode"`
}

type ReportNotifyReq struct { // 三方报备回调
	NotifyType string `json:"notifyType"`
	Report     Report `json:"report"`
	Status     string `json:"status"`
}

type Report struct {
	AuthID        string      `json:"authId"`
	MercID        string      `json:"mercId"`
	ThridMercList []ThridMerc `json:"thridMercList"`
}

type ThridMerc struct {
	ApprSts       string `json:"apprSts"`
	MercID        string `json:"mercId"`
	Remark        string `json:"remark"`
	ReportChannel string `json:"reportChannel"`
	ThridMercID   string `json:"thridMercId"`
}

type SignNotifyReq struct { // 进件签约回调
	Auth       Auth   `json:"auth"`
	NotifyType string `json:"notifyType"`
	Status     string `json:"status"`
}

type Auth struct {
	CustID string `json:"custId"`
	MercID string `json:"mercId"`
	Note   string `json:"note"`
	AuthID string `json:"authId"`
}

type OpenSharingNotify struct { // 申请分账开通回调
	OrderNo string `json:"orderNo"`           // 申请单号
	Status  string `json:"status"`            // 申请单状态 枚举值： 0:已完成 2:审核中 3:待签约 9:审核拒绝
	Remark  string `json:"remark,omitempty"`  // 备注
	SignUrl string `json:"signUrl,omitempty"` // 签约URL（状态为待签约时返回签约链接）
	SignId  string `json:"signId,omitempty"`  // 签约ID（后签约时使用）（暂未实现）
}

type SharingNotify struct {
	Status        string `json:"status"`
	Note          string `json:"note"`
	Amount        string `json:"amount"`
	PayeeUsercode string `json:"payeeUsercode"`
	FeePayer      string `json:"feePayer"`
	Fee           string `json:"fee"`
	SrcUsercode   string `json:"srcUsercode"`
	TradeSource   string `json:"tradeSource"`
	SrcMercID     string `json:"srcMercId"`
	DivList       string `json:"divList"`
	OrderID       string `json:"orderId"`
	DivisionMode  string `json:"divisionMode"`
}

type SharingRefundNotify struct { // 分账退款回调
	SettlementAmount               string `json:"settlementAmount"`
	SharingRefundNotifyNotifyType  string `json:"notify_type"`
	Amount                         string `json:"amount"`
	Version                        string `json:"version"`
	SharingRefundNotifyTotalAmount string `json:"total_amount"`
	TradeSn                        string `json:"tradeSn"`
	BgURL                          string `json:"bgUrl"`
	MerCharset                     string `json:"merCharset"`
	SharingRefundNotifyAccountDate string `json:"account_date"`
	Sign                           string `json:"sign"`
	State                          string `json:"state"`
	ChannelRecvTime                string `json:"channelRecvTime"`
	AccountDate                    string `json:"accountDate"`
	TradeStatus                    string `json:"tradeStatus"`
	SrcMercID                      string `json:"srcMercId"`
	OrderID                        string `json:"orderId"`
	UserCode                       string `json:"userCode"`
	TradeSource                    string `json:"tradeSource"`
	SignType                       string `json:"sign_type"`
	SharingRefundNotifyNotifyTime  string `json:"notify_time"`
	TradeNo                        string `json:"tradeNo"`
	NotifyTime                     string `json:"notifyTime"`
	NotifyType                     string `json:"notifyType"`
	SharingRefundNotifyTradeStatus string `json:"trade_status"`
	TotalAmount                    string `json:"totalAmount"`
	OutTradeNo                     string `json:"out_trade_no"`
	SharingRefundNotifyTradeNo     string `json:"trade_no"`
}
