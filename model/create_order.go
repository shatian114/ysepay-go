package yse_model

type CreateOrderReq struct { // 下单请求
	OrderId           string                                    `json:"orderId" v:"required"`          // 是否必填: Y;长度：20;商户订单号,商户系统生成的订单号，要求如下：; 1、须保证在商户端不重复。; 2、只能由大小写英文字母、数字、下划线及横杠组成。; 建议前8位为交易日期，如20220525，范围跨度支持包含当天在内的前后一天，示例值：202205256843192280647118。
	MsgCode           string                                    `json:"msgCode" v:"required"`          // 是否必填: Y;长度：265;报文编号,S3001即时到账、S3002担保交易，为空时默认S3001
	MercId            string                                    `json:"mercId" v:"required"`           // 是否必填: Y;长度：265;商户号
	BusiCode          string                                    `json:"busiCode" v:"required"`         // 是否必填: Y;长度：265;业务代码,默认传00510103，如果有通过会拓客分润需求，则传00510102，具体请咨询客户经理; 注意：如果有传非以上2个业务代码的请尽快修改，将在2024年4月25日开启强制校验
	ShopDate          string                                    `json:"shopDate" v:"required"`         // 是否必填: Y;长度：265;交易日期,商户系统的交易发生日期格式yyyyMMdd 示例值：20180525
	Amount            string                                    `json:"amount" v:"required"`           // 是否必填: Y;长度：265;交易金额,该笔订单的资金总额，单位：分。
	PaymentValidTime  string                                    `json:"paymentValidTime" v:"required"` // 是否必填: Y;长度：265;支付有效时间, 单位分钟，最小1分钟，最长不超过30分钟.。
	Currency          string                                    `json:"currency" v:"required"`         // 是否必填: Y;长度：265;币种,默认CNY（人民币）。
	Remark            string                                    `json:"remark,omitempty" v:""`         // 是否必填: N;长度：265;订单备注,该字段废弃
	Note              string                                    `json:"note,omitempty" v:""`           // 是否必填: N;长度：265;订单说明,可以用于商品备注，该字段会上送给渠道，建议必填，例如用户使用微信支付，则用户可以在微信账单看到订单说明
	BackUrl           string                                    `json:"backUrl" v:"required"`          // 是否必填: Y;长度：265;回调地址,交易成功异步通知到商户的后台地址，支持多个url进行异步通知，多个url用分隔符“,”分开，格式如：url1,url2,url3。
	LimitPay          string                                    `json:"limitPay,omitempty" v:""`       // 是否必填: N;长度：265;指定支付方式,(0或空不禁用,1禁用信用卡支付,2禁用花呗，3禁用花呗分期，4禁用所有信用支付类型); 微信，只支持1；云闪付，均不支持；支付宝，均支持
	PayMode           string                                    `json:"payMode" v:"required"`          // 是否必填: Y;长度：265;支付方式,; 26 支付宝生活号; 28 微信公众号; 29 微信小程序; 30 银联行业码支付
	Detail            CreateOrderDetail                         `json:"detail,omitempty" v:""`         // 是否必填: N;长度：265;商品购物明细,电商场景/担保交易时建议传值。
	StoreId           string                                    `json:"storeId,omitempty" v:""`        // 是否必填: N;长度：265;商户门店编号
	BuyerRealnameInfo `json:"buyerRealnameInfo,omitempty" v:""` // 是否必填: N;长度：265;实名信息
	IsRepeatPay       string                                    `json:"isRepeatPay,omitempty" v:""` // 是否必填: N;长度：265;是否允许重复下单,（空或者1代表是，0代表否）
	H5Join            string                                    `json:"h5Join,omitempty" v:""`      // 是否必填: N;长度：265;H5接入方式,商家如果本身是小程序接入收银台，则该字段默认传空，如果是H5接入请联系银盛客户经理
	IsFastPay         string                                    `json:"isFastPay,omitempty" v:""`   // 是否必填: N;长度：2;是否立即支付,传入01时，商户预下单之后拉起的是立即支付收银台页面，其他则进入普通收银台页面。; 说明：立即支付收银台与普通收银台相比没有确认按钮，用户从商户页面下单之后进入收银台页面，不需要确认直接进入密码支付
	MercHomeUrl       string                                    `json:"mercHomeUrl,omitempty" v:""` // 是否必填: N;长度：265;返回商户页面URL
	AppType           string                                    `json:"appType,omitempty" v:""`     // 是否必填: N;长度：265;应用类型
}

