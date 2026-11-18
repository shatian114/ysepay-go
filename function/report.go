package yse_func

import (
	"context"

	yse_model "github.com/shatian114/ysepay-go/model"
)

// 给报备新增appid
func (s *Client) ReportAppidAdd(ctx context.Context, req *yse_model.ReportAppidAddReq) (res *yse_model.ReportAppidAddRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.ReportAppidAddRes](ctx, req, "report.scan.union.appIdAddOrUpdate", "1.0")
	if err != nil {
		return
	}
	return
}

// 查询报备
func (s *Client) ReportQuery(ctx context.Context, req *yse_model.ReportQueryReq) (res *yse_model.ReportQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.ReportQueryRes](ctx, req, "report.scan.union.repAndAppIdQry", "1.0")
	if err != nil {
		return
	}
	return
}

// 查询商户报备信息
func (s *Client) MercReportQuery(ctx context.Context, req *yse_model.MercReportQueryReq) (res *yse_model.MercReportQueryRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.MercReportQueryRes](ctx, req, "cust.reportSearchServer.queryReportResult", "1.0")
	if err != nil {
		return
	}
	return
}

// 新增报备
func (s *Client) ReportAdd(ctx context.Context, req *yse_model.ReportAddReq) (res *yse_model.ReportAddRes, commonRes *yse_model.CommonRes, err error) {
	res, commonRes, err = s.Request[*yse_model.ReportAddRes](ctx, req, "report.scan.union.reportAdd", "1.0")
	if err != nil {
		return
	}
	return
}
