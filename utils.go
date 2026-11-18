package ysepaygo

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"github.com/emmansun/gmsm/cfca"
	"github.com/emmansun/gmsm/pkcs"
	"github.com/emmansun/gmsm/sm2"
	"github.com/emmansun/gmsm/smx509"
)

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7 unpadding implementation
func unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// aes ecb 加密
func AESEncryptECB(plainText []byte, aesKey []byte) (string, error) {
	// Create AES cipher
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	// Pad the plain text
	padded := pad([]byte(plainText))

	// Encrypt the padded data
	encrypted := make([]byte, len(padded))
	for i := 0; i < len(padded); i += aes.BlockSize {
		block.Encrypt(encrypted[i:i+aes.BlockSize], padded[i:i+aes.BlockSize])
	}

	// Return the base64-encoded encrypted data
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// aes ecb 解密
func AESDecryptECB(cipherText string, aesKey []byte) ([]byte, error) {
	// Create AES cipher
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	// Decode the base64-encoded input
	decoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	// Decrypt the input data
	decrypted := make([]byte, len(decoded))
	for i := 0; i < len(decoded); i += aes.BlockSize {
		block.Decrypt(decrypted[i:i+aes.BlockSize], decoded[i:i+aes.BlockSize])
	}

	// Unpad the decrypted data
	unpadded := unpad(decrypted)

	// Return the decrypted plaintext
	return unpadded, nil
}

func decodeCFCAKeyFile(data []byte) ([]byte, error) {
	data = bytes.TrimSpace(data)

	// 支持 PEM 包装的 CFCA KEY。
	if bytes.HasPrefix(data, []byte("-----BEGIN")) {
		block, _ := pem.Decode(data)
		if block == nil {
			return nil, fmt.Errorf("无法解析 PEM 私钥文件")
		}
		if block.Type != "CFCA KEY" {
			return nil, fmt.Errorf("PEM 类型不正确: %s", block.Type)
		}
		return block.Bytes, nil
	}

	// 如果文件已经是 DER，就直接返回。
	if len(data) > 0 && data[0] == 0x30 {
		return data, nil
	}

	// 普通多行 Base64：移除换行和空白后解码。
	encoded := strings.Join(strings.Fields(string(data)), "")
	der, err := base64.StdEncoding.Strict().DecodeString(encoded)
	if err != nil {
		// 兼容未带 Base64 padding 的内容。
		der, err = base64.RawStdEncoding.Strict().DecodeString(encoded)
		if err != nil {
			return nil, fmt.Errorf("私钥文件不是有效的 Base64 或 DER 数据: %w", err)
		}
	}
	return der, nil
}

func SignMsgSM2(priv *sm2.PrivateKey, msg string) (string, error) {

	sig, err := priv.SignMessage(rand.Reader, []byte(msg), sm2.DefaultSM2SignerOpts)
	if err != nil {
		return "", fmt.Errorf("SM2 签名失败: %w", err)
	}

	sign := base64.StdEncoding.EncodeToString(sig)

	// 对齐 Java String.format("%-256s", checkValue) 的右侧补空格行为。
	if len(sign) < 256 {
		sign += strings.Repeat(" ", 256-len(sign))
	}
	return sign, nil
}

func EncryptData(cert *smx509.Certificate, data []byte) (string, error) {

	// 对应 Java 的 Mechanism.SM4_ECB。
	envelopeDER, err := cfca.EnvelopeMessage(
		pkcs.SM4ECB,
		data, // Go 字符串字节为 UTF-8
		[]*smx509.Certificate{cert},
	)
	if err != nil {
		return "", fmt.Errorf("CFCA 数字信封加密失败: %w", err)
	}

	// Java 方法返回 Base64 信封字符串；gmsm 返回 DER，所以这里编码为 Base64。
	return base64.StdEncoding.EncodeToString(envelopeDER), nil
}

func GetPubCertFromBase64(ctx context.Context, base64Str string) (cert *smx509.Certificate, err error) {
	certBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return
	}

	// 证书可以是 PEM 或 DER 格式。
	der := certBytes
	if block, _ := pem.Decode(certBytes); block != nil {
		if block.Type != "CERTIFICATE" {
			err = errors.New("银盛公钥证书类型不正确")
			Logger.Debug(ctx, "银盛公钥证书类型是: ", block.Type)
			return
		}
		der = block.Bytes
	}

	cert, err = smx509.ParseCertificate(der)
	return
}

func GetPrivKeyFromBase64(ctx context.Context, privBase64 string, password string) (privKey *sm2.PrivateKey, err error) {
	fileData, err := base64.StdEncoding.DecodeString(privBase64)
	if err != nil {
		return
	}

	der, err := decodeCFCAKeyFile(fileData)
	if err != nil {
		return
	}

	// ParseSM2 的参数顺序是：密码、DER 编码的 CFCA 私钥容器。
	privKey, _, err = cfca.ParseSM2([]byte(password), der)
	if err != nil {
		return
	}
	return
}

func VerifyMsgSignSM2(
	cert *smx509.Certificate,
	sigDER []byte,
	msg []byte,
) (bool, error) {
	if cert == nil {
		return false, fmt.Errorf("证书不能为空")
	}

	pub, ok := cert.PublicKey.(*ecdsa.PublicKey)
	if !ok || !sm2.IsSM2PublicKey(pub) {
		return false, fmt.Errorf("证书公钥不是 SM2")
	}

	// SM2 常见默认 UID；需与签名方使用的 UID 一致。1234567812345678
	uid := []byte("1234567812345678")
	return sm2.VerifyASN1WithSM2(pub, uid, msg, sigDER), nil
}

// VerifyMsgSignSM2Base64 接收接口返回的 Base64 签名文本。
// 可处理 Java %-256s 产生的首尾空格。
func VerifyMsgSignSM2Base64(
	cert *smx509.Certificate,
	signBase64 string,
	msg []byte,
) (bool, error) {
	sigDER, err := base64.StdEncoding.DecodeString(strings.TrimSpace(signBase64))
	if err != nil {
		return false, fmt.Errorf("Base64 签名解码失败: %w", err)
	}
	return VerifyMsgSignSM2(cert, sigDER, msg)
}
