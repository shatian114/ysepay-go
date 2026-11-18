package yse_model

type AlipaySubmitAuthMessageReq struct { // 支付宝实名认证请求
	MercId           string `json:"mercId" v:"required"`             // 是否必填: Y;长度：20;商户号
	MercName         string `json:"mercName,omitempty" v:""`         // 是否必填: N;长度：128;开户意愿确认商户名称
	ContactCorpId    string `json:"contactCorpId" v:"required"`      // 是否必填: Y;长度：20;联系人证件号码
	ContactsTel      string `json:"contactsTel" v:"required"`        // 是否必填: Y;长度：20;联系人证件手机号
	IdHolderType     string `json:"idHolderType,omitempty" v:""`     // 是否必填: N;长度：265;证件持有人类型,类型值：LEGAL-经营者/法人 、SUPER-经办人，示例值：LEGAL。; 注意：;  一、主体类型为政府机关、事业单位时选传：; （1）若上传的是法人证件，则不需要上传该字段。; （2）若因特殊情况，无法提供法人证件时，可上传经办人。; （3）经办人：经商户授权办理支付宝支付业务的人员，授权范围包括但不限于签约，入驻过程需完成账户验证。;  二、主体类型为企业、个体户、社会组织时，默认为经营者/法人，不需要上传该字段。
	IdValidDateBegin string `json:"idValidDateBegin,omitempty" v:""` // 是否必填: N;长度：265;法人证件有效开始日期,格式：yyyy-MM-dd，例：2001-12-31
	IdValidDateEnd   string `json:"idValidDateEnd,omitempty" v:""`   // 是否必填: N;长度：265;法人证件有效结束日期,格式：yyyy-MM-dd，若为永久，此时间为2999-12-31
	BusLincenceBegin string `json:"busLincenceBegin,omitempty" v:""` // 是否必填: N;长度：265;营业执照有效开始日期,格式：yyyy-MM-dd，例：2001-12-31
	BusLincenceEnd   string `json:"busLincenceEnd,omitempty" v:""`   // 是否必填: N;长度：265;营业执照有效结束日期,格式：yyyy-MM-dd，若为永久，此时间为2999-12-31
	ManagementType   string `json:"managementType,omitempty" v:""`   // 是否必填: N;长度：265;小微经营类型,00-门店场所、01-流动经营/便民服务。
	StoreName        string `json:"storeName,omitempty" v:""`        // 是否必填: N;长度：128;门店名称
	ShopPhoto        string `json:"shopPhoto,omitempty" v:""`        // 是否必填: N;长度：200;门头照,要求如下：; 1、经营类型为00，提交门店门口照片，要求招牌清晰可见。; 2、经营类型为01，提交经营/服务现场照片。
	StoreEnvirPhoto  string `json:"storeEnvirPhoto,omitempty" v:""`  // 是否必填: N;长度：200;店内环境照片,要求如下：; 1、经营类型为00，提交店内环境照片。; 2、经营类型为01，可提交另一张经营/服务现场照片。
	LinencePhoto     string `json:"linencePhoto,omitempty" v:""`     // 是否必填: N;长度：265;营业执照营业执照复印件,营业执照营业执照复印件(盖公章)
	IndentityPhoto   string `json:"indentityPhoto,omitempty" v:""`   // 是否必填: N;长度：265;法人身份证复印件,法人身份证复印件（盖私章，签字）/原件扫描件
	ProtocolPhoto    string `json:"protocolPhoto,omitempty" v:""`    // 是否必填: N;长度：265;服务协议复印件,服务协议复印件（盖公章）/原件扫描件
	Cocc             string `json:"cocc,omitempty" v:""`             // 是否必填: N;长度：265;组织机构代码证,组织机构代码证（盖公章）/原件扫描件
	Sfzbackphoto     string `json:"sfzbackphoto,omitempty" v:""`     // 是否必填: N;长度：265;法人身份证反面照
	Sfzfrontphoto    string `json:"sfzfrontphoto,omitempty" v:""`    // 是否必填: N;长度：265;法人身份证正面照
	GrantAuth        string `json:"grantAuth,omitempty" v:""`        // 是否必填: N;长度：265;法定代表人授权函,当证件持有人类型为"经办人"时，必须上传。其他情况，无需上传。
	IdImgHand        string `json:"idImgHand,omitempty" v:""`        // 是否必填: N;长度：265;手持身份证
	BankCardImg      string `json:"bankCardImg,omitempty" v:""`      // 是否必填: N;长度：265;银行卡照片或开户许可证
	BdShopPhoto      string `json:"bdShopPhoto,omitempty" v:""`      // 是否必填: N;长度：265;BD与门头照合影
	BdPosterPhoto    string `json:"bdPosterPhoto,omitempty" v:""`    // 是否必填: N;长度：265;BD与门店摇摇乐海报合影
	WxPlatform       string `json:"wxPlatform,omitempty" v:""`       // 是否必填: N;长度：265;平台入驻照片(支付宝)
	UnitPhoto        string `json:"unitPhoto,omitempty" v:""`        // 是否必填: N;长度：265;单位函
	BookType         string `json:"bookType,omitempty" v:""`         // 是否必填: N;长度：265;证书类型,类型如下：; CERTIFICATE_TYPE_2388-事业单位法人证书; CERTIFICATE_TYPE_2389-统一社会信用代码证书; CERTIFICATE_TYPE_2390-有偿服务许可证（军队医院适用); CERTIFICATE_TYPE_2391-医疗机构执业许可证（军队医院适用); CERTIFICATE_TYPE_2392-企业营业执照（挂靠企业的党组织适用); CERTIFICATE_TYPE_2393-组织机构代码证（政府机关适用); CERTIFICATE_TYPE_2394-社会团体法人登记证书; CERTIFICATE_TYPE_2395-民办非企业单位登记证书; CERTIFICATE_TYPE_2396-基金会法人登记证书; CERTIFICATE_TYPE_2397-慈善组织公开募捐资格证书; CERTIFICATE_TYPE_2398-农民专业合作社法人营业执照; CERTIFICATE_TYPE_2399-宗教活动场所登记证; CERTIFICATE_TYPE_2400-其他证书/批文/证明
	HotLine          string `json:"hotLine,omitempty" v:""`          // 是否必填: N;长度：265;客服电话
	MerContactType   string `json:"merContactType" v:"required"`     // 是否必填: Y;长度：265;联系人类型,65-经营者/法人、66-经办人，示例值：65。; 当商户主体为小微商户时，无需填写,只能选择 65；; 主体为“个体工商户/企业/政府机关/事业单位/社会组织”，可选择 65 或者 66）
	IdTypeCd         string `json:"idTypeCd" v:"required"`           // 是否必填: Y;长度：265;联系人证件类型,00-身份证、04-户口簿、06-外国护照、07-港澳通行证、08-台胞通行证。; 若联系人类型为“法人”，则该姓名需与法人代表证件类型一致。; 若联系人类型为“经办人”，则可填写实际经办人的证件类型。
	IdFrontImg       string `json:"idFrontImg,omitempty" v:""`       // 是否必填: N;长度：265;联系人证件证面照片,联系人为经办人时必填，类型为其他国家或地区居民-护照，此项选填。
	IdBackImg        string `json:"idBackImg,omitempty" v:""`        // 是否必填: N;长度：265;联系人证件反面照片,联系人为经办人时必填，类型为其他国家或地区居民-护照，此项选填。
	BusAutLetterImg  string `json:"busAutLetterImg,omitempty" v:""`  // 是否必填: N;长度：265;业务办理授权函,可不填。
	ContDateBegin    string `json:"contDateBegin,omitempty" v:""`    // 是否必填: N;长度：265;联系人证件有效期开始时间,格式 yyyy-MM-dd
	ContDateEnd      string `json:"contDateEnd,omitempty" v:""`      // 是否必填: N;长度：265;联系人证件有效期结束时间,格式 yyyy-MM-dd，若为永久，此时间为2999-12-31
}

