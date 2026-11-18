package yse_model

type ReportAppidAddReq struct { // 报备新增appid请求
	ChannelId         string `json:"channelId" v:"required"`           // 是否必填: Y;长度：265;渠道号
	MercId            string `json:"mercId" v:"required"`              // 是否必填: Y;长度：265;商户号,需配置APPID的商户号
	JsPayRelatedAppId string `json:"jsPayRelatedAppId,omitempty" v:""` // 是否必填: N;长度：265;公众号appid,公众号appid、小程序appid、app支付appid不能同时为空
	AppletId          string `json:"appletId,omitempty" v:""`          // 是否必填: N;长度：265;小程序appid,公众号appid、小程序appid、app支付appid不能同时为空
	AppPayAppId       string `json:"appPayAppId,omitempty" v:""`       // 是否必填: N;长度：265;app支付appid,公众号appid、小程序appid、app支付appid不能同时为空
}

type ReportAppidAddRes struct { // 报备新增appid响应
	State    string `json:"state,omitempty" v:""`    // 是否必填: N;长度：265;交易状态 错误码成功时响应
	ReqMsgId string `json:"reqMsgId,omitempty" v:""` // 是否必填: N;长度：265;请求流水号
}

type ReportQueryReq struct { // 查询报备请求
	MercId    string `json:"mercId" v:"required"`    // 是否必填: Y;长度：265;商户号
	ChannelId string `json:"channelId" v:"required"` // 是否必填: Y;长度：265;交易报备渠道编号,CUPS_WECHAT-银联微信；CUPS_ALIPAY-银联支付宝；NUCC_WECHAT-网联微信；NUCC_ALIPAY-网联支付宝
}

type ReportQueryRes struct { // 查询报备响应
	ReqMsgId       string `json:"reqMsgId,omitempty" v:""`       // 是否必填: N;长度：265;请求流水号
	State          string `json:"state,omitempty" v:""`          // 是否必填: N;长度：265;交易状态
	MercId         string `json:"mercId,omitempty" v:""`         // 是否必填: N;长度：265;商户号
	MercType       string `json:"mercType,omitempty" v:""`       // 是否必填: N;长度：265;商户信息-商户类别,0或null：普通商户，1：小微商户,2：个体商户,3：事业单位，4：社会组织（原其他组织）， 5：政府机关
	OpenId         string `json:"openId,omitempty" v:""`         // 是否必填: N;长度：265;微信openid
	Remark         string `json:"remark,omitempty" v:""`         // 是否必填: N;长度：265;备注
	CreatDate      string `json:"creatDate,omitempty" v:""`      // 是否必填: N;长度：265;录入时间
	AppIdType      string `json:"appIdType,omitempty" v:""`      // 是否必填: N;长度：265;公众号类型
	ReportDate     string `json:"reportDate,omitempty" v:""`     // 是否必填: N;长度：265;报备时间
	AppId          string `json:"appId,omitempty" v:""`          // 是否必填: N;长度：265;公众号
	BusiType       string `json:"busiType,omitempty" v:""`       // 是否必填: N;长度：265;业务类型,0-微信 1-支付宝 2-银盛通 4-QQ钱包
	ChannelId      string `json:"channelId,omitempty" v:""`      // 是否必填: N;长度：265;交易报备渠道编号
	ThirdMercId    string `json:"thirdMercId,omitempty" v:""`    // 是否必填: N;长度：265;第三方商户号
	ThirdMercIdTwo string `json:"thirdMercIdTwo,omitempty" v:""` // 是否必填: N;长度：265;第三方商户号2,银联或网联通道时，表示收单机构的商户号. 银联支付宝的appid
	AppidType1     string `json:"appidType1,omitempty" v:""`     // 是否必填: N;长度：265;appid1类型
	AppidType2     string `json:"appidType2,omitempty" v:""`     // 是否必填: N;长度：265;appid2类型
	AppidType3     string `json:"appidType3,omitempty" v:""`     // 是否必填: N;长度：265;appid3类型
	AppidType4     string `json:"appidType4,omitempty" v:""`     // 是否必填: N;长度：265;appid4类型
	AppId1         string `json:"appId1,omitempty" v:""`         // 是否必填: N;长度：265;公众号1
	AppId2         string `json:"appId2,omitempty" v:""`         // 是否必填: N;长度：265;公众号2
	AppId3         string `json:"appId3,omitempty" v:""`         // 是否必填: N;长度：265;公众号3
	AppId4         string `json:"appId4,omitempty" v:""`         // 是否必填: N;长度：265;公众号4
	MchType        string `json:"mchType,omitempty" v:""`        // 是否必填: N;长度：265;报备行业类型
	ApprSts        string `json:"apprSts,omitempty" v:""`        // 是否必填: N;长度：265;报备状态,00-已注册（交易开通）01-待审核 02-审核中 03-注册失败 04-交易关闭
	JsapiPath1     string `json:"jsapiPath1,omitempty" v:""`     // 是否必填: N;长度：265;公众号授权目录1
	JsapiPath2     string `json:"jsapiPath2,omitempty" v:""`     // 是否必填: N;长度：265;公众号授权目录2
	JsapiPath3     string `json:"jsapiPath3,omitempty" v:""`     // 是否必填: N;长度：265;公众号授权目录3
	JsapiPath4     string `json:"jsapiPath4,omitempty" v:""`     // 是否必填: N;长度：265;公众号授权目录4
	JsapiPath5     string `json:"jsapiPath5,omitempty" v:""`     // 是否必填: N;长度：265;公众号授权目录5
}