type CreateOrderRes struct { // 下单响应
	ReqMsgId           string `json:"reqMsgId,omitempty" v:""`           // 是否必填: N;长度：265;请求流水号
	Amount             string `json:"amount,omitempty" v:""`             // 是否必填: N;长度：265;订单金额,请求对应的订单金额
	MercId             string `json:"mercId,omitempty" v:""`             // 是否必填: N;长度：265;商户号
	OrderId            string `json:"orderId,omitempty" v:""`            // 是否必填: N;长度：265;商户订单号,商户请求对应的商户订单号
	OrderCreateTime    string `json:"orderCreateTime,omitempty" v:""`    // 是否必填: N;长度：265;创建订单时间,创建时间，格式"yyyy-MM-dd HH:mm:ss"
	OrderEfficientTime string `json:"orderEfficientTime,omitempty" v:""` // 是否必填: N;长度：265;订单有效截至时间,订单有效时间，格式"yyyy-MM-dd HH:mm:ss"
	EncryData          string `json:"encryData,omitempty" v:""`          // 是否必填: N;长度：265;加密校验参数,加密校验信息，拉起银盛收银台需要
	PayUrl             string `json:"payUrl,omitempty" v:""`             // 是否必填: N;长度：265;收银台地址信息
	AppId              string `json:"appId,omitempty" v:""`              // 是否必填: N;长度：265;appId
}

type GoodsDetail struct { // 下单商品详情
	GoodsId      string `json:"goodsId" v:"required"`        // 是否必填: Y;长度：265;商户商品编码,由半角的⼤⼩写字母，数字，下划线，中划线中的⼀种或者⼏种组成，Shouji_abcef_001-001
	WxpayGoodsId string `json:"wxpayGoodsId,omitempty" v:""` // 是否必填: N;长度：265;微信商品编号,微信支付定义的统⼀商品编号,微信支付定义的统⼀商品编号
	GoodsName    string `json:"goodsName,omitempty" v:""`    // 是否必填: N;长度：265;商品名称,商品的实际名称
	Quantity     string `json:"quantity" v:"required"`       // 是否必填: Y;长度：265;商品数量,只能是正整数 ,1000
	Price        string `json:"price" v:"required"`          // 是否必填: Y;长度：265;商品单价,单位：分
}

type BuyerRealnameInfo struct { // 下单实名信息
	CertType string `json:"certType" v:"required"` // 是否必填: Y;长度：265;证件类型,只支持身份证，枚举值：1
	CertName string `json:"certName" v:"required"` // 是否必填: Y;长度：265;证件名称
	CertNo   string `json:"certNo" v:"required"`   // 是否必填: Y;长度：265;证件号码
}
type CreateOrderDetail struct { // 下单商品购物明细
	CostPrice    string        `json:"costPrice,omitempty" v:""`  // 是否必填: N;长度：265;订单原价,单位：分
	ReceiptId    string        `json:"receiptId,omitempty" v:""`  // 是否必填: N;长度：265;商品小票,长度不能超过32位 Ys001001
	GoodsDetails []GoodsDetail `json:"goodsDetails" v:"required"` // 是否必填: Y;长度：265;商品列表
}

