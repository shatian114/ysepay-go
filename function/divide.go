package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 开通分账
func (s *Client) DivideOpen(ctx context.Context, req *yse_model.DivideReceiveConfigReq) (res *yse_model.SharingOpenRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.SharingOpenRes](ctx, req, "smsc.revenueSharing.open", "1.5")
	if err != nil {
		return
	}
	return
}

// 配置分账接收方
func (s *Client) DivideReceiveConfig(ctx context.Context, req *yse_model.SharingOpenReq) (res any, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[any](ctx, req, "eqt.divideReceiveMerchant.config", "1.0")
	if err != nil {
		return
	}
	return
}

// 发起分账
func (s *Client) DivideAdd(ctx context.Context, req *yse_model.DivideAddReq) (res *yse_model.DivideAddRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.DivideAddRes](ctx, req, "eqt.divide.single", "1.0")
	if err != nil {
		return
	}
	return
}

// 查询分账
func (s *Client) DivideQuery(ctx context.Context, req *yse_model.DivideQueryReq) (res *yse_model.DivideQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.DivideQueryRes](ctx, req, "eqt.divide.query", "1.0")
	if err != nil {
		return
	}
	return
}