type ReportAddReq struct { // 新增报备请求
	ChannelCode             string `json:"channelCode" v:"required"`               // 是否必填: Y;长度：20;交易报备渠道编号,类型如下：; CUPS_WECHAT-银联微信; CUPS_ALIPAY-银联支付宝; NUCC_WECHAT-网联微信; NUCC_ALIPAY-网联支付宝
	MercId                  string `json:"mercId" v:"required"`                    // 是否必填: Y;长度：20;商户信息-商户号
	LinencePhoto            string `json:"linencePhoto,omitempty" v:""`            // 是否必填: N;长度：100;商户信息-营业执照营业执照复印件(盖公章),; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 小微实名认证必填；微信新绿洲计划报备必填；蓝海行动报名必填
	IndentityPhoto          string `json:"indentityPhoto,omitempty" v:""`          // 是否必填: N;长度：100;商户信息-法人身份证复印件（盖私章，签字）/原件扫描件,商户进件时未传，法人身份证复印件（盖私章，签字）/原件扫描件必传；; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg
	ProtocolPhoto           string `json:"protocolPhoto,omitempty" v:""`           // 是否必填: N;长度：100;商户信息-服务协议复印件（盖公章）/原件扫描件,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg
	Cocc                    string `json:"cocc,omitempty" v:""`                    // 是否必填: N;长度：100;商户信息-组织机构代码证（盖公章）/原件扫描件,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 蓝海行动报名必填
	SfzFrontPhoto           string `json:"sfzFrontPhoto,omitempty" v:""`           // 是否必填: N;长度：100;商户信息-法人身份证正面照,商户进件时未传，法人身份证正面照必传；; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg
	SfzBackPhoto            string `json:"sfzBackPhoto,omitempty" v:""`            // 是否必填: N;长度：100;商户信息-法人身份证反面照,商户进件时未传，法人身份证反面照必传；; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg
	IdcarNo                 string `json:"idcarNo,omitempty" v:""`                 // 是否必填: N;长度：50;商户信息法人身份证号码
	MccSubCd                string `json:"mccSubCd" v:"required"`                  // 是否必填: Y;长度：265;商户信息-银盛商户类型,传商户类型编码，详见"银盛商户类型及mcc码对照表"。
	BusiType                string `json:"busiType" v:"required"`                  // 是否必填: Y;长度：1;业务类型,类型如下：; 0-微信 ; 1-支付宝
	MercShortName           string `json:"mercShortName" v:"required"`             // 是否必填: Y;长度：50;商户信息-商户简称
	MercProv                string `json:"mercProv" v:"required"`                  // 是否必填: Y;长度：6;商户信息-商户归属省,所属省编码，例:5800。通过"地区信息查询"接口获取cityCd，详见地区信息查询
	MercCity                string `json:"mercCity" v:"required"`                  // 是否必填: Y;长度：6;商户信息-商户归属市,所属市编， 例:5840。通过"地区信息查询"接口获取cityCd，详见地区信息查询
	MercArea                string `json:"mercArea" v:"required"`                  // 是否必填: Y;长度：6;商户信息-商户归属区,所属区编码，例:5840。通过"地区信息查询"接口获取cityCd，详见地区信息查询
	BusAddr                 string `json:"busAddr" v:"required"`                   // 是否必填: Y;长度：255;商户信息-营业地址（联系地址）
	ContactsName            string `json:"contactsName" v:"required"`              // 是否必填: Y;长度：50;商户信息-商户联系人名称
	ContactsTel             string `json:"contactsTel" v:"required"`               // 是否必填: Y;长度：11;商户信息-商户联系人手机
	ContactsEmail           string `json:"contactsEmail" v:"required"`             // 是否必填: Y;长度：50;商户信息-商户联系人邮箱
	CrpNm                   string `json:"crpNm" v:"required"`                     // 是否必填: Y;长度：50;商户信息-法人名称
	OrgNo                   string `json:"orgNo" v:"required"`                     // 是否必填: Y;长度：20;商户信息-商户归属机构
	CertNo                  string `json:"certNo" v:"required"`                    // 是否必填: Y;长度：50;商户信息-法人证件号码
	AgentName               string `json:"agentName,omitempty" v:""`               // 是否必填: N;长度：50;商户信息-代理商名称
	AgtMercId               string `json:"agtMercId,omitempty" v:""`               // 是否必填: N;长度：20;商户信息-代理商编号
	BankType                string `json:"bankType" v:"required"`                  // 是否必填: Y;长度：20;商户信息-开户行行别,通过"行名行号查询"获取bankCode，详见行名行号查询
	BankName                string `json:"bankName" v:"required"`                  // 是否必填: Y;长度：50;商户信息-开户行名称
	AccountType             string `json:"accountType" v:"required"`               // 是否必填: Y;长度：2;商户信息-账户类型,11-个人借记卡 12个人贷记卡 13个人存折 14对私借记卡（二类户） 21对公借记卡 22对公贷记卡 23对公存折 24单位结算卡 31有卡bin预付费卡 32无卡bin预付费卡 00平台账户 **所有卡类型
	AccountName             string `json:"accountName" v:"required"`               // 是否必填: Y;长度：50;商户信息-账户名称
	AccountNo               string `json:"accountNo" v:"required"`                 // 是否必填: Y;长度：32;商户信息-账户账号
	CustId                  string `json:"custId" v:"required"`                    // 是否必填: Y;长度：20;商户信息-客户号
	MchType                 string `json:"mchType" v:"required"`                   // 是否必填: Y;长度：3;商户信息-商户类型,传编码，详见"商户类型"。
	MercName                string `json:"mercName,omitempty" v:""`                // 是否必填: N;长度：50;商户信息-商户名称
	ReportName              string `json:"reportName,omitempty" v:""`              // 是否必填: N;长度：265;商户信息-微信报备名称
	MercType                string `json:"mercType" v:"required"`                  // 是否必填: Y;长度：1;商户信息-商户类别,类别如下：; 0或null-普通企业商户; 1-小微商户; 2-个体商户; 3-事业单位; 4-社会组织（原其他组织）; 5-政府机关
	CertType                string `json:"certType,omitempty" v:""`                // 是否必填: N;长度：2;商户信息-法人证件类型,类型如下：; 00-身份证; 01-中国护照; 02-军人身份证; 03-警官证; 04-户口簿; 05-临时身份证; 06-外国护照; 07-外国护照; 08-台胞通行证; 09-离休干部荣誉证书; 10-军官退休证书; 11-文职干部退休证书; 12-军事院校哦学员证书; 13-武装警察身份证; 14-军官证; 15-文职干部证书; 16-军人士兵证书; 17-武警士兵证书; 18-其他证件
	PId                     string `json:"pId,omitempty" v:""`                     // 是否必填: N;长度：265;商户信息-支付宝推荐关系pid
	ShopPhoto               string `json:"shopPhoto,omitempty" v:""`               // 是否必填: N;长度：100;商户信息-门头照,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 小微实名认证必填；微信新绿洲计划报备必填；蓝海行动报名必填
	StoreEnvirPhoto         string `json:"storeEnvirPhoto,omitempty" v:""`         // 是否必填: N;长度：100;商户信息-店内环境照,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 小微实名认证必填；微信新绿洲计划报备必填；蓝海行动报名必填
	IdImgHand               string `json:"idImgHand,omitempty" v:""`               // 是否必填: N;长度：100;商户信息-手持身份证照,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 小微实名认证必填
	BankCardImg             string `json:"bankCardImg,omitempty" v:""`             // 是否必填: N;长度：100;商户信息-银行卡照片或开户许可证,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg
	BdShopPhoto             string `json:"bdShopPhoto,omitempty" v:""`             // 是否必填: N;长度：100;商户信息-客户经理与门头照合影,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 微信新绿洲计划报备必填
	BdPosterPhoto           string `json:"bdPosterPhoto,omitempty" v:""`           // 是否必填: N;长度：100;商户信息-客户经理与门店摇摇乐海报合影,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 微信新绿洲计划报备必填
	WxPlatform              string `json:"wxPlatform,omitempty" v:""`              // 是否必填: N;长度：100;商户信息-平台入驻照片(微信),; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 微信新绿洲计划报备必填
	IdValidDateBegin        string `json:"idValidDateBegin,omitempty" v:""`        // 是否必填: N;长度：8;商户信息-法人证件有效开始期,格式：yyyyMMdd
	IdValidDateEnd          string `json:"idValidDateEnd,omitempty" v:""`          // 是否必填: N;长度：8;商户信息-法人证件有效结束期,格式：yyyyMMdd，如果为永久值，则为29991231。
	BusLincenceBegin        string `json:"busLincenceBegin,omitempty" v:""`        // 是否必填: N;长度：8;商户信息-营业执照有效开始期,格式：yyyyMMdd
	BusLincenceEnd          string `json:"busLincenceEnd,omitempty" v:""`          // 是否必填: N;长度：8;商户信息-营业执照失效期,格式：yyyyMMdd，如果为永久值，则为29991231。
	ManagementType          string `json:"managementType,omitempty" v:""`          // 是否必填: N;长度：265;商户信息-小微商户经营类型,类型如下：; 00-门店场所; 01-流动经营/便民服务; 02-线上商品/服务交易
	StoreName               string `json:"storeName,omitempty" v:""`               // 是否必填: N;长度：265;商户信息-门店名称
	WechatCheckStandPhoto   string `json:"wechatCheckStandPhoto,omitempty" v:""`   // 是否必填: N;长度：100;商户信息-微信收银台照片,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 微信新绿洲计划报备必
	AlipayCheckStandPhoto   string `json:"alipayCheckStandPhoto,omitempty" v:""`   // 是否必填: N;长度：100;商户信息-支付宝收银台主扫照片,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg; 蓝海行动报名必填
	AliPayStoreCashierPhoto string `json:"aliPayStoreCashierPhoto,omitempty" v:""` // 是否必填: N;长度：100;商户信息-支付宝收银台被扫照片,; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg
	BookType                string `json:"bookType,omitempty" v:""`                // 是否必填: N;长度：22;商户信息-证书类型,类型如下：; CERTIFICATE_TYPE_2388：事业单位法人证书; CERTIFICATE_TYPE_2389：统一社会信用代码证书; CERTIFICATE_TYPE_2390：有偿服务许可证（军队医院适用); CERTIFICATE_TYPE_2391：医疗机构执业许可证（军队医院适用); CERTIFICATE_TYPE_2392：企业营业执照（挂靠企业的党组织适用); CERTIFICATE_TYPE_2393：组织机构代码证（政府机关适用); CERTIFICATE_TYPE_2394：社会团体法人登记证书; CERTIFICATE_TYPE_2395：民办非企业单位登记证书; CERTIFICATE_TYPE_2396：基金会法人登记证书; CERTIFICATE_TYPE_2397：慈善组织公开募捐资格证书; CERTIFICATE_TYPE_2398：农民专业合作社法人营业执照; CERTIFICATE_TYPE_2399：宗教活动场所登记证; CERTIFICATE_TYPE_2400：其他证书/批文/证明
	UnitPhoto               string `json:"unitPhoto,omitempty" v:""`               // 是否必填: N;长度：100;商户信息-单位证明函,商户类别为事业单位时，单位证明函必传; 当picType=1时，传图片ID(调用图片上传接口得到图片ID)；; 当picType=2时，传url地址，例如：group1,M00/26/0D/CtUiIlYg-AcTmgAAU2krJ8NgE60.jpeg
	HotLine                 string `json:"hotLine,omitempty" v:""`                 // 是否必填: N;长度：20;商户信息-客服电话
	BusinessLicenseType     string `json:"businessLicenseType" v:"required"`       // 是否必填: Y;长度：1;商户信息-证照类型,类型如下：; 0-营业执照; 1-事业单位法人证书; 2-小微商户负责人身份证; 3-社会信用等级证书
	BusinessLicense         string `json:"businessLicense" v:"required"`           // 是否必填: Y;长度：50;商户信息-证件编号
	AppletAppId             string `json:"appletAppId,omitempty" v:""`             // 是否必填: N;长度：265;appid信息-小程序APPID
	AppId1                  string `json:"appId1,omitempty" v:""`                  // 是否必填: N;长度：265;appid信息-公众号1
	AppId2                  string `json:"appId2,omitempty" v:""`                  // 是否必填: N;长度：265;appid信息-公众号2
	AppId3                  string `json:"appId3,omitempty" v:""`                  // 是否必填: N;长度：265;appid信息-公众号3
	AppId4                  string `json:"appId4,omitempty" v:""`                  // 是否必填: N;长度：265;appid信息-公众号4
	ApplyServices           string `json:"applyServices,omitempty" v:""`           // 是否必填: N;长度：2;商户信息-申请的服务,如有多个，必须用英文分号“;”隔开。; 类型如下：; 01-正扫&反扫; 02-JS支付 (包括公众号/生活号,A/T的小程序支付)。
	PicType                 string `json:"picType,omitempty" v:""`                 // 是否必填: N;长度：1;上传图片类型,类型如下：; 1-图片id; 2-图片url; 可为空，1或空时为图片id，识别url时必须传2
	JsapiPath1              string `json:"jsapiPath1,omitempty" v:""`              // 是否必填: N;长度：265;公众号授权目录1
	JsapiPath2              string `json:"jsapiPath2,omitempty" v:""`              // 是否必填: N;长度：265;公众号授权目录2
	JsapiPath3              string `json:"jsapiPath3,omitempty" v:""`              // 是否必填: N;长度：265;公众号授权目录3
	JsapiPath4              string `json:"jsapiPath4,omitempty" v:""`              // 是否必填: N;长度：265;公众号授权目录4
	JsapiPath5              string `json:"jsapiPath5,omitempty" v:""`              // 是否必填: N;长度：265;公众号授权目录5
}

