package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 进件资料上送
func (s *Client) CreateAccount(ctx context.Context, req *yse_model.CreateAccountReq) (res *yse_model.CreateAccountRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.CreateAccountRes](ctx, req, "t1.smsc.addCustInfoApply", "1.5")
	if err != nil {
		return
	}
	return
}

// 进件资料修改
func (s *Client) CreateAccountUpdate(ctx context.Context, req *yse_model.UpdateCreateAccountReq) (res *yse_model.CreateAccountRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.CreateAccountRes](ctx, req, "smsc.modifyCustInfoApply", "1.4")
	if err != nil {
		return
	}
	return
}
