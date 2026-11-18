package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// --------------- 支付宝 ---------------------
// 支付宝实名认证请求
func (s *Client) AlipaySubmitAuthMessage(ctx context.Context, req *yse_model.AlipaySubmitAuthMessageReq) (res *yse_model.AlipaySubmitAuthMessageRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.AlipaySubmitAuthMessageRes](ctx, req, "pregate.alipay.submitAuthMessages", "1.0")
	if err != nil {
		return
	}
	return
}

// 支付宝实名认证查询
func (s *Client) AlipaySubmitAuthMessageQuery(ctx context.Context, req *yse_model.AlipaySubmitAuthMessageQueryReq) (res *yse_model.AlipaySubmitAuthMessageQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.AlipaySubmitAuthMessageQueryRes](ctx, req, "pregate.alipay.getAuthMessagesByApplyNo", "1.0")
	if err != nil {
		return
	}
	return
}

// 支付宝授权状态查询
func (s *Client) AlipayAuthQuery(ctx context.Context, req *yse_model.AlipayAuthStatusQueryReq) (res *yse_model.AlipayAuthStatusQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.AlipayAuthStatusQueryRes](ctx, req, "pregate.alipay.getAuthState", "1.0")
	if err != nil {
		return
	}
	return
}

// --------------- 微信 ---------------------

// 微信实名认证请求
func (s *Client) WechatSubmitAuthMessage(ctx context.Context, req *yse_model.WechatSubmitAuthMessageReq) (res *yse_model.WechatSubmitAuthMessageRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.WechatSubmitAuthMessageRes](ctx, req, "trade.scan.wechat.auth.submitAuthMessages", "1.0")
	if err != nil {
		return
	}
	return
}

// 微信实名认证查询
func (s *Client) WechatSubmitAuthMessageQuery(ctx context.Context, req *yse_model.WechatSubmitAuthMsgQueryReq) (res *yse_model.WechatSubmitAuthMsgQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.WechatSubmitAuthMsgQueryRes](ctx, req, "trade.scan.wechat.auth.getAuthMessagesByApplyNo", "1.0")
	if err != nil {
		return
	}
	return
}

// 微信实名认证详情查询
func (s *Client) WechatSubmitAuthMessageDetailQuery(ctx context.Context, req *yse_model.WechatAuthMsgByMercIdReq) (res *yse_model.WechatAuthMsgByMercIdRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.WechatAuthMsgByMercIdRes](ctx, req, "trade.scan.wechat.getAuthMessagesByMercId", "1.0")
	if err != nil {
		return
	}
	return
}

// 微信授权状态查询
func (s *Client) WechatAuthQuery(ctx context.Context, req *yse_model.WechatAuthQueryReq) (res *yse_model.WechatAuthQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.WechatAuthQueryRes](ctx, req, "trade.scan.wechat.auth.getAuthState", "1.0")
	if err != nil {
		return
	}
	return
}
