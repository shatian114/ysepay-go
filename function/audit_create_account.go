package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 审核进件资料，返回审核结果
func (s *Client) AuditCreateAccount(ctx context.Context, req *yse_model.AuditCreateAccountReq) (res *yse_model.AuditCreateAccountRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.AuditCreateAccountRes](ctx, req, "t1.smsc.auditCustInfoApply", "1.0")
	if err != nil {
		return
	}
	return
}
