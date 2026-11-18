package yse_model

// /////////////  开通订单分账
type OpenSharingReq struct { // 开通订单分账请求
	MercId                      string                 `json:"mercId" v:"required"`                        // 是否必填: Y;长度：265;商户号
	RevenueSharingType          string                 `json:"revenueSharingType" v:"required"`            // 是否必填: Y;长度：265;分账类型,分账类型(0余额分账1订单分账)
	RevenueSharingRule          string                 `json:"revenueSharingRule,omitempty" v:""`          // 是否必填: N;长度：265;分账规则,订单分账必填，分账规则( 1 接口动态分账 2 按百分比 3 按固定金额 ); 场景：; 接口动态分账：开通时无需指定商户，发起分账交易时可自主指定商户。; 按百分比/按固定金额：开通时需要指定商户，发起分账交易时按开通时指定商户进行分账。
	MinSplitRatio               string                 `json:"minSplitRatio,omitempty" v:""`               // 是否必填: N;长度：265;最低分账比例,交易收款方最低分账比例（%），分账规则=1 时必填0-100 （2位小数）
	MaxSplitMerchantCount       string                 `json:"maxSplitMerchantCount,omitempty" v:""`       // 是否必填: N;长度：265;参与分账商户数,参与分账商户数，分账规则=1时必填1-20
	RevenueSharingScene         string                 `json:"revenueSharingScene" v:"required"`           // 是否必填: Y;长度：265;分账场景,场景00(代售保险) ; 场景01(下级单位财政款) ; 场景02115(司机/车队劳务费(网络货运平台)) ; 场景03(供应商贷款) ; 场景04(分销商/经销商/站点/加盟店应收账款(线上)) ; 场景05(分销商/经销商/站点/加盟店应收账款(线下)) ; 场景06(推广返佣) ; 场景07(系统商/平台商服务费); 场景08(平台补贴) ; 场景09(货主应收账款) ; 场景10(入驻商户应收账款); 场景11(用户应收储值款); 场景12(公司内部/集团/股东资金往来); 场景13(资方应收账款) ; 场景14(助贷服务费/担保费等); 场景15(司机/车队劳务费(交通运输))
	RevenueSharingDetails       []RevenueSharingDetail `json:"revenueSharingDetails,omitempty" v:""`       // 是否必填: N;长度：265;分账明细,订单分账必填，分账明细，分账规则=2、3时必填；需包含交易商户；分账规则=2 时分账商户相加比例为100%；; 分账规则=3 时交易商户分账金额传0
	RevenueSharingPics          []RevenueSharingPic    `json:"revenueSharingPics,omitempty" v:""`          // 是否必填: N;长度：265;图片信息
	RevenueSharingRate          string                 `json:"revenueSharingRate,omitempty" v:""`          // 是否必填: N;长度：265;余额分账费率,余额分账费率（%），余额分账时必填 [0.02-100)（2位小数）
	NotifyUrl                   string                 `json:"notifyUrl,omitempty" v:""`                   // 是否必填: N;长度：265;回调地址,申请单状态变更通知（申请单状态（待签约、审核中（签约完成）、审核通过、审核驳回））
	RevenueSharingFeeRateBearer string                 `json:"revenueSharingFeeRateBearer,omitempty" v:""` // 是否必填: N;长度：265;余额分账费率承担方,1分账商户，2分账接收方。默认=1
	RevenueSharingRuleSource    string                 `json:"revenueSharingRuleSource,omitempty" v:""`    // 是否必填: N;长度：265;余额分账规则来源,1=接口发起，2=门户发起，3=接口+门户。默认=2
	OrderRevenueSharingRate     string                 `json:"orderRevenueSharingRate,omitempty" v:""`     // 是否必填: N;长度：265;订单分账费率,订单分账费率（%），订单分账时必填 [0.02-100)（2位小数）注：如果存量订单分账费率为0，支持费率传0
	ContractType                string                 `json:"contractType,omitempty" v:""`                // 是否必填: N;长度：265;签约类型,签约类型：1=纸质签约，2=电子签约。纸质签约时，图片信息字段中，分账确认书首页照片、分账确认书盖章页照片必传。签约类型未传时默认为电子签约。（暂未实现）
}