type ReportAddRes struct { // 新增报备响应
	ReqMsgId         string       `json:"reqMsgId,omitempty" v:""`         // 是否必填: N;长度：64;请求流水号
	State            string       `json:"state,omitempty" v:""`            // 是否必填: N;长度：2;交易状态 错误码成功时响应
	MercId           string       `json:"mercId,omitempty" v:""`           // 是否必填: N;长度：20;商户号
	MchFlag          string       `json:"mchFlag,omitempty" v:""`          // 是否必填: N;长度：1;商户标识,0.线下 1.线上
	BusiType         string       `json:"busiType,omitempty" v:""`         // 是否必填: N;长度：1;业务类型,0-微信 1-支付宝 2-银盛通 4-QQ钱包
	ChannelId        string       `json:"channelId,omitempty" v:""`        // 是否必填: N;长度：20;交易报备渠道编号
	BakModel         string       `json:"bakModel,omitempty" v:""`         // 是否必填: N;长度：1;报备模式,0:一户一报 1:大商户模式 默认0
	ThirdMercId      string       `json:"thirdMercId,omitempty" v:""`      // 是否必填: N;长度：32;三方商户号
	ThirdMercIdTwo   string       `json:"thirdMercIdTwo,omitempty" v:""`   // 是否必填: N;长度：32;第三方商户号
	ApprSts          string       `json:"apprSts,omitempty" v:""`          // 是否必填: N;长度：2;报备状态,00-已注册（交易开通）01-待审核 02-审核中 03-注册失败 04-交易关闭 99-失败
	Remark           string       `json:"remark,omitempty" v:""`           // 是否必填: N;长度：255;备注
	Key              string       `json:"key,omitempty" v:""`              // 是否必填: N;长度：265;秘钥
	OpenId           string       `json:"openId,omitempty" v:""`           // 是否必填: N;长度：265;id号
	CreatDate        string       `json:"creatDate,omitempty" v:""`        // 是否必填: N;长度：265;录入时间
	ReportDate       string       `json:"reportDate,omitempty" v:""`       // 是否必填: N;长度：265;报备时间
	MercType         string       `json:"mercType,omitempty" v:""`         // 是否必填: N;长度：1;商户信息-商户类别,0或null：普通企业商户，1：小微商户,2：个体商户,3：事业单位，4：社会组织（原其他组织）， 5：政府机关
	MchType          string       `json:"mchType,omitempty" v:""`          // 是否必填: N;长度：265;商户信息-报备行业类型
	AppId            string       `json:"appId,omitempty" v:""`            // 是否必填: N;长度：265;公众号
	AppletAppId      string       `json:"appletAppId,omitempty" v:""`      // 是否必填: N;长度：265;程序公众号
	ThirdMercIdThree string       `json:"thirdMercIdThree,omitempty" v:""` // 是否必填: N;长度：265;微信一级行业类目
	Level            string       `json:"level,omitempty" v:""`            // 是否必填: N;长度：265;支付宝商户等级
	TerminalInfoList TerminalInfo `json:"terminalInfoList,omitempty" v:""` // 是否必填: N;长度：265;终端报备结果列表
	AppId1Flag       string       `json:"appId1Flag,omitempty" v:""`       // 是否必填: N;长度：265;公众号1新增结果,true-成功；false-失败
	AppId2Flag       string       `json:"appId2Flag,omitempty" v:""`       // 是否必填: N;长度：265;公众号2新增结果,true-成功；false-失败
	AppId3Flag       string       `json:"appId3Flag,omitempty" v:""`       // 是否必填: N;长度：265;公众号3新增结果,true-成功；false-失败
	AppId4Flag       string       `json:"appId4Flag,omitempty" v:""`       // 是否必填: N;长度：265;公众号4新增结果,true-成功；false-失败
}

