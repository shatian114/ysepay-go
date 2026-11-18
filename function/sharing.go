package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 上传开通分账的图片
func (s *Client) SharingUploadPic(ctx context.Context, req *yse_model.SharingUploadPicReq) (res *yse_model.SharingUploadPicRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.SharingUploadPicRes](ctx, req, "smsc.revenueSharing.upload.pic", "1.1")
	if err != nil {
		return
	}
	return
}

// 开通订单分账
func (s *Client) SharingOpen(ctx context.Context, req *yse_model.OpenSharingReq) (res *yse_model.OpenSharingRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.OpenSharingRes](ctx, req, "smsc.revenueSharing.open", "1.7")
	if err != nil {
		return
	}
	return
}

// 开通订单分账进度查询
func (s *Client) SharingOpenQuery(ctx context.Context, req *yse_model.OpenSharingQueryReq) (res *yse_model.OpenSharingQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.OpenSharingQueryRes](ctx, req, "smsc.revenueSharing.progress.query", "1.2")
	if err != nil {
		return
	}
	return
}

// 关闭订单分账
func (s *Client) SharingClose(ctx context.Context, req *yse_model.CloseSharingReq) (commonRes *yse_model.CommonRes, err error) {
	_, commonRes, err = s.Request[any](ctx, req, "smsc.revenueSharing.close", "1.0")
	if err != nil {
		return
	}
	return
}

// 登记订单分账
func (s *Client) SharingRegist(ctx context.Context, req *yse_model.SharingRegistReq) (res *yse_model.SharingRegistRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.SharingRegistRes](ctx, req, "unify.division.online.divRegist", "1.1")
	if err != nil {
		return
	}
	return
}

// 登记订单分账查询
func (s *Client) SharingRegistQuery(ctx context.Context, req *yse_model.SharingRegistQueryReq) (res *yse_model.SharingRegistQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.SharingRegistQueryRes](ctx, req, "unify.division.online.divQuery", "1.1")
	if err != nil {
		return
	}
	return
}

// // 分账退款
func (s *Client) SharingRefund(ctx context.Context, req *yse_model.SharingRefundReq) (res *yse_model.SharingRefundRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.SharingRefundRes](ctx, req, "unify.division.online.divRefund", "1.0")
	if err != nil {
		return
	}
	return
}