type OpenSharingRes struct { // 开通订单分账响应
	OrderNo string `json:"orderNo"` //申请单号
}

type RevenueSharingDetail struct { // 开通分账请求明细
	DivMercId string `json:"divMercId,omitempty" v:""` // 是否必填: N;长度：265;分账商户号
	Ratio     string `json:"ratio,omitempty" v:""`     // 是否必填: N;长度：265;分账比例,分账比例（%），当分账规则=2时必填 ，且分账明细中比例加起来等于100，例：10.25，大于0（2位小数）
	Amount    string `json:"amount,omitempty" v:""`    // 是否必填: N;长度：265;分账金额,分账金额（元/2位小数），当分账规则=3时必填，分账商户大于0（交易商户等于0）
	FeeFlag   string `json:"feeFlag,omitempty" v:""`   // 是否必填: N;长度：265;是否手续费承担方,是否手续费承担方，1是 2否；手续费承担方有且只有一个；如果所有分账接收方该字段均传空，默认交易商户号作为手续费承担方。）
}

type RevenueSharingPic struct { // 开通分账请求图片
	PicType string `json:"picType,omitempty" v:""` // 是否必填: N;长度：265;图片类型,图片类型（枚举同图片上传接口）
	PicUrl  string `json:"picUrl,omitempty" v:""`  // 是否必填: N;长度：265;图片地址,图片接口返回的地址
}
type OpenSharingQueryReq struct { // 开通订单分账进度查询请求
	OrderNo string `json:"orderNo"` //申请单号
}

type OpenSharingQueryRes struct { // 开通订单分账进度查询请求
	OrderInfo OpenSharingQueryResOrder `json:"orderInfo"`
}

type OpenSharingQueryResOrder struct { // 开通分账请求查询响应订单内容
	OrderNo string `json:"orderNo,omitempty" v:""` // 是否必填: N;长度：265;申请单号
	Status  string `json:"status,omitempty" v:""`  // 是否必填: N;长度：265;当前状态,0.已完成   1.待审核   2:审核中 3.待签约   9.审核拒绝
	Remark  string `json:"remark,omitempty" v:""`  // 是否必填: N;长度：265;备注
	SignUrl string `json:"signUrl,omitempty" v:""` // 是否必填: N;长度：265;签约链接,3:待签约时返回签约链接
	SignId  string `json:"signId,omitempty" v:""`  // 是否必填: N;长度：265;签约ID,签约ID。后签约模式使用（暂未实现）
}

// /////////////// 关闭订单分账
type CloseSharingReq struct { // 关闭订单分账请求
	MercId             string `json:"mercId"`             // 商户号
	RevenueSharingType string `json:"revenueSharingType"` // 分账类型,分账类型(0余额分账1订单分账)
}

type SharingRegistReq struct { // 登记分账请求
	NotifyUrl       string                `json:"notifyUrl" v:"required"`       // 是否必填: Y;长度：190;异步通知地址,异步通知地址，商户通过该地址接收异步通知
	OrderId         string                `json:"orderId" v:"required"`         // 是否必填: Y;长度：32;订单号
	PayeeMercId     string                `json:"payeeMercId" v:"required"`     // 是否必填: Y;长度：20;主商户号,原交易收款方商户号
	Amount          string                `json:"amount" v:"required"`          // 是否必填: Y;长度：10;原订单交易金额,必需为正整数。单位：分
	SysFlag         string                `json:"sysFlag,omitempty" v:""`       // 是否必填: N;长度：2;系统标志,DD：订单交易，DS：代收交易，默认是DD，订单交易
	IsDivistion     string                `json:"isDivistion" v:"required"`     // 是否必填: Y;长度：2;原订单是否参与分账,01：是，02否
	IsDivisionAgain string                `json:"isDivisionAgain" v:"required"` // 是否必填: Y;长度：1;是否重新分账,Y：是，N：否
	DivisionMode    string                `json:"divisionMode" v:"required"`    // 是否必填: Y;长度：2;分账模式,01 ：比例，02：金额
	DivList         []SharingRegistDetail `json:"divList,omitempty" v:""`       // 是否必填: N;长度：265;分账明细,参与分账时不可以空，明细中的是否必填按照分账来描述。
	TradeSn         string                `json:"tradeSn,omitempty" v:""`       // 是否必填: N;长度：30;银盛原交易流水号
}

