package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 查询银行支行联行号
func (s *Client) QueryArea(ctx context.Context, req *yse_model.AreaReq) (res *yse_model.AreaRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.AreaRes](ctx, req, "pregate.trade.findCmmtAreaInfoList", "1.0")
	if err != nil {
		return
	}
	return
}