type CreateOrderUnifyWechatReq struct { // 微信统一扫码下单请求
	OrderId          string `json:"orderId" v:"required"`            // 是否必填: Y;长度：32;商户订单号,商户系统生成的订单号，要求如下：; 1、须保证在商户端不重复。; 2、只能由大小写英文字母、数字、下划线及横杠组成。; 建议前8位为交易日期，如20220525，范围跨度支持包含当天在内的前后一天，示例值：202205256843192280647118。
	ShopDate         string `json:"shopDate,omitempty" v:""`         // 是否必填: N;长度：8;交易日期,商户系统的交易发生日期格式yyyyMMdd 20180525
	Note             string `json:"note" v:"required"`               // 是否必填: Y;长度：125;订单备注,订单备注
	Amount           string `json:"amount" v:"required"`             // 是否必填: Y;长度：10;订单金额,该笔订单的资金总额，单位为：分。必须为正整数
	Currency         string `json:"currency,omitempty" v:""`         // 是否必填: N;长度：3;币种,默认CNY（⼈民币）
	PayeeMercId      string `json:"payeeMercId" v:"required"`        // 是否必填: Y;长度：20;收款方商户号
	PayeeMercName    string `json:"payeeMercName,omitempty" v:""`    // 是否必填: N;长度：50;收款方商户名称,收款商户号对应商户名称
	TimeOutExpress   string `json:"timeOutExpress" v:"required"`     // 是否必填: Y;长度：265;订单失效时间,设置未付款交易的超时时间，⼀旦超时，该笔交易就会自动被关闭。单 ; 位：分钟。例如：30，则订单失效时间为30分钟。
	ExtraCommonParam string `json:"extraCommonParam,omitempty" v:""` // 是否必填: N;长度：1000;公用回传参数,扩展参数，调用方可用于回调时携带业务参数，异步通知时会返回该参数。
	BusiCode         string `json:"busiCode" v:"required"`           // 是否必填: Y;长度：10;业务代码,请联系银盛客户经理获取。注意：业务代码非固定值，不同到账方式需 ; 要传不同的业务代码
	SubOpenId        string `json:"subOpenId" v:"required"`          // 是否必填: Y;长度：128;用户在商户appid下的唯⼀标识,用户在商户appid下的唯⼀标识。可参考https://pay.weixin.qq.com/ ; wiki/doc/apiv3/terms_definition/chapter1_1_3.shtml#part-3获 ; 取
	IsMiniPg         string `json:"isMiniPg,omitempty" v:""`         // 是否必填: N;长度：1;是否为⼩程序支付,1：微信⼩程序支付； ; 2：微信公众号支付（默认值）；
	AppId            string `json:"appId" v:"required"`              // 是否必填: Y;长度：32;商户公众平台ID, ; 商户在微信侧申请的⼩程序或公众号APPID。当发起公众号支付时，需 ; 填写公众号的APPID(应用ID)；当发起⼩程序支付时，需填写⼩程序的 ; APPID(应用ID)
	LimitCreditPay   string `json:"limitCreditPay,omitempty" v:""`   // 是否必填: N;长度：2;是否限制信用卡,0：不限制（默认值）； ; 1：禁用信用卡支付；
	AllowRepeatPay   string `json:"allowRepeatPay,omitempty" v:""`   // 是否必填: N;长度：2;是否允许多次支付,是否允许多次支付,Y：允许;N：不允许（参数为空或者Y时，默认该笔 ; 订单交易状态非成功状态时，支持选择其他支付方式继续付款， 当为N ; 时表示该笔订单交易状态为失败状态时，不支持选择其他支付方式继续 ; 付款）
	FailNotifyUrl    string `json:"failNotifyUrl,omitempty" v:""`    // 是否必填: N;长度：190;失败通知地址,失败通知地址（交易状态为失败时，银盛通知商户网站里指定的页面ht ; tp路径，支持多个url进行异步通知，多个url用分隔符“,”分开，格式如 ; ：url1,url2,url3）注：只有不允许重复支付的交易才会通知
	NotifyUrl        string `json:"notifyUrl" v:"required"`          // 是否必填: Y;长度：190;通知地址,交易成功异步通知到商户的后台地址 ，详见《微信JS下单异步通知》
	SrcIP            string `json:"srcIP,omitempty" v:""`            // 是否必填: N;长度：16;发起方IP,请求方ip。例如：112.112.112.112
	PayerIP          string `json:"payerIP,omitempty" v:""`          // 是否必填: N;长度：16;付款方IP,付款方ip。例如：112.112.112.113
	GpsInfo          string `json:"gpsInfo,omitempty" v:""`          // 是否必填: N;长度：265;付款方gps信息,付款方gps信息，示例值："{\"userGps\":{\"lat\":24.48579,\"lng\":118.126076}}"。
	AgentNo          string `json:"agentNo,omitempty" v:""`          // 是否必填: N;长度：265;代理商编号,发起方银盛代理商编号。
	PaygateNo        string `json:"paygateNo,omitempty" v:""`        // 是否必填: N;长度：265;支付网关编号,9000010：银联-微信；10010001：网联-微信
	MsgCode          string `json:"msgCode,omitempty" v:""`          // 是否必填: N;长度：265;报文编号,S3001即时到账、S3002担保交易，为空时默认S3001
	OrderExtendInfo  string `json:"orderExtendInfo,omitempty" v:""`  // 是否必填: N;长度：265;扩展参数,扩展参数，示例：{\"qr_code_no\":\"20251210247697961\"}
	YsStoresNo       string `json:"ysStoresNo,omitempty" v:""`       // 是否必填: N;长度：265;银盛门店编号,银盛侧的门店编号
	YsCashierNo      string `json:"ysCashierNo,omitempty" v:""`      // 是否必填: N;长度：265;银盛收银员编号,银盛侧的收银员编号
}