type SharingRegistDetail struct { // 登记分账列表明细
	DivMercId   string  `json:"divMercId" v:"required"`   // 是否必填: Y;长度：20;分账商户号,参与分账商户号
	IsChargeFee string  `json:"isChargeFee" v:"required"` // 是否必填: Y;长度：2;是否收取手续费,01：是，02否。若配置手续费收支分离，该参数配置的手续费承担方不生效
	DivAmount   string  `json:"divAmount,omitempty" v:""` // 是否必填: N;长度：10;分账金额,必需为正整数。单位：分单位：分
	DivRatio    float64 `json:"divRatio,omitempty" v:""`  // 是否必填: N;长度：265;分账比例,divMode=01时必填，小数格式，最小长度1位，最大4位小数位（例：0.2365）, 分账明细中所有分账比例之和必须为1。
}

type SharingRegistRes struct { // 登记分账响应
	BusinessData `json:"businessData,omitempty" v:""` // 是否必填: N;长度：265;响应数据
	ReqMsgId     string                               `json:"reqMsgId" v:"required"` // 是否必填: Y;长度：265;请求流水号
}

type SharingRegistResInfo struct { // 登记分账响应内容
	ReturnInfo string `json:"returnInfo"` //返回说明
}

type SharingRegistQueryReq struct { // 登记分账查询请求
	OrderId string `json:"orderId" v:"required"`   // 是否必填: Y;长度：30;订单号,商户系统生成的订单号
	SysFlag string `json:"sysFlag,omitempty" v:""` // 是否必填: N;长度：2;系统标志,DD：订单交易，DS：代收交易，默认是DD，订单交易
	TradeSn string `json:"tradeSn,omitempty" v:""` // 是否必填: N;长度：30;银盛原交易流水号
}

type SharingRegistQueryRes struct { //登记分账查询响应
	Data       SharingRegistQueryResData `json:"data"`
	SystemCode string                    `json:"systemCode"`
	ReqMsgID   string                    `json:"reqMsgId"`
}

type SharingRegistQueryResData struct {
	Note               string    `json:"note"`
	Amount             string    `json:"amount"`
	OrderID            string    `json:"orderId"`
	ShopDate           string    `json:"shopDate"`
	Fee                string    `json:"fee"`
	UpdateTime         string    `json:"updateTime"`
	DivisionStatusCode string    `json:"divisionStatusCode"`
	DivList            []DivList `json:"divList"`
	FeePayer           string    `json:"feePayer"`
	DivisionStatus     string    `json:"divisionStatus"`
	SrcMercID          string    `json:"srcMercId"`
	TradeSn            string    `json:"tradeSn"`
	DivisionMode       string    `json:"divisionMode"`
}

type DivList struct {
	DivMercID string `json:"divMercId"`
	DvsnFee   string `json:"dvsnFee"`
	DivAmount string `json:"divAmount"`
}

type SharingUploadPicReq struct { // 开通分账前先上传图片
	PicType string `json:"picType"`
	PicStr  string `json:"picStr"` // 图片base64
}

type SharingUploadPicRes struct { // 开通分账前先上传图片
	PicUrl string `json:"picUrl"`
}

///////////////// 订单分账退款