type TerminalInfo struct { // 新增报备响应终端信息
	MercId              string `json:"mercId,omitempty" v:""`              // 是否必填: N;长度：265;商户号
	ChannelCode         string `json:"channelCode,omitempty" v:""`         // 是否必填: N;长度：265;终端报备渠道编号
	TerminalType        string `json:"terminalType,omitempty" v:""`        // 是否必填: N;长度：265;终端类型,01 自动柜员机（含 ATM 和 CDM）和多媒体自助终端、02 传统 POS、03 mPOS、04 智能 POS、05 II 型固定电话、 06 云闪付终端、07 保留使用、08 手机 POS、09 刷脸付终端、10 条码支付受理终端、11 辅助受理终端、12 行业终端（公交、地铁用于指定行业的终端）、13 MIS 终端
	TerminalNo          string `json:"terminalNo,omitempty" v:""`          // 是否必填: N;长度：265;终端编号
	SerialNum           string `json:"serialNum,omitempty" v:""`           // 是否必填: N;长度：265;终端序列号
	TerminalAddr        string `json:"terminalAddr,omitempty" v:""`        // 是否必填: N;长度：265;终端布放地址
	TerminalReportState string `json:"terminalReportState,omitempty" v:""` // 是否必填: N;长度：265;终端报备状态,00-成功；99-失败
	Note                string `json:"note,omitempty" v:""`                // 是否必填: N;长度：265;终端报备说明
}

