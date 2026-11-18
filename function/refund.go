package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 退款请求
func (s *Client) Refund(ctx context.Context, req *yse_model.RefundReq) (res *yse_model.RefundRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.RefundRes](ctx, req, "unify.trade.refund", "1.0")
	if err != nil {
		return
	}
	return
}

// 退款查询
func (s *Client) RefundQuery(ctx context.Context, req *yse_model.RefundQueryReq) (res *yse_model.RefundQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.RefundQueryRes](ctx, req, "unify.trade.refund.query", "1.1")
	if err != nil {
		return
	}
	return
}
