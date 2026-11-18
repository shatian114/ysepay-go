package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 聚合收银台下单
func (s *Client) CreateOrder(ctx context.Context, req *yse_model.CreateOrderReq) (res *yse_model.CreateOrderRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.CreateOrderRes](ctx, req, "order.createOrder", "1.4")
	if err != nil {
		return
	}
	return
}

// 微信聚合扫码下单
func (s *Client) CreateOrderUnifyWechat(ctx context.Context, req *yse_model.CreateOrderUnifyWechatReq) (res *yse_model.CreateOrderUnifyWechatRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.CreateOrderUnifyWechatRes](ctx, req, "unify.basePay.scan.weChatPay.js", "1.4")
	if err != nil {
		return
	}
	return
}

// 支付宝聚合扫码下单
func (s *Client) CreateOrderUnifyAlipay(ctx context.Context, req *yse_model.CreateOrderUnifyAlipayReq) (res *yse_model.CreateOrderUnifyAlipayRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.CreateOrderUnifyAlipayRes](ctx, req, "unify.alipay.js", "1.4")
	if err != nil {
		return
	}
	return
}
