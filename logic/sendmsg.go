package logic

import (
	"encoding/json"
	"errors"
	"fca/dal"
	"fca/libs"
	"fca/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (logic LogicBase) CheckCodeByMobile(mobile string, code string) (err error) {
	var smsinfo model.SMSInfo
	if !dal.DB.Where("mobile=?", mobile).Last(&smsinfo).RecordNotFound() {
		if time.Now().Sub(smsinfo.SendTime).Minutes() > 60*24 {
			return errors.New("验证码错误(001)！")
		}
		if smsinfo.Code != code {
			return errors.New("验证码错误(003)！")
		}
	} else {
		return errors.New("验证码错误(002)！")

	}
	return nil
}
func (logic LogicBase) SendMsgByMobile(mobile string, code string) (err error) {

	var smsinfo model.SMSInfo

	//查询发送
	if !dal.DB.Where("mobile=?", mobile).Last(&smsinfo).RecordNotFound() {
		if time.Now().Sub(smsinfo.SendTime).Minutes() < 60*24 {
			return errors.New("在24小时内已发送过验证码！")
		}
	}

	logic.SendMsgAPI(mobile, code)

	return nil
}

func (logic LogicBase) SendMsg(uid int, code string) (err error) {
	user, err := logic.GetUserByUid(uid)
	if err != nil {
		var smsinfo model.SMSInfo

		if user.Mobile == "" {
			return errors.New("未绑定手机号")
		}
		//查询发送
		if !dal.DB.Where("uid=?", user.UID).Last(&smsinfo).RecordNotFound() {
			if time.Now().Sub(smsinfo.SendTime).Minutes() < 60*24 {
				return errors.New("在24小时内已发送过验证码！")
			}
		}

		logic.SendMsgAPI(user.Mobile, code)

	} else {
		return errors.New("用户不存在")
	}
	return nil
}

func (logic LogicBase) SendMsgAPI(mobile string, code string) (rcode int, err error) {
	smsinfo := new(model.SMSInfo)
	smsinfo.Code = code
	smsinfo.Mobile = mobile
	smsinfo.SendTime = time.Now()
	dal.DB.NewRecord(smsinfo)
	dal.DB.Create(smsinfo)
	params := url.Values{}
	params.Add("mobile", mobile)
	params.Add("tpl_id", "29840")
	params.Add("tpl_value", "#code#="+code)
	params.Add("key", "df9840a6c6bc7b9c9d480f494d87b1e6")
	params.Add("dtype", "json")
	juheURL, _ := url.Parse("http://v.juhe.cn/sms/send")
	juheURL.RawQuery = params.Encode()
	libs.Log.Info(fmt.Sprintf("SendMsg:%s\n", juheURL.String()))
	resp, err := http.Get(juheURL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if buf, err := ioutil.ReadAll(resp.Body); err != nil {
		var result SMSResult
		json.Unmarshal(buf, &result)
		return result.error_code, nil
	}

	return
}

type SMSResult struct {
	reason     string
	result     Result
	error_code int
}
type Result struct {
	count int
	fee   int
	sid   string
}
