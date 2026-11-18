package yse_model

type CommonReq struct { // 请求公共参数
	TimeStamp  string `json:"timeStamp,omitempty"`
	Method     string `json:"method,omitempty"`
	Charset    string `json:"charset,omitempty"`
	Sign       string `json:"sign,omitempty"`
	Check      string `json:"check,omitempty"`
	BizContent string `json:"bizContent,omitempty"`
	ReqID      string `json:"reqId,omitempty"`
	CertID     string `json:"certId,omitempty"`
	Version    string `json:"version,omitempty"`
}

type CommonRes struct { // 请求统一响应
	Code         string  `json:"code" v:""`                   // 是否必填: Y;长度：5;网关响应码，示例值：00000 ，详见网关公共响应码
	Msg          string  `json:"msg" v:""`                    // 是否必填: Y;长度：50;网关响应码描述
	SubCode      string  `json:"subCode" v:""`                // 是否必填: Y;长度：265;业务响应码，参见具体的API接口文档
	SubMsg       string  `json:"subMsg" v:""`                 // 是否必填: Y;长度：265;业务响应描述
	TimeStamp    string  `json:"timeStamp" v:""`              // 是否必填: Y;长度：265;响应时间，格式"yyyy-MM-dd HH:mm:ss"
	Norce        string  `json:"norce" v:""`                  // 是否必填: Y;长度：128;随机参数
	Sign         string  `json:"sign" v:""`                   // 是否必填: Y;长度：265;响应参数的签名串，详见demo,; 注意:请用银盛公钥进行验签
	BusinessData *string `json:"businessData,omitempty" v:""` // 是否必填: Y;长度：265;业务响应参数集合，; 注意：银盛网关通过AES加密业务响应参数集合，得到businessData，商户需要对其进行解密，详情请见demo
}

func (s *CommonRes) IsErr() bool {
	return s.Code != "00000"
}

func (s *CommonRes) IsSubErr() bool {
	return s.SubCode != "0000"
}