type AlipaySubmitAuthMessageRes struct { // 支付宝实名认证响应
	ReqMsgId           string `json:"reqMsgId,omitempty" v:""`           // 是否必填: N;长度：265;请求流水号
	SendChannelApplyNo string `json:"sendChannelApplyNo,omitempty" v:""` // 是否必填: N;长度：265;实名认证申请单单号,该单号用于查询实名认证结果
}

type AlipaySubmitAuthMessageQueryReq struct { // 支付宝实名认证状态查询请求
	MercId             string `json:"mercId" v:"required"`             // 是否必填: Y;长度：20;商户号
	SendChannelApplyNo string `json:"sendChannelApplyNo" v:"required"` // 是否必填: Y;长度：20;实名认证申请单单号
}

type AlipaySubmitAuthMessageQueryRes struct { // 支付宝实名认证状态查询响应
	ReqMsgId     string `json:"reqMsgId,omitempty" v:""`     // 是否必填: N;长度：265;请求流水号
	ApplyState   string `json:"applyState,omitempty" v:""`   // 是否必填: N;长度：265;申请单状态,类型如下：; AUDITING-审核中；; LEGAL_CONFIRM-待法人确认；; CONTACT_CONFIRM-待确认联系信息；; AUDIT_REJECT-审核失败；; AUDIT_PASS-审核通过；; CANCELED-已撤回；; AUDIT_FREEZE-已冻结；
	QrcodeData   string `json:"qrcodeData,omitempty" v:""`   // 是否必填: N;长度：265;小程序码图片,返回支付宝二维码地址，例如：http://mobilecodec.stable.alipay.net/show.htm?code=0ki0929qqeeid52uvga5
	RejectParam  string `json:"rejectParam,omitempty" v:""`  // 是否必填: N;长度：265;驳回参数,当申请单状态为“审核驳回”时，会返回该字段，标识被驳回的字段名
	RejectReason string `json:"rejectReason,omitempty" v:""` // 是否必填: N;长度：265;驳回原因,当申请单状态为“审核驳回”时，会返回该字段，表示驳回原因
}