type CreateOrderUnifyWechatRes struct { // 微信统一扫码下单响应
	ReqMsgId string                  `json:"reqMsgId,omitempty" v:""` // 是否必填: N;长度：265;请求流水号,同请求
	Data     CreateOrderUnifyResData `json:"data,omitempty" v:""`     // 是否必填: N;长度：265;响应业务数据,请求成功时响应
}
type CreateOrderUnifyResData struct { // 微信统一扫码下单响应数据
	OrderId          string `json:"orderId,omitempty" v:""`          // 是否必填: N;长度：265;商户订单号
	TradeSn          string `json:"tradeSn,omitempty" v:""`          // 是否必填: N;长度：265;交易流水号
	TradeStatus      string `json:"tradeStatus,omitempty" v:""`      // 是否必填: N;长度：265;交易状态,00-交易成功 ; 02-支付成功,合单交易渠道成功但参与营销，需⼈⼯调账待支付 ; 11-支付中（已向渠道跳转） ; 12-部分支付 ; 13-全部已支付(担保交易使用)等待发货 ; 14-未确认收货（等待付款方确认） ; 60-提现登记 ; 61-提现完成 ; 70-自动关闭处理中 ; 80-部分退款 ; 81-全部退款 ; 97-客户主动关闭 ; 99-到时关闭' ; 01-支付成功 ; 98-支付失败 ; 96-交易失败（超金额支付） ; 50-业务超时 ; 95-关闭订单 ; 94-无支付订 ; 93-订单撤销
	Amount           string `json:"amount,omitempty" v:""`           // 是否必填: N;长度：265;订单金额,单位：分
	Currency         string `json:"currency,omitempty" v:""`         // 是否必填: N;长度：265;交易币种
	ExtraCommonParam string `json:"extraCommonParam,omitempty" v:""` // 是否必填: N;长度：265;公共回传参数
	JsapiPayInfo     string `json:"jsapiPayInfo,omitempty" v:""`     // 是否必填: N;长度：265;微信返回参数,JSON格式。具体字段参考以下jsapiPayInfo详细参数
}

