package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 查询进件资料的审核结果
func (s *Client) QueryCreateAccountStatus(ctx context.Context, req *yse_model.QueryCreateAccountStatusReq) (res *yse_model.QueryCreateAccountStatusRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.QueryCreateAccountStatusRes](ctx, req, "smsc.queryCustApply", "1.1")
	if err != nil {
		return
	}
	return
}
