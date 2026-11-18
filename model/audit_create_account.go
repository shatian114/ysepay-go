package yse_model

type AuditCreateAccountReq struct { // 对进件资料进行初审请求
	AuditFlag string `json:"auditFlag,omitempty" v:""` // 审核标志,Y通过,N拒绝
	SysFlowId string `json:"sysFlowId,omitempty" v:""` // 入网申请流水号
}
type AuditCreateAccountRes struct { // 对进件资料进行初审结果
	CustId    string `json:"custId,omitempty" v:""`    // 客户号,入网成功后返回
	SysFlowId string `json:"sysFlowId,omitempty" v:""` // 入网申请流水号
	Status    string `json:"status,omitempty" v:""`    // 入网状态,00:入网成功 01:待上传图片 02:待审核 90:审核拒绝 10:转人工审核
}