type AlipayAuthStatusQueryReq struct { // 支付宝授权状态查询请求
	MercId string `json:"mercId" v:"required"` // 是否必填: Y;长度：20;商户号
}

type AlipayAuthStatusQueryRes struct { // 支付宝授权状态查询响应
	ReqMsgId  string `json:"reqMsgId,omitempty" v:""`  // 是否必填: N;长度：265;请求流水号
	AuthState string `json:"authState,omitempty" v:""` // 是否必填: N;长度：265;授权状态,0-未授权；1-已授权；2-已销户
}

type WechatSubmitAuthMessageReq struct { // 微信实名认证请求
	MercId               string `json:"mercId" v:"required"`                 // 是否必填: Y;长度：265;商户号
	MercName             string `json:"mercName,omitempty" v:""`             // 是否必填: N;长度：265;商户名称,① 商户名称必须与商家营业执照上的名称一致; ②若个体工商户营业执照没有商家名称或者名称为“”或者“**”，则商户名称应填 “个体户XXX”（XXX为营业执照上经营者姓名）; ③若商家为小微商户，商户名称应填 “商户_XXX”（XXX为小微商户身份证上的姓名）
	MerContactType       string `json:"merContactType" v:"required"`         // 是否必填: Y;长度：265;联系人类型,必填项，校验联系人与法人信息是否一致，一致的则为65：经营者/法人 ，不一致时，校验商户类型是否为个体、企业、事业单位、政府机关、社会组织，如不是则报错，如是则为66：经办人; ; 1、主体为“小微/个人卖家”时，可选择：65：经营者/法人。; 2、主体为“个体工商户/企业/政府机关/事业单位/社会组织”，可选择：65：经营者/法人、66：经办人。
	IdTypeCd             string `json:"idTypeCd" v:"required"`               // 是否必填: Y;长度：265;联系人证件类型,证件类型：; 00-身份证; 04-户口薄; 06-外国护照; 07-港澳通行证; 08-台胞通行证; 示例：00
	IdFrontImg           string `json:"idFrontImg,omitempty" v:""`           // 是否必填: N;长度：265;联系人证件正面照片
	IdBackImg            string `json:"idBackImg,omitempty" v:""`            // 是否必填: N;长度：265;联系人证件反面照片
	ContIdValidDateBegin string `json:"contIdValidDateBegin,omitempty" v:""` // 是否必填: N;长度：265;联系人证件有效期开始时间,格式 yyyy-MM-dd
	ContIdValidDateEnd   string `json:"contIdValidDateEnd,omitempty" v:""`   // 是否必填: N;长度：265;联系人证件有效期结束时间,格式 yyyy-MM-dd,若为永久，传2999-12-31
	ContactCorpType      string `json:"contactCorpType,omitempty" v:""`      // 是否必填: N;长度：265;联系人证件类型,只支持身份证 00：身份证；
	ContactCorpId        string `json:"contactCorpId,omitempty" v:""`        // 是否必填: N;长度：265;联系人证件号码
	IndentityPhoto       string `json:"indentityPhoto,omitempty" v:""`       // 是否必填: N;长度：265;法人身份证复印件,（盖私章，签字），原件扫描件
	SfzbackPhoto         string `json:"sfzbackPhoto,omitempty" v:""`         // 是否必填: N;长度：265;法人身份证反面照
	SfzfrontPhoto        string `json:"sfzfrontPhoto,omitempty" v:""`        // 是否必填: N;长度：265;法人身份证正面照
	IdImgHandPhoto       string `json:"idImgHandPhoto,omitempty" v:""`       // 是否必填: N;长度：265;手持身份证
	GrantAuth            string `json:"grantAuth,omitempty" v:""`            // 是否必填: N;长度：265;法定代表人授权函
	IdValidDateBegin     string `json:"idValidDateBegin,omitempty" v:""`     // 是否必填: N;长度：265;法人证件有效开始日期,格式：yyyy-MM-dd
	IdValidDateEnd       string `json:"idValidDateEnd,omitempty" v:""`       // 是否必填: N;长度：265;法人证件有效结束日期,若为永久，此时间为2999-12-31
	IdHolderType         string `json:"idHolderType,omitempty" v:""`         // 是否必填: N;长度：265;法人证件持有人类型,1. 主体类型为政府机关、事业单位时选传：; （1）若上传的是法人证件，则不需要上传该字段。; （2）若因特殊情况，无法提供法人证件时，可上传经办人。; （经办人：经商户授权办理微信支付业务的人员，授权范围包括但不限于签约，入驻过程需完成账户验证）。; 2. 主体类型为企业、个体户、社会组织时，默认为经营者/法人，不需要上传该字段。; 字典值：; LEGAL：经营者/法人; SUPER：经办人
	IdAddress            string `json:"idAddress,omitempty" v:""`            // 是否必填: N;长度：265;法人居住地址,企业商户必填。
	BusLincenceBegin     string `json:"busLincenceBegin,omitempty" v:""`     // 是否必填: N;长度：265;营业执照/证书有效开始日期,格式：yyyy-MM-dd
	BusLincenceEnd       string `json:"busLincenceEnd,omitempty" v:""`       // 是否必填: N;长度：265;营业执照/证书有效结束日期,若为永久，此时间为2999-12-31
	ManagementType       string `json:"managementType,omitempty" v:""`       // 是否必填: N;长度：265;小微经营类型,小微商户必填。; 小微经营类型,00-门店场所、01-流动经营/便民服务。
	StoreName            string `json:"storeName,omitempty" v:""`            // 是否必填: N;长度：265;门店名称,小微商户必填，根据managementType小微经营类型填写不同值。; 小微经营类型为00时，填写门店名称；; 小微经营类型为01时，填写经营/服务名称；; 小微经营类型为02时，填写线上店铺名称；
	ShopPhoto            string `json:"shopPhoto,omitempty" v:""`            // 是否必填: N;长度：265;门头照,报备时未上送则必填。; 根据managementType小微经营类型上传不同照片。; 小微经营类型为00时，提交门店门口照片，要求招牌清晰可见；; 小微经营类型为01时，提交经营/服务现场照片；; 小微经营类型为02时，提交店铺首页截图；
	StoreEnvirPhoto      string `json:"storeEnvirPhoto,omitempty" v:""`      // 是否必填: N;长度：265;店内环境照片,报备时未上送则必填。; 根据managementType小微经营类型上传不同照片。; 小微经营类型为00时，提交店内环境照片；; 小微经营类型为01时，可提交另一张经营/服务现场照片；; 小微经营类型为02时，提交店铺管理后台截图
	BookType             string `json:"bookType,omitempty" v:""`             // 是否必填: N;长度：265;证书类型,证书类型：; CERTIFICATE_TYPE_2388-事业单位法人证书; CERTIFICATE_TYPE_2389-统一社会信用代码证书; CERTIFICATE_TYPE_2390-有偿服务许可证（军队医院适用); CERTIFICATE_TYPE_2391-医疗机构执业许可证（军队医院适用); CERTIFICATE_TYPE_2392-企业营业执照（挂靠企业的党组织适用); CERTIFICATE_TYPE_2393-组织机构代码证（政府机关适用); CERTIFICATE_TYPE_2394-社会团体法人登记证书; CERTIFICATE_TYPE_2395-民办非企业单位登记证书; CERTIFICATE_TYPE_2396-基金会法人登记证书; CERTIFICATE_TYPE_2397-慈善组织公开募捐资格证书; CERTIFICATE_TYPE_2398-农民专业合作社法人营业执照; CERTIFICATE_TYPE_2399-宗教活动场所登记证; CERTIFICATE_TYPE_2400-其他证书/批文/证明; 示例：CERTIFICATE_TYPE_2389
	UnitPhoto            string `json:"unitPhoto,omitempty" v:""`            // 是否必填: N;长度：265;单位函照
	LinencePhoto         string `json:"linencePhoto,omitempty" v:""`         // 是否必填: N;长度：265;营业执照营业执照,需盖公章
	ProtocolPhoto        string `json:"protocolPhoto,omitempty" v:""`        // 是否必填: N;长度：265;服务协议复印件,（盖公章）/原件扫描件
	CoccPhoto            string `json:"coccPhoto,omitempty" v:""`            // 是否必填: N;长度：265;组织机构代码证,（盖公章）/原件扫描件
	BankCardImgPhoto     string `json:"bankCardImgPhoto,omitempty" v:""`     // 是否必填: N;长度：265;银行卡照片或开户许可证
	BdShopPhoto          string `json:"bdShopPhoto,omitempty" v:""`          // 是否必填: N;长度：265;BD与门头照合影,BD与门头照合影
	BdPosterPhoto        string `json:"bdPosterPhoto,omitempty" v:""`        // 是否必填: N;长度：265;BD与门店摇摇乐海报合影,BD与门店摇摇乐海报合影
	WxPlatform           string `json:"wxPlatform,omitempty" v:""`           // 是否必填: N;长度：265;平台入驻照片(微信),平台入驻照片(微信)
	HotLine              string `json:"hotLine,omitempty" v:""`              // 是否必填: N;长度：265;客服电话,客服电话
}

