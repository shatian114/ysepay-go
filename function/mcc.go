package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 查询银行支行联行号
func (s *Client) QueryMcc(ctx context.Context, req *yse_model.MccReq) (res *yse_model.MccRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.MccRes](ctx, req, "aggregation.scan.mccQuery", "1.0")
	if err != nil {
		return
	}
	return
}
