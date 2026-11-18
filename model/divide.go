package yse_model

type SharingOpenReq struct { // 开通分账请求
	MercId                      string `json:"mercId" v:"required"`                        // 是否必填: Y;长度：265;商户号
	RevenueSharingType          string `json:"revenueSharingType" v:"required"`            // 是否必填: Y;长度：265;分账类型,分账类型(0余额分账1订单分账)
	RevenueSharingRule          string `json:"revenueSharingRule,omitempty" v:""`          // 是否必填: N;长度：265;分账规则,订单分账必填，分账规则( 1 接口动态分账 2 按百分比 3 按固定金额 ); 场景：; 接口动态分账：开通时无需指定商户，发起分账交易时可自主指定商户。; 按百分比/按固定金额：开通时需要指定商户，发起分账交易时按开通时指定商户进行分账。
	MinSplitRatio               string `json:"minSplitRatio,omitempty" v:""`               // 是否必填: N;长度：265;最低分账比例,交易收款方最低分账比例（%），分账规则=1 时必填0-100 （2位小数）
	MaxSplitMerchantCount       string `json:"maxSplitMerchantCount,omitempty" v:""`       // 是否必填: N;长度：265;参与分账商户数,参与分账商户数，分账规则=1时必填1-20
	RevenueSharingScene         string `json:"revenueSharingScene" v:"required"`           // 是否必填: Y;长度：265;分账场景,场景00(代售保险) ; 场景01(下级单位财政款) ; 场景02115(司机/车队劳务费(网络货运平台)) ; 场景03(供应商贷款) ; 场景04(分销商/经销商/站点/加盟店应收账款(线上)) ; 场景05(分销商/经销商/站点/加盟店应收账款(线下)) ; 场景06(推广返佣) ; 场景07(系统商/平台商服务费); 场景08(平台补贴) ; 场景09(货主应收账款) ; 场景10(入驻商户应收账款); 场景11(用户应收储值款); 场景12(公司内部/集团/股东资金往来); 场景13(资方应收账款) ; 场景14(助贷服务费/担保费等); 场景15(司机/车队劳务费(交通运输))
	RevenueSharingRate          string `json:"revenueSharingRate,omitempty" v:""`          // 是否必填: N;长度：265;余额分账费率,余额分账费率（%），余额分账时必填 [0.02-100)（2位小数）
	NotifyUrl                   string `json:"notifyUrl,omitempty" v:""`                   // 是否必填: N;长度：265;回调地址,申请单状态变更通知（申请单状态（待签约、审核中（签约完成）、审核通过、审核驳回））
	RevenueSharingFeeRateBearer string `json:"revenueSharingFeeRateBearer,omitempty" v:""` // 是否必填: N;长度：265;余额分账费率承担方,1分账商户，2分账接收方。默认=1
	RevenueSharingRuleSource    string `json:"revenueSharingRuleSource,omitempty" v:""`    // 是否必填: N;长度：265;余额分账规则来源,1=接口发起，2=门户发起，3=接口+门户。默认=2
	OrderRevenueSharingRate     string `json:"orderRevenueSharingRate,omitempty" v:""`     // 是否必填: N;长度：265;订单分账费率,订单分账费率（%），订单分账时必填 [0.02-100)（2位小数）
}

type SharingOpenRes struct { // 开通分账响应
	OrderNo string `json:"orderNo,omitempty" v:""` // 是否必填: N;长度：265;申请单号
}

type DivideReceiveConfigReq struct { // 配置分账接收方请求
	DivideMerchantNo        string `json:"divideMerchantNo" v:"required"`        // 是否必填: Y;长度：20;分账商户
	DivideReceiveMerchantNo string `json:"divideReceiveMerchantNo" v:"required"` // 是否必填: Y;长度：20;分账接收方商户
	OperType                string `json:"operType" v:"required"`                // 是否必填: Y;长度：1;操作类型,1.新增 2.删除
}

type DivideAddReq struct { // 发起分账请求
	OrderNo                 string `json:"orderNo" v:"required"`                 // 是否必填: Y;长度：32;订单号,保证唯一
	OrderAmount             string `json:"orderAmount" v:"required"`             // 是否必填: Y;长度：10;订单金额,单位：元
	DivideMerchantNo        string `json:"divideMerchantNo" v:"required"`        // 是否必填: Y;长度：20;分账商户
	DivideReceiveMerchantNo string `json:"divideReceiveMerchantNo" v:"required"` // 是否必填: Y;长度：20;分账接收方商户
	DivideMerchantName      string `json:"divideMerchantName" v:"required"`      // 是否必填: Y;长度：100;分账接收方商户名称
	BankAccountNo           string `json:"bankAccountNo,omitempty" v:""`         // 是否必填: N;长度：32;银行账号,若银行账号传值，则银行名称必填
	BankName                string `json:"bankName,omitempty" v:""`              // 是否必填: N;长度：100;银行名称
	BankRemark              string `json:"bankRemark,omitempty" v:""`            // 是否必填: N;长度：20;附言
	NotifyUrl               string `json:"notifyUrl,omitempty" v:""`             // 是否必填: N;长度：100;异步通知地址
}
type DivideAddRes struct { // 发起分账响应
	OrderNo string `json:"orderNo,omitempty" v:""` // 是否必填: N;长度：265;订单号
	Status  int    `json:"status,omitempty" v:""`  // 是否必填: N;长度：265;状态,1.交易成功 2.交易失败 3.交易关闭 4.处理中
}

type DivideQueryReq struct { // 分账查询请求
	OrderNo          string `json:"orderNo" v:"required"`          // 是否必填: Y;长度：265;订单号,保证唯一
	DivideMerchantNo string `json:"divideMerchantNo" v:"required"` // 是否必填: Y;长度：20;分账商户
}
type DivideQueryRes struct { // 分账查询响应
	OrderNo                   int    `json:"orderNo,omitempty" v:""`                   // 是否必填: N;长度：265;订单号
	DivideReceiveMerchantNo   int    `json:"divideReceiveMerchantNo,omitempty" v:""`   // 是否必填: N;长度：265;分账接收方商户
	DivideReceiveMerchantName string `json:"divideReceiveMerchantName,omitempty" v:""` // 是否必填: N;长度：265;分账接收方商户名称
	BankAccountNo             string `json:"bankAccountNo,omitempty" v:""`             // 是否必填: N;长度：265;银行账号
	OrderAmount               string `json:"orderAmount,omitempty" v:""`               // 是否必填: N;长度：265;订单金额,单位：元
	Fee                       string `json:"fee,omitempty" v:""`                       // 是否必填: N;长度：265;手续费,单位：元
	Status                    string `json:"status,omitempty" v:""`                    // 是否必填: N;长度：265;状态,1交易成功 2交易失败 3交易关闭 4处理中 5待确认
	CreateTime                string `json:"createTime,omitempty" v:""`                // 是否必填: N;长度：265;创建时间
	PayTime                   string `json:"payTime,omitempty" v:""`                   // 是否必填: N;长度：265;支付时间
	Remark                    string `json:"remark,omitempty" v:""`                    // 是否必填: N;长度：265;备注
}
