package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 查询银行支行联行号
func (s *Client) QueryOcrBank(ctx context.Context, req *yse_model.BankOcrReq) (res string, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[string](ctx, req, "pregate.auth.bankOcr", "1.0")
	if err != nil {
		return
	}
	return
}
