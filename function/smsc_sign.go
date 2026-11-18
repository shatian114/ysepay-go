package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 查询银行支行联行号
func (s *Client) SismSign(ctx context.Context, req *yse_model.SmscSignReq) (res *yse_model.SmscSignRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.SmscSignRes](ctx, req, "t1.smsc.sign", "1.6")
	if err != nil {
		return
	}
	return
}
