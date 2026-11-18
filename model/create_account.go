package yse_model

type CreateAccountReq struct { // 新增银盛进件请求
	ThirdFlowId string         `json:"thirdFlowId,omitempty"` // 是否必填: N;长度：32;业务方入网申请流水号,唯一（商户状态 查询接口可通过此字段查询）
	CustInfo    CustInfoType   `json:"custInfo"`              // 基本信息
	CrpInfo     CrpInfoType    `json:"crpInfo"`               // 法人信息
	StlAccInfo  StlAccInfoType `json:"stlAccInfo"`            // 结算信息
	BusInfo     BusInfoType    `json:"busInfo"`               // 营业信息
}

type UpdateCreateAccountReq struct { // 修改银盛进件请求
	SysFlowId  string         `json:"sysFlowId,omitempty"` // 是否必填: N;长度：32;业务方入网申请流水号,唯一（商户状态 查询接口可通过此字段查询）
	CustInfo   CustInfoType   `json:"custInfo"`            // 基本信息
	CrpInfo    CrpInfoType    `json:"crpInfo"`             // 法人信息
	StlAccInfo StlAccInfoType `json:"stlAccInfo"`          // 结算信息
	BusInfo    BusInfoType    `json:"busInfo"`             // 营业信息
}

type CustInfoType struct { // 基本信息
	AgtMercId               string `json:"agtMercId,omitempty"`               // 是否必填: N;长度：15;代理商编号,默认同发起方代理商编号，注意：需要上传子商户代理商编号（801开头），否则商户无法交易
	MercName                string `json:"mercName"`                          // 是否必填: Y;长度：100;商户名称,(小微商户名称增加城市名称的校验，有门头（招牌），命名规则为“城市+门头名称”，无门头（招牌），命名规则为“城市 +法人姓名+主营”，其中城市录入规范为，如果是区的，就选地级市名称，如果是县市的选县市名称，也可以支持地级市或者县市名 称。 或者直接全部上送地市级名称)
	MercShortName           string `json:"mercShortName"`                     // 是否必填: Y;长度：20;商户简称,默认为商户名称(小微商户名称增加城市名称的校验，有门头（招牌），命名 规则为“城市+门头名称”，无门头（招牌）， 命名规则为“城市+法人姓名+主营”，其中城市录入规范为，如果是区的，就选地级市名称，如果是县市的选县市名称，也可以支持 地级市或者县市名称。 或者直接全部上送地 市级名称)
	MercType                string `json:"mercType"`                          // 是否必填: Y;长度：2;商户类型,2 小微 3个体 4企业 5 社会组织 6 事业单位 7 政府机关
	MccCd                   string `json:"mccCd"`                             // 是否必填: Y;长度：4;mcc码,通过"查询MCC码"接口获取mccCd，详见MCC码查询
	ContactMail             string `json:"contactMail"`                       // 是否必填: Y;长度：50;联系人邮箱
	ContactMan              string `json:"contactMan,omitempty"`              // 是否必填: N;长度：20;联系人,默认同法人
	ContactPhone            string `json:"contactPhone,omitempty"`            // 是否必填: N;长度：11;联系人电话, 默认同法人电话号码
	CusMgrNm                string `json:"cusMgrNm"`                          // 是否必填: Y;长度：20;客户经理
	IsOpenMarket            string `json:"isOpenMarket,omitempty"`            // 是否必填: N;长度：1;是否开通营销,Y开通N不开通,默认不开通
	NotifyUrl               string `json:"notifyUrl"`                         // 是否必填: Y;长度：100;异步通知地址,银盛通过商户提供的地址下发通知，商户可根据异步通知业务参数bizContent中的notifyType确定通知类型，具体通知内容、场景以及接口要求详见如下:; 客户进件异步通知（notifyType = 1）; 合同签约异步通知（notifyType = 2）; 入网资料变更通知（notifyType = 3）; 费率变更异步通知（notifyType = 5）; 报备异步通知（notifyType = 6）
	Remark                  string `json:"remark,omitempty"`                  // 是否必填: N;长度：200;备注
	SmartScene              string `json:"smartScene,omitempty"`              // 是否必填: N;长度：10;智慧场景,智慧场景，1： 智慧校园 2：智慧食堂 3：智慧商超 4：智慧餐饮 5：智慧酒店 6：智慧医疗 7：智慧菜市场 8：智慧交通 9：智慧景区 10：智慧乡村 11：智慧社区 12：智慧批发市场 13：其他
	Referrer                string `json:"referrer,omitempty"`                // 是否必填: N;长度：50;推荐人
	MerchantUserDefinedInfo string `json:"merchantUserDefinedInfo,omitempty"` // 是否必填: N;长度：1000;商户自定义信息
}

type CrpInfoType struct { // 法人信息
	CrpCertNo   string `json:"crpCertNo"`         // 是否必填: Y;长度：30;法人证件号
	CrpCertType string `json:"crpCertType"`       // 是否必填: Y;长度：2;法人证件类型,00：身份证   06：国外护照 07：港澳通行证 08：台胞通行证
	CertBgn     string `json:"certBgn"`           // 是否必填: Y;长度：8;证件开始日期,证件开始日期yyyyMMdd
	CertExpire  string `json:"certExpire"`        // 是否必填: Y;长度：8;证件有效期,日期格式yyyyMMdd ，如果为长期或者永久，请填值“29991231”
	CrpNati     string `json:"crpNati,omitempty"` // 是否必填: N;长度：3;法人国籍,法人国籍(默认中国(249))
	CrpNm       string `json:"crpNm"`             // 是否必填: Y;长度：20;法人姓名
	CrpPhone    string `json:"crpPhone"`          // 是否必填: Y;长度：11;法人手机号
	CrpAddr     string `json:"crpAddr,omitempty"` // 是否必填: N;长度：100;法人地址
}

