package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 查询银行支行联行号
func (s *Client) QueryBank(ctx context.Context, req *yse_model.BankReq) (res *yse_model.BankRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.BankRes](ctx, req, "pregate.trade.findBankNameAndBankCode", "1.0")
	if err != nil {
		return
	}
	return
}