type WechatSubmitAuthMessageRes struct { // 微信实名认证响应
	ReqMsgId           string `json:"reqMsgId,omitempty" v:""`           // 是否必填: N;长度：265;请求流水号, 同请求
	State              string `json:"state,omitempty" v:""`              // 是否必填: N;长度：265;交易状态, 错误码成功时响应 00：成功，99/ZA/ZB/ZC/ZD/ZZ：失败,具体含义见全局中的状态码
	ApplyNo            string `json:"applyNo,omitempty" v:""`            // 是否必填: N;长度：265;申请单编号
	SendChannelApplyNo string `json:"sendChannelApplyNo,omitempty" v:""` // 是否必填: N;长度：265;业务编号
}

type WechatSubmitAuthMsgQueryReq struct { // 微信实名认证状态查询请求
	ApplyNo string `json:"applyNo" v:"required"` // 是否必填: Y;长度：265;申请单单号,或传业务编号（sendChannelApplyNo）
}

type WechatSubmitAuthMsgQueryRes struct { // 微信实名认证状态查询响应
	ReqMsgId    string `json:"reqMsgId,omitempty" v:""`    // 是否必填: N;长度：265;请求流水号
	State       string `json:"state,omitempty" v:""`       // 是否必填: N;长度：265;交易状态, 错误码成功时响应 00：成功，99/ZA/ZB/ZC/ZD/ZZ：失败,具体含义见全局中的状态码
	ApplyState  string `json:"applyState,omitempty" v:""`  // 是否必填: N;长度：265;申请单状态,申请单状态：; APPLYMENT_STATE_WAITTING_FOR_AUDIT-审核中，请耐心等待1~2个工作日，微信支付将会完成审核; APPLYMENT_STATE_EDITTING-编辑中，可能提交申请发生了错误导致，可用同一个业务申请编号重新提交; APPLYMENT_STATE_WAITTING_FOR_CONFIRM_CONTACT-待确认联系信息，请扫描微信支付返回的小程序码确认联系信息(此过程可修改超级管理员手机号) ; APPLYMENT_STATE_WAITTING_FOR_CONFIRM_LEGALPERSON-待账户验证，请扫描微信支付返回的小程序码在小程序端完成账户验证; APPLYMENT_STATE_PASSED-审核通过，请扫描微信支付返回的小程序码在小程序端完成授权流程; APPLYMENT_STATE_REJECTED-审核驳回，请按照驳回原因修改申请资料，并更换业务申请编码，重新提交申请; APPLYMENT_STATE_FREEZED-已冻结，可能是该主体已完成过入驻，请查看驳回原因，并通知驳回原因中指定的联系人扫描微信支付返回的小程序码在小程序端完成授权流程; APPLYMENT_STATE_CANCELED-已作废，表示申请单已被撤销，无需再对其进行操作。; 示例：APPLYMENT_STATE_PASSED
	ChannelCode string `json:"channelCode,omitempty" v:""` // 是否必填: N;长度：265;驳回参数,当申请单状态为“审核驳回”时，会返回该字段，标识被驳回的字段名
	ChannelNote string `json:"channelNote,omitempty" v:""` // 是否必填: N;长度：265;驳回原因,当申请单状态为“审核驳回”时，会返回该字段，表示驳回原因
	QrcodeData  string `json:"qrcodeData,omitempty" v:""`  // 是否必填: N;长度：265;小程序码图片,当申请单状态为APPLYMENT_STATE_WAITTING_FOR_CONFIRM_CONTACT、APPLYMENT_STATE_WAITTING_FOR_CONFIRM_LEGALPERSON、APPLYMENT_STATE_PASSED、APPLYMENT_STATE_FREEZED时，会返回小程序码图片。使用base64解码该字段，可得到图片二进制数据。可用img标签直接加载该图片。
}