type CreateOrderUnifyWechatJs struct { // 微信统一扫码下单js调用数据
	AppId     string `json:"appId" v:"required"`     // 是否必填: Y;长度：32;商户APPID,商户在微信侧申请的APPID，接⼝返回jsapi_pay_info中的 ; 信息 wxd678efh567hg6787
	TimeStamp string `json:"timeStamp" v:"required"` // 是否必填: Y;长度：32;时间戳,时间戳，接⼝返回jsapi_pay_info中的信息
	NonceStr  string `json:"nonceStr" v:"required"`  // 是否必填: Y;长度：32;随机字机串,随机字机串，接⼝返回jsapi_pay_info中的信息 5K8264ILT ; KCH16CQ2502SI8
	Package   string `json:"package" v:"required"`   // 是否必填: Y;长度：128;订单详情扩展字 符串,订单详情扩展字符串，接⼝返回jsapi_pay_info中的信息 pr ; epay_id=u802345jgfjsdfg
	SignType  string `json:"signType" v:"required"`  // 是否必填: Y;长度：32;签名方式, ; 签名方式，接⼝返回 jsapi_pay_info中的信息 MD5
	PaySign   string `json:"paySign" v:"required"`   // 是否必填: Y;长度：256;签名,签名，接⼝返回jsapi_pay_info中的信息 C380BEC2BFD7 ; 27A4B6845132
}

