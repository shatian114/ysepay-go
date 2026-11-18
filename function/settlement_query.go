package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 查询银行支行联行号
func (s *Client) QuerySettlement(ctx context.Context, req *yse_model.SettlementQueryReq) (res *yse_model.SettlementQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.SettlementQueryRes](ctx, req, "prepose.settlementAccount.query", "1.2")
	if err != nil {
		return
	}
	return
}
