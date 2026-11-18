package yse_model

type AreaReq struct { // area查询请求
	ParentCityCd string `json:"parentCityCd,omitempty" v:""` // 是否必填: N;长度：265;父级地区代码,比如 parentCityCd传1200（河北省cityCd），就会分页查询河北省下面的市级
	AreaLevel    string `json:"areaLevel,omitempty" v:""`    // 是否必填: N;长度：265;级别,1-省；2市；3-区 例如areaLevel=1时，返回结果只会有省级的地区码
	PageNumber   string `json:"pageNumber" v:"required"`     // 是否必填: Y;长度：265;当前页码,未上送默认为1
	PageSize     string `json:"pageSize" v:"required"`       // 是否必填: Y;长度：265;每页显示条数,未上送默认为50，最大50条
	CityCd       string `json:"cityCd,omitempty" v:""`       // 是否必填: N;长度：265;地区码
	CityName     string `json:"cityName,omitempty" v:""`     // 是否必填: N;长度：265;地区码名称
}
type AreaRes struct {
	ReqMsgID     string     `json:"reqMsgId"`
	MaxPage      int64      `json:"maxPage"`
	MaxRowCount  int64      `json:"maxRowCount"`
	AreaInfoList []AreaInfo `json:"areaInfoList"`
}

type AreaInfo struct {
	CityCD    string `json:"cityCd"`
	CityNm    string `json:"cityNm"`
	AreaLevel string `json:"areaLevel"`
	IsAble    string `json:"isAble"`
	AreaCode  string `json:"areaCode"`
}