type StlAccInfoType struct { // 结算信息
	IsSettInPlatAcc     string `json:"isSettInPlatAcc,omitempty"`     // 是否必填: N;长度：1;是否平台内账户,Y是 N否，默认否
	IsUncrpSett         string `json:"isUncrpSett,omitempty"`         // 是否必填: N;长度：1;是否非同名结算,Y:非同名结算（备用结算账户信息必传） N:同名结算,默认 N:同名结算; ; 同名结算; 1.小微商户：对私且结算账户名为法人姓名; 2.个体商户:; (1). 对私且结算账户名为法人姓名; (2). 对公且结算账户名为营业执照名; 3.企业类商户: 对公结算; ; 非同名结算; 1.小微商户：不支持; 2.个体商户：对私且结算账户名与法人姓名不相同; 3.企业类商户: 对私结算
	StlAccNm            string `json:"stlAccNm,omitempty"`            // 是否必填: N;长度：20;结算户名,默认法人姓名, 非同名结算必填
	StlAccNo            string `json:"stlAccNo"`                      // 是否必填: Y;长度：30;结算账号
	BankSubCode         string `json:"bankSubCode"`                   // 是否必填: Y;长度：20;开户支行联行号,通过"行名行号查询"获取bankCode，详见行名行号查询
	StlAccType          string `json:"stlAccType"`                    // 是否必填: Y;长度：2;结算账户类型,11 对私 21 对公 23 对公存折 24 单位结算卡
	BankMobile          string `json:"bankMobile,omitempty"`          // 是否必填: N;长度：11;银行预留手机号,非同名结算必填
	OpenCertNo          string `json:"openCertNo,omitempty"`          // 是否必填: N;长度：20;开开户证件号,默认法人证件号,非同名结算必填
	BankProince         string `json:"bankProince,omitempty"`         // 是否必填: N;长度：265;银行开户行所属省代码,银行开户行所属市代码都填或都不填 详见地区信息查询
	BankCity            string `json:"bankCity,omitempty"`            // 是否必填: N;长度：265;银行开户行所属市代码,银行开户行所属省代码都填或都不填 详见地区信息查询
	StandByStlAccType   string `json:"standByStlAccType,omitempty"`   // 是否必填: N;长度：2;备用结算账户类型,该字段已下架，无需上送
	StandByStlAccNo     string `json:"standByStlAccNo,omitempty"`     // 是否必填: N;长度：30;备用结算账号,该字段已下架，无需上送
	StandByStlAccNm     string `json:"standByStlAccNm,omitempty"`     // 是否必填: N;长度：20;备用结算户名,该字段已下架，无需上送
	StandByBankSubCode  string `json:"standByBankSubCode,omitempty"`  // 是否必填: N;长度：20;备用结算账户开户行支行编号,该字段已下架，无需上送
	StandByBankProvince string `json:"standByBankProvince,omitempty"` // 是否必填: N;长度：10;备用结算账户银行开户行省代码,该字段已下架，无需上送
	StandByBankCity     string `json:"standByBankCity,omitempty"`     // 是否必填: N;长度：10;备用结算账户银行开户行市代码,该字段已下架，无需上送
}

type BusInfoType struct {
	BusNo          string `json:"busNo,omitempty"`         // 是否必填: N;长度：50;营业执照号,非小微必填
	BusNm          string `json:"busNm"`                   // 是否必填: Y;长度：50;营业执照名称,/商户名称(小微商户名称增加城市名称的校验，有门头（招牌），命名 规则为“城市+门头名称”，无门头（招牌）， 命名规则为“城市+法人姓名+主营”，其中城市录入规范为，如果是区的，就选地级市名称，如果是县市的选县市名称，也可以支持地级市或者县市名称。 或者直接全部上送地市级名称)
	BusCertBgn     string `json:"busCertBgn,omitempty"`    // 是否必填: N;长度：8;营业执照有效开始日期,非小微必填 格式为yyyyMMdd
	BusCertExpire  string `json:"busCertExpire,omitempty"` // 是否必填: N;长度：8;营业执照有效期,非小微必填 格式为yyyyMMdd
	BusProviceCode string `json:"busProviceCode"`          // 是否必填: Y;长度：4;营业归属省代码,通过"地区信息查询"接口获取cityCd，详见地区信息查询
	BusCityCode    string `json:"busCityCode"`             // 是否必填: Y;长度：4;营业归属市代码,通过"地区信息查询"接口获取cityCd，详见地区信息查询
	BusAreaCode    string `json:"busAreaCode,omitempty"`   // 是否必填: N;长度：4;营业归属区(县)代码,通过"地区信息查询"接口获取cityCd，详见地区信息查询，城市下有区(县必填)，城市下无区(县)则不上送，如广东省中山市，广东省东莞市
	BusCertType    string `json:"busCertType,omitempty"`   // 是否必填: N;长度：265;营业执照证件类型,非小微商户必填 ; 个体及企业商户：仅支持 19：社会统一信用代码; 事业单位：仅支持 z2：事业单位法人证书; 政府机关：仅支持 z2：事业单位法人证书，z4:其他证明文件; 社会组织：仅支持 z5:民办非企业单位登记证书，z6:基金会法人登记证书，z4:其他证明文件
	BusAddr        string `json:"busAddr"`                 // 是否必填: Y;长度：265;营业详细地址,需包含省市区，示例：广东省深圳市龙华新区大道513号远景大厦
	RegAddr        string `json:"regAddr,omitempty"`       // 是否必填: N;长度：265;注册地址,默认 营业地址一致
}

type CreateAccountRes struct {
	SysFlowId string `json:"sysFlowId"` //入网申请流水号
}
