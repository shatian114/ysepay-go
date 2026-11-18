package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 查询商户余额
func (s *Client) QueryBalance(ctx context.Context, req *yse_model.BalanceQueryReq) (res *yse_model.BalanceQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.BalanceQueryRes](ctx, req, "unify.merchant.balance.query", "1.0")
	if err != nil {
		return
	}
	return
}
