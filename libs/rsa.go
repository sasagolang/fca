package libs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 公钥和私钥可以从文件中读取
var AliAppPrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEArsHwLh/ZWdf1sE9JiCmnv3xoc1zZyBUIm4on2r0XzLkYG9Hn7a/9iPXSvMuVRGoEV1aUSrLfxeM7+AOrxjtUwJa8LyyDJSubgYhZ2jKMmdPP+WIuba95l0PMWDZlLeeQByurAxY4UyI4VRL3xKpkm+GshskC51K8KWcqmBBXqu0IxU/IKKxFceZ9sxuraHE1W8XSqRn33LX50shCWj0xoZnfYWU9kWjwX/XlObe/H65QoRYGL9Lpm0xoYSY7j4id68HNPVP2vqgl+dUGCIkaTfwlIAWJcfPtVzMJ3zTr/CdBRzlRwK1WcjwFAN4nFiEVNE/YB+Ydcy/x5YGP5aGVfQIDAQABAoIBAGvO/MqXn1CLp+fioew8Nl4s84rnpvqF7091BW0t5jbHoRXYgonhXjlV3XHW7fjMpmzy3nzOqLm9m0s/iIK8K2a8Fs/LWARgSIi6gyPWt2JWhyDD7vx/mjU1ACUTOYv+JJ4n1FINLSYsIDBriSpsK7D2oaZs7zuhkK7CFEbHHi41E5CM2xuaIXIERqv7ICLCVAaAWthb7E3rXYLb6lEsv+DhGUPgF4jmEphtcl899+N0KJXaxDgEpTqBn+e9/u6UM+y40xblMeg0nfqKxwPtbQn6Brzi3dKNflmTkSq0jb7bzTFcUjbX981p4ZqGkx5IT0vgqUGZN2SIRsFyrl0aXMkCgYEA1LDs3/c/kxdbBjwedusiexmEq8yzUjwyF97gSy20ZBa6PaOH3/mPZb8GIXQ6n+aLvtOgm883bpao6vQDMxyDxUpJpunTJ9Qv65EvdmMG11ywimGw+eTrULgpR2s/m4BA4XE6VpfZWzLfKSRNqbpWfANE/8gcQS/xf+T5M6Ohfg8CgYEA0leg+YonBrGvxQREt5YmrfG8OS5YDdOG/Sc/wbNCK3xA1rVb0WkcJNvSwml7z3hG2exWVoP1mfLaaReX09pF2/sWNWXgalSj4RvI8l6Z35xqVcuuN2bEjOyHYWH6MyVBsXJ5SpevtUIc3+LYNRfXOhjxfAd+oXqGFtysEtwDf7MCgYBOAJDu1nt/U2pzj+rQTA26PbKVWx5Mw3zPmlKB38IvjtJAts+nCZxYgUUbUcgKSn/nvS9C1S0MJr7OZC0kOons+gCm8UwaaEwmxXk1nr2sj/bC2W8RRq4yTUf+REvwmImy5Faz2T7CnpaPPRwqagc7tetBLz+FfLLB59So5pLcawKBgAsFgG/S51yYbudZ4+fivEAdDInKfd6rmMUnC1Yw+GSi7BrUAe2lHk5oHlEFifFDEiNVGOkLnRDmAr/C9repFkQCkhVWMz4fFT80X4Ejp9hpr8CzHXvVrLLdqfJWWe/YIesUXnqkHBbZUf4BOub9Ss/GgtDG68G2U9Ra18FdOem1AoGBAKTBedMw5K4e3iqsErwH4mbQED1fwyY/nRpC13iOHpO93ma2vNbE4ofm3dRtBI/SxwxrwAz+SoW3DWaqGzQ5mBpgZOxmsqPEhch0G6w09o4xx+dS/J+vN8k42AYvAU1rRjsGv9H/Mx7CaxxVqWNUEAjn+FQYooytbTeyDHsmn79+
-----END RSA PRIVATE KEY-----
`)

// var AliAppPublicKey = []byte(`
// -----BEGIN PUBLIC KEY-----
// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArsHwLh/ZWdf1sE9JiCmnv3xoc1zZyBUIm4on2r0XzLkYG9Hn7a/9iPXSvMuVRGoEV1aUSrLfxeM7+AOrxjtUwJa8LyyDJSubgYhZ2jKMmdPP+WIuba95l0PMWDZlLeeQByurAxY4UyI4VRL3xKpkm+GshskC51K8KWcqmBBXqu0IxU/IKKxFceZ9sxuraHE1W8XSqRn33LX50shCWj0xoZnfYWU9kWjwX/XlObe/H65QoRYGL9Lpm0xoYSY7j4id68HNPVP2vqgl+dUGCIkaTfwlIAWJcfPtVzMJ3zTr/CdBRzlRwK1WcjwFAN4nFiEVNE/YB+Ydcy/x5YGP5aGVfQIDAQAB
// -----END PUBLIC KEY-----
// `)
var AliAppPublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnxj/9qwVfgoUh/y2W89L6BkRAFljhNhgPdyPuBV64bfQNN1PjbCzkIM6qRdKBoLPXmKKMiFYnkd6rAoprih3/PrQEB/VsW8OoM8fxn67UDYuyBTqA23MML9q1+ilIZwBC2AQ2UBVOrFXfFl75p6/B5KsiNG9zpgmLCUYuLkxpLQIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(AliAppPrivateKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(AliAppPublicKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