type SharingRefundReq struct { // 银盛分账退款请求
	OrderId             string                    `json:"orderId" v:"required"`               // 是否必填: Y;长度：32;商户原交易订单号,示例值：201805256843192280647118
	TradeSn             string                    `json:"tradeSn,omitempty" v:""`             // 是否必填: N;长度：30;银盛原交易流水号
	ShopDate            string                    `json:"shopDate" v:"required"`              // 是否必填: Y;长度：8;商户系统的交易发生日期,格式yyyyMMdd 示例值：20180525
	Amount              string                    `json:"amount" v:"required"`                // 是否必填: Y;长度：10;需要退款的金额,必须为正整数。该金额不能大于订单金额。单位为：分
	RefundReason        string                    `json:"refundReason" v:"required"`          // 是否必填: Y;长度：50;退款的原因说明,该参数最长为50个汉字
	RefundOrderId       string                    `json:"refundOrderId" v:"required"`         // 是否必填: Y;长度：32;商户系统生成的退款订单号,标识一次退款请求，同一笔交易多次退款需要保证唯一。该参数支持汉字，最大长度为16个。用同一订单号同一退款流水号继续退款时，则会返回第一次退款的结果不会继续退款 示例值：RD2012061713107
	IsDivistion         string                    `json:"isDivistion,omitempty" v:""`         // 是否必填: N;长度：2;原交易是否参与分账,01或空-是；02-否
	DivisionMode        string                    `json:"divisionMode,omitempty" v:""`        // 是否必填: N;长度：2;原交易分账模式,01 -比例；02-金额
	NotifyUrl           string                    `json:"notifyUrl" v:"required"`             // 是否必填: Y;长度：190;异步通知地址,异步通知地址
	MsgCode             string                    `json:"msgCode" v:"required"`               // 是否必填: Y;长度：265;报文编号,S3106-担保交易退款；S3107-分账退款
	RefundDivInfoList   []SharignRefundDivInfo    `json:"refundDivInfoList,omitempty" v:""`   // 是否必填: N;长度：265;退款分账明细,分账退款金额总和等于退款金额，当原交易参与分账时，该参数必传，原交易不参与分账，
	OriginalDivInfoList []SharingRefundOriginInfo `json:"originalDivInfoList,omitempty" v:""` // 是否必填: N;长度：265;原订单交易分帐信息,当原交易参与分账时，该参数必传，原交易不参与分账，该参数无值。
	SrcIP               string                    `json:"srcIP" v:"required"`                 // 是否必填: Y;长度：16;发起方IP,示例：10.20.61.119
}

type SharignRefundDivInfo struct { // 银盛分账退款请求退款分账明细
	DivMercId    string `json:"divMercId,omitempty" v:""`    // 是否必填: N;长度：20;分账退款商户号,参与分账商户号
	DivAmount    string `json:"divAmount,omitempty" v:""`    // 是否必填: N;长度：10;分账退款金额,必须为正整数，单位为：分单位：分
	DivisionMode string `json:"divisionMode,omitempty" v:""` // 是否必填: N;长度：2;原交易分账模式
}

type SharingRefundOriginInfo struct { // 银盛分账退款请求原订单分账明细
	DivMercId   string  `json:"divMercId,omitempty" v:""`   // 是否必填: N;长度：20;原订单分账收款方商户号,参与分账商户号
	DivAmount   string  `json:"divAmount,omitempty" v:""`   // 是否必填: N;长度：10;原订单分账金额,必须为正整数，单位为：分单位：分
	DivRatio    float64 `json:"divRatio,omitempty" v:""`    // 是否必填: N;长度：265;原订单分账比例,原订单分账模式为按比如时必填，小数格式，最大4位小数位
	IsChargeFee string  `json:"isChargeFee,omitempty" v:""` // 是否必填: N;长度：2;是否收取手续费,01-是；02-否
}

type SharingRefundRes struct { // 银盛分账退款响应
	BusinessData []SharingRefundResData `json:"businessData,omitempty" v:""` // 是否必填: N;长度：265;响应数据
	ReqMsgId     string                 `json:"reqMsgId" v:"required"`       // 是否必填: Y;长度：265;请求流水号
}

type SharingRefundResData struct { // 银盛分账退款响应数据
	SrcMercId     string `json:"srcMercId" v:"required"`     // 是否必填: Y;长度：265;发起方商户号
	Amount        string `json:"amount" v:"required"`        // 是否必填: Y;长度：265;本次发生的退款金额,单位：分
	RefundOrderId string `json:"refundOrderId" v:"required"` // 是否必填: Y;长度：265;商户系统生成的订单号
	AccDate       string `json:"accDate" v:"required"`       // 是否必填: Y;长度：265;退款会计日期,格式yyyyMMdd。 示例值：20140724
}
