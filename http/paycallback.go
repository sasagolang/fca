package http

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/guidao/gopay/client"
	"github.com/guidao/gopay/common"
	"github.com/guidao/gopay/util"
)

func AliWebCallback(w http.ResponseWriter, r *http.Request) (*common.AliWebPayResult, error) {
	var m = make(map[string]string)
	var signSlice []string
	r.ParseForm()
	for k, v := range r.Form {
		// k不会有多个值的情况
		m[k] = v[0]
		if k == "sign" || k == "sign_type" {
			continue
		}
		signSlice = append(signSlice, fmt.Sprintf("%s=%s", k, v[0]))
	}

	sort.Strings(signSlice)
	signData := strings.Join(signSlice, "&")
	if m["sign_type"] != "RSA" {
		//错误日志
		return nil, errors.New("签名类型未知")
	}

	err := client.DefaultAliWebClient().CheckSign(signData, m["sign"])
	if err != nil {
		//log.Error("签名验证失败：", err, signData, m)
		return nil, err
	}

	var aliPay common.AliWebPayResult
	err = util.MapStringToStruct(m, &aliPay)
	if err != nil {
		//log.Error(err)
		w.Write([]byte("error"))
		return nil, err
	}

	// err = biz.AliWebCallBack(&aliPay)
	// if err != nil {
	// 	//log.Error(err)
	// 	w.Write([]byte("error"))
	// 	return nil, err
	// }
	w.Write([]byte("success"))
	return &aliPay, nil
}

// 支付宝app支付回调
func AliAppCallback(w http.ResponseWriter, r *http.Request) (*common.AliWebPayResult, error) {
	var m = make(map[string]string)
	var signSlice []string
	r.ParseForm()
	for k, v := range r.Form {
		m[k] = v[0]
		if k == "sign" || k == "sign_type" {
			continue
		}
		signSlice = append(signSlice, fmt.Sprintf("%s=%s", k, v[0]))
	}
	sort.Strings(signSlice)
	signData := strings.Join(signSlice, "&")
	if m["sign_type"] != "RSA" {
		//log.Error(m)
		return nil, errors.New("签名类型未知")
	}

	err := client.DefaultAliAppClient().CheckSign(signData, m["sign"])
	if err != nil {
		//log.Error(err, m, signData)
		w.Write([]byte("error"))
		return nil, err
	}

	var aliPay common.AliWebPayResult
	err = util.MapStringToStruct(m, &aliPay)
	if err != nil {
		//log.Error(err)
		w.Write([]byte("error"))
		return nil, err
	}

	// err = biz.AliAppCallBack(&aliPay)
	// if err != nil {
	// 	//log.Error(err)
	// 	w.Write([]byte("error"))
	// }

	w.Write([]byte("success"))
	return &aliPay, nil
}