type CreateOrderUnifyAlipayReq struct { // 聚合扫码支付宝生活号下单请求
	OrderId          string `json:"orderId" v:"required"`            // 是否必填: Y;长度：32;订单号,商户系统生成的订单号，要求如下：; 1、须保证在商户端不重复。; 2、只能由大小写英文字母、数字、下划线及横杠组成。; 建议前8位为交易日期，如20220525，范围跨度支持包含当天在内的前后一天，示例值：202205256843192280647118。
	ShopDate         string `json:"shopDate,omitempty" v:""`         // 是否必填: N;长度：8;交易日期,商户系统的交易发生日期格式yyyyMMdd， 示例值：20180525。
	Note             string `json:"note" v:"required"`               // 是否必填: Y;长度：125;订单备注
	Amount           string `json:"amount" v:"required"`             // 是否必填: Y;长度：10;订单总金额,该笔订单的资金总额，必须为正整数，单位：分。
	Currency         string `json:"currency,omitempty" v:""`         // 是否必填: N;长度：3;交易币种,默认CNY（人民币）。
	PayeeMercId      string `json:"payeeMercId" v:"required"`        // 是否必填: Y;长度：20;收款商户号
	PayeeMercName    string `json:"payeeMercName,omitempty" v:""`    // 是否必填: N;长度：265;商户名称,收款商户号对应商户名称。
	TimeOutExpress   string `json:"timeOutExpress" v:"required"`     // 是否必填: Y;长度：6;交易超时时间,设置未付款交易的超时时间，一旦超时，该笔交易就会自动被关闭。单位：分钟。注意：设置了未付款交易超时时间的情况下，若我司在限定时间内没有收到成功支付通知，则会关闭交易，关闭后该笔交易若付款方支付成功的情况下，会自动原路退款至付款方。
	ExtraCommonParam string `json:"extraCommonParam,omitempty" v:""` // 是否必填: N;长度：1000;公用回传参数,扩展参数，调用方可用于回调时携带业务参数，异步通知时会返回该参数。
	BusiCode         string `json:"busiCode" v:"required"`           // 是否必填: Y;长度：10;业务代码,请联系银盛客户经理获取。; 注意：业务代码非固定值，不同到账方式需要传不同的业务代码。
	LimitCreditPay   string `json:"limitCreditPay,omitempty" v:""`   // 是否必填: N;长度：265;是否限制贷记卡,0：不限制（默认值）；; 1：禁用信用卡；; 2：禁用花呗；; 3：禁用花呗分期；; 4：禁用所有信用支付类型；
	AllowRepeatPay   string `json:"allowRepeatPay,omitempty" v:""`   // 是否必填: N;长度：265;是否允许多次支付,Y：允许；N：不允许（参数为空或者Y时，默认该笔订单交易状态非成功状态时，支持选择其他支付方式继续付款， 当为N时表示该笔订单交易状态为失败状态时，不支持选择其他支付方式继续付款）
	NotifyUrl        string `json:"notifyUrl,omitempty" v:""`        // 是否必填: N;长度：190;通知地址
	FailNotifyUrl    string `json:"failNotifyUrl,omitempty" v:""`    // 是否必填: N;长度：190;失败通知地址,失交易状态为失败时，银盛支付服务器主动通知商户网站里指定的页面http路径，支持多个url进行异步通知，多个url用分隔符“,”分开，格式如：url1,url2,url3。; 注：只有不允许重复支付的交易才会通知。
	BuyerLogonId     string `json:"buyerLogonId,omitempty" v:""`     // 是否必填: N;长度：100;买家支付宝账号,和bugerId不能同时为空。
	BuyerId          string `json:"buyerId,omitempty" v:""`          // 是否必填: N;长度：100;买家支付宝用户ID,买家支付宝用户ID，和buyerLogonId不能同时为空。bugerId获取方法：https://doc.open.alipay.com/doc2/detail.htm?spm=a219a.7629140.0.0.8ujLD6&treeId=115&articleId=104114&docType=1通过网页授权获取用户信息，同步响应结果中的user_id对应文档中的bugerId。
	SrcIP            string `json:"srcIP,omitempty" v:""`            // 是否必填: N;长度：16;发起方IP
	PayerIP          string `json:"payerIP,omitempty" v:""`          // 是否必填: N;长度：16;付款方IP
	GpsInfo          string `json:"gpsInfo,omitempty" v:""`          // 是否必填: N;长度：265;付款方gps信息,示例值："{\"userGps\":{\"lat\":24.48579,\"lng\":118.126076}}"。
	AgentNo          string `json:"agentNo,omitempty" v:""`          // 是否必填: N;长度：265;代理商编号,发起方的银盛代理商编号。
	PaygateNo        string `json:"paygateNo,omitempty" v:""`        // 是否必填: N;长度：265;支付网关编号,900000001：银联-支付宝；10000012：网联-支付宝。
	HbFqNum          string `json:"hbFqNum,omitempty" v:""`          // 是否必填: N;长度：265;花呗分期期数,仅在支付宝支付的时候生效,目前期数支持3,6,12,24。
	StoreId          string `json:"storeId,omitempty" v:""`          // 是否必填: N;长度：32;商户门店编号
	AlipayStoreId    string `json:"alipayStoreId,omitempty" v:""`    // 是否必填: N;长度：28;支付宝的店铺编号
	OperatorId       string `json:"operatorId,omitempty" v:""`       // 是否必填: N;长度：28;商户操作员编号
	TerminalId       string `json:"terminalId,omitempty" v:""`       // 是否必填: N;长度：32;商户机具终端编号
	BusinessParams   string `json:"businessParams,omitempty" v:""`   // 是否必填: N;长度：32;支付宝花呗分期专属字段
	FqType           string `json:"fqType,omitempty" v:""`           // 是否必填: N;长度：265;分期类型,值为00表示花呗分期,值01表示信用卡分期。不填默认为花呗分期。
	YsStoresNo       string `json:"ysStoresNo,omitempty" v:""`       // 是否必填: N;长度：265;银盛门店编号,银盛侧的门店编号
	YsCashierNo      string `json:"ysCashierNo,omitempty" v:""`      // 是否必填: N;长度：265;银盛收银员编号,银盛侧的收银员编号
}
type CreateOrderUnifyAlipayRes struct {
	Data       Data   `json:"data"`
	SystemCode string `json:"systemCode"`
	ReqMsgID   string `json:"reqMsgId"`
}

type Data struct {
	Amount       string `json:"amount"`
	OrderID      string `json:"orderId"`
	TradeStatus  string `json:"tradeStatus"`
	Currency     string `json:"currency"`
	JsapiPayInfo string `json:"jsapiPayInfo"`
	TradeSn      string `json:"tradeSn"`
}
