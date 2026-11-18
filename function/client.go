package yse_func

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json/v2"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/emmansun/gmsm/sm2"
	"github.com/emmansun/gmsm/smx509"
	"github.com/shatian114/putils-go/pconv"
	"github.com/shatian114/putils-go/prand"
	"github.com/shatian114/verifysign"
	ysepaygo "github.com/shatian114/ysepay-go"
	yse_model "github.com/shatian114/ysepay-go/model"
)

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{
		SessionTicketsDisabled: true,
		NextProtos:             nil,
		MinVersion:             tls.VersionTLS12,
		MaxVersion:             tls.VersionTLS12,
	},
}

type Client struct {
	CertId      string
	PrivKey     *sm2.PrivateKey
	PubCert     *smx509.Certificate
	KeyPassword string
}

func getCommonReq(method, version string) *yse_model.CommonReq {
	now := time.Now()
	return &yse_model.CommonReq{
		ReqID:     prand.RandStr(18) + now.Format("20060102150405"),
		TimeStamp: now.Format("2006-01-02 15:04:05"),
		Method:    method,
		Charset:   "utf-8", // 固定值
		Version:   version,
	}
}

func (s *Client) Request[T any](ctx context.Context, paramObj any, method, apiVer string, fileList ...[]byte) (res T, commonRes *yse_model.CommonRes, err error) {

	ysepaygo.Logger.Debug(ctx, "开始http请求")
	// path 是把 method 里的 . 更换为 /
	path := strings.ReplaceAll(method, ".", "/")

	commonReq := getCommonReq(method, apiVer)
	commonReq.CertID = s.CertId
	paramByte, err := json.Marshal(paramObj)
	ysepaygo.Logger.Debug(ctx, "请求业务参数: ", string(paramByte))
	if err != nil {
		err = fmt.Errorf("请求，参数转json出错: %s", err.Error())
		return
	}

	aesKey := []byte(prand.RandStr(16))
	hexAesKey := hex.EncodeToString(aesKey)
	checkByte, err := ysepaygo.EncryptData(s.PubCert, []byte(hexAesKey))
	if err != nil {
		ysepaygo.Logger.Error(ctx, "check字段加密出错: ", err)
		return
	}
	commonReq.Check = string(checkByte)

	commonReq.BizContent, err = ysepaygo.AESEncryptECB(paramByte, aesKey)
	if err != nil {
		return
	}

	ysepaygo.Logger.Debug(ctx, "开始构造签名")
	// 构造签名
	signMap := pconv.AnyToMap(commonReq)
	delete(*signMap, "sign")
	signOriginStr := verifysign.BuildSignStrByBody(*signMap)
	// 将签名进行Base64编码
	commonReq.Sign, err = ysepaygo.SignMsgSM2(s.PrivKey, signOriginStr)
	if err != nil {
		ysepaygo.Logger.Error(ctx, "生成签名出错: ", err)
		return
	}

	var httpReq *http.Request

	// 判断是文件上传还是普通 post application/json
	if len(fileList) > 0 {
		file := fileList[0]
		body := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(body)
		// 创建文件字段
		part, _ := writer.CreateFormFile("file", "1gtgrt.jpg")
		// part.Write(bArr)
		io.Copy(part, bytes.NewReader(file))

		// 创建普通文本字段
		// 循环遍历 commonReq

		part, _ = writer.CreateFormField("timeStamp")
		io.Copy(part, bytes.NewReader([]byte(commonReq.TimeStamp)))

		part, _ = writer.CreateFormField("method")
		io.Copy(part, bytes.NewReader([]byte(commonReq.Method)))

		part, _ = writer.CreateFormField("charset")
		io.Copy(part, bytes.NewReader([]byte(commonReq.Charset)))

		part, _ = writer.CreateFormField("sign")
		io.Copy(part, bytes.NewReader([]byte(commonReq.Sign)))

		part, _ = writer.CreateFormField("check")
		io.Copy(part, bytes.NewReader([]byte(commonReq.Check)))

		part, _ = writer.CreateFormField("bizContent")
		io.Copy(part, bytes.NewReader([]byte(commonReq.BizContent)))

		part, _ = writer.CreateFormField("reqId")
		io.Copy(part, bytes.NewReader([]byte(commonReq.ReqID)))

		part, _ = writer.CreateFormField("certId")
		io.Copy(part, bytes.NewReader([]byte(commonReq.CertID)))

		part, _ = writer.CreateFormField("version")
		io.Copy(part, bytes.NewReader([]byte(commonReq.Version)))

		part, _ = writer.CreateFormField("meta")
		metaJsonByte, _ := json.Marshal(paramByte)
		io.Copy(part, bytes.NewReader(metaJsonByte))

		writer.Close()

		httpReq, err = http.NewRequest(http.MethodPost, ysepaygo.ProUrlPrefix+path, body)
		contentType := writer.FormDataContentType()
		ysepaygo.Logger.Info(ctx, "设置请求体类型: ", contentType)
		httpReq.Header.Set("Content-Type", contentType)
	} else {
		pj, _ := json.Marshal(commonReq)
		httpReq, err = http.NewRequest(http.MethodPost, ysepaygo.ProUrlPrefix+path, bytes.NewReader(pj))
		httpReq.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		err = fmt.Errorf("请求，创建http request出错: %s", err)
		return
	}

	reqAllByte, err := httputil.DumpRequest(httpReq, true)
	if err != nil {

		ysepaygo.Logger.Error(ctx, "获取请求数据出错: ", err.Error())
	}
	ysepaygo.Logger.Debug(ctx, "请求: ", reqAllByte)

	httpClient := &http.Client{Transport: tr, Timeout: 5 * time.Second}
	httpRes, err := httpClient.Do(httpReq)

	if err != nil {
		err = fmt.Errorf("银盛接口请求出错: %s", err.Error())
		return
	}
	defer httpRes.Body.Close()

	resAllByte, err := httputil.DumpResponse(httpRes, true)
	if err != nil {
		ysepaygo.Logger.Error(ctx, "获取银盛请响应据出错: ", err.Error())
	}
	ysepaygo.Logger.Debug(ctx, "银盛响应: ", resAllByte)

	resByte, err := io.ReadAll(httpRes.Body)
	if err != nil {
		err = fmt.Errorf("银盛请求，读取响应byte出错: %s", err)
		return
	}
	resStr := string(resByte)
	if !strings.HasPrefix(resStr, "{") {
		var bsByte []byte
		bsByte, err = base64.StdEncoding.DecodeString(resStr)
		if err != nil {
			return
		}
		resStr = string(bsByte)
	}

	ysepaygo.Logger.Debug(ctx, "银盛返回: ", resStr)

	err = json.Unmarshal([]byte(resStr), &commonRes)
	if err != nil {
		ysepaygo.Logger.Error(ctx, "银盛结果转 struct 出错: ", err)
		err = errors.New("银盛结果转换出错")
		return
	}

	// 判断结果是否正确
	if commonRes.Code != "00000" {
		err = fmt.Errorf("银盛公共出错: %s", commonRes.Msg)
		return
	}

	if commonRes.SubCode != "0000" {
		err = fmt.Errorf("银盛业务出错: %s", commonRes.SubMsg)
		return
	}
	// 验签
	// 构建待验签字符串
	signMap = pconv.AnyToMap(commonRes)
	delete(*signMap, "sign")

	signOriginStr = verifysign.BuildSignStrByBody(*signMap)
	ok, err := ysepaygo.VerifyMsgSignSM2Base64(s.PubCert, strings.TrimRight(commonRes.Sign, " "), []byte(signOriginStr))
	if err != nil {
		ysepaygo.Logger.Error(ctx, "银盛验签出错")
		return
	}
	if !ok {
		err = fmt.Errorf("银盛验签不通过")
		ysepaygo.Logger.Error(ctx, "银盛验签不通过")
		return
	}
	ysepaygo.Logger.Debug(ctx, "银盛验签通过，开始解密内容")

	if commonRes.BusinessData == nil {
		ysepaygo.Logger.Info(ctx, "业务字符串为空，不解析")
		return
	}

	aesHexByte, err := hex.DecodeString(hexAesKey)
	if err != nil {
		ysepaygo.Logger.Error(ctx, "aeshex 解析出错: ", err)
		return
	}
	businessDeByte, err := ysepaygo.AESDecryptECB(*commonRes.BusinessData, aesHexByte)
	if err != nil {
		ysepaygo.Logger.Error(ctx, "银盛结果里的业务数据解密出错: ", err)
		return
	}
	ysepaygo.Logger.Debug(ctx, "银盛结果业务内容: ", string(businessDeByte))

	// 字符串结果不用转 json
	if _, ok := any(res).(string); ok {
		res = any(string(businessDeByte)).(T)
	} else {
		err = json.Unmarshal(businessDeByte, &res)
	}
	return

}