type WechatAuthMsgByMercIdReq struct { // 微信实名认证状态详情查询请求
	MercId string `json:"mercId" v:"required"` // 是否必填: Y;长度：265;商户号
}

type WechatAuthMsgByMercIdRes struct { // 微信实名认证状态详情查询响应
	ReqMsgId                 string `json:"reqMsgId,omitempty" v:""`                 // 是否必填: N;长度：265;请求流水号
	State                    string `json:"state,omitempty" v:""`                    // 是否必填: N;长度：265;交易状态, 错误码成功时响应 00：成功，99/ZA/ZB/ZC/ZD/ZZ：失败,具体含义见全局中的状态码
	SendChannelAuthApplyId   string `json:"sendChannelAuthApplyId,omitempty" v:""`   // 是否必填: N;长度：265;实名认证申请单编号,银盛系统生成的申请单编号
	ApplyStatus              string `json:"applyStatus,omitempty" v:""`              // 是否必填: N;长度：265;申请单审核状态,; 申请单审核状态：; 00-审核中; 01-待确认联系信息; 02-待账户验证; 03-审核通过; 04-审核驳回; 05-冻结; 06-已作废; 07-申请单提交失败; 示例：00
	QrcodePic                string `json:"qrcodePic,omitempty" v:""`                // 是否必填: N;长度：265;小程序码图片,fastDfs上的图片地址
	QrcodeData               string `json:"qrcodeData,omitempty" v:""`               // 是否必填: N;长度：265;小程序码图片,base64编码
	OutMercIdAuthStatus      string `json:"outMercIdAuthStatus,omitempty" v:""`      // 是否必填: N;长度：265;授权状态,第三方商户号授权状态
	OutMercIdRemark          string `json:"outMercIdRemark,omitempty" v:""`          // 是否必填: N;长度：265;授权状态备注,第三方商户号授权状态备注
	AuthTime                 string `json:"authTime,omitempty" v:""`                 // 是否必填: N;长度：265;授权认证时间,格式：yyyyMMddHHmmss
	ChannelReturnAuthApplyId string `json:"channelReturnAuthApplyId,omitempty" v:""` // 是否必填: N;长度：265;渠道申请单编号,渠道返回实名认证申请单编号
}