type MercReportQueryReq struct { // 商户报备结果查询请求
	MercId     string `json:"mercId" v:"required"`       // 是否必填: Y;长度：265;银盛商户号
	TerminalNo string `json:"terminalNo,omitempty" v:""` // 是否必填: N;长度：265;终端号
}

type MercReportQueryRes struct { // 商户报备结果查询响应
	MercId     string                   `json:"mercId,omitempty" v:""`     // 是否必填: N;长度：265;银盛商户号
	TerminalNo string                   `json:"terminalNo,omitempty" v:""` // 是否必填: N;长度：265;终端号
	List       []MercReportQueryResData `json:"list,omitempty" v:""`       // 是否必填: N;长度：265;结果集合
}

type MercReportQueryResData struct { // 商户报备结果查询响应数据
	ThirdMercId string `json:"thirdMercId,omitempty" v:""` // 是否必填: N;长度：265;第三方商户号
	ApprSts     string `json:"apprSts,omitempty" v:""`     // 是否必填: N;长度：265;报备状态,01：报备中；00：报备成功；02：报备失败；03：冻结；04：注销；05：关闭；
	ChannelId   string `json:"channelId,omitempty" v:""`   // 是否必填: N;长度：265;报备渠道,银联（银联刷卡）:0001；万事网联:0002；连通:0003；讯联:0004；银联微信：0005；网联微信：0006; 银联支付宝：0007； 网联支付宝：0008 ;网联互联互通：0009；
	Remark      string `json:"remark,omitempty" v:""`      // 是否必填: N;长度：265;报备结果描述
}
