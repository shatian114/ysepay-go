package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 发起提现
func (s *Client) Withdraw(ctx context.Context, req *yse_model.WithdrawReq) (res *yse_model.WithdrawRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.WithdrawRes](ctx, req, "unify.basePay.withdraw.quickAcccept", "1.0")
	if err != nil {
		return
	}
	return
}

// 查询提现结果
func (s *Client) WithdrawQuery(ctx context.Context, req *yse_model.WithdrawQueryReq) (res *yse_model.WithdrawQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.WithdrawQueryRes](ctx, req, "unify.merchant.withdraw.quick.query", "1.0")
	if err != nil {
		return
	}
	return
}
