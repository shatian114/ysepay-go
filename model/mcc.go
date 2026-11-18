package yse_model

type MccReq struct { // Mcc查询请求
	MccCd    string `json:"mccCd,omitempty" v:""`    // 是否必填: N;长度：265;mcc码,mcc码和商户类型不能全部为空。
	MercType string `json:"mercType,omitempty" v:""` // 是否必填: N;长度：265;商户类型,1-个体工商户 2-企业 3-小微 5-社会组织  6-事业单位 7-政府机关，mcc码和商户类型不能全部为空。
}
type MccRes struct {
	SystemCode  string    `json:"systemCode"`
	MccInfoList []MccInfo `json:"mccInfoList"`
	ReqMsgID    string    `json:"reqMsgId"`
}

type MccInfo struct {
	MccCD      string `json:"mccCd"`
	MercClsNm  string `json:"mercClsNm"`
	MchType    string `json:"mchType"`
	CreateTime string `json:"createTime"`
	MercType   string `json:"mercType"`
	MchTypeNm  string `json:"mchTypeNm"`
}