type WechatAuthQueryReq struct { // 微信授权状态查询请求
	MercId string `json:"mercId" v:"required"` // 是否必填: Y;长度：265;商户号
}

type WechatAuthQueryRes struct { // 微信授权状态查询响应
	ReqMsgId     string `json:"reqMsgId" v:"required"`       // 是否必填: Y;长度：265;请求流水号
	State        string `json:"state" v:"required"`          // 是否必填: Y;长度：265;交易状态, 错误码成功时响应 00：成功，99/ZA/ZB/ZC/ZD/ZZ：失败,具体含义见全局中的状态码
	ThirdMercIdX string `json:"thirdMercIdX" v:"required"`   // 是否必填: Y;长度：265;第三方商户号1
	AuthStateX   string `json:"authStateX" v:"required"`     // 是否必填: Y;长度：265;授权状态,AUTHORIZE_STATE_UNAUTHORIZED：未授权; AUTHORIZE_STATE_AUTHORIZED：已授权
	ThirdMercIdY string `json:"thirdMercIdY,omitempty" v:""` // 是否必填: N;长度：265;第三方商户号2,上送商户号可能对应微信的1个或2个第三方商户号，所以当有两个第三方商户的时候thirdMercIdY，authStateY会有值
	AuthStateY   string `json:"authStateY,omitempty" v:""`   // 是否必填: N;长度：265;授权状态2,AUTHORIZE_STATE_UNAUTHORIZED：未授权; AUTHORIZE_STATE_AUTHORIZED：已授权)
}
