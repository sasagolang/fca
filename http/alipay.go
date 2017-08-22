package http

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fca/libs"
	"fca/logic"
	"fmt"
	"log"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/guidao/gopay/client"
	"github.com/guidao/gopay/common"
	"github.com/guidao/gopay/constant"
)

func AliAppPay(uid int, paymentAmout int64) string {
	//Create Order
	order := Logic.CreateOrder(uid, paymentAmout, "AliApp")
	charge := new(common.Charge)
	charge.PayMethod = constant.ALI_APP
	charge.MoneyFee = order.PaymentAmount
	charge.Describe = "充值"
	charge.TradeNum = strconv.FormatInt(order.OrderNo, 10)
	charge.CallbackURL = "https://appgphitec.com/callback/aliappcallback"
	fmt.Printf("tradnum:%v,AliAppPay:%v\n", string(order.OrderNo), charge)
	fdata, err := logic.Pay(charge)
	if err != nil {
		fmt.Printf("AliAppPay:%v\n", err)
	}
	return fdata
}
func WeChatAppPay(uid int, paymentAmout int64) string {
	//Create Order
	order := Logic.CreateOrder(uid, paymentAmout, "WeChatApp")
	charge := new(common.Charge)
	charge.PayMethod = constant.WECHAT
	charge.MoneyFee = order.PaymentAmount
	charge.Describe = "充值"
	charge.TradeNum = strconv.FormatInt(order.OrderNo, 10)
	charge.CallbackURL = "https://appgphitec.com/callback/wechatappcallback"
	fmt.Printf("tradnum:%v,WeChatAppPay:%v\n", string(order.OrderNo), charge)
	fdata, err := logic.Pay(charge)
	if err != nil {
		fmt.Printf("WeChatAppPay:%v\n", err)
	}
	return fdata
}
func TestPay() {
	//initClient()
	//	initHandle()
	charge := new(common.Charge)
	charge.PayMethod = constant.ALI_APP
	charge.MoneyFee = 1
	charge.Describe = "充值"
	charge.TradeNum = "1111111111"
	charge.CallbackURL = "http://www.aaa.com"
	fdata, err := logic.Pay(charge)
	if err != nil {
		fmt.Printf("TestPay:%v\n", err)
	}
	fmt.Println(fdata)
}

//2017031306200530
func InitClient() {
	ali := &client.AliAppClient{
		PartnerID:  "2088521469262911",
		SellerID:   "xuyanmei@glelec.com",
		AppID:      "2017031306200530",
		PrivateKey: nil,
		PublicKey:  nil,
	}
	client.InitAliAppClient(ali)
	InitKeys(ali)

	wechatapp := &client.WechatAppClient{
		AppID: "wx02e17b72574934e7",
		MchID: "1487065122",

		Key:    "ec76f5c7377e3291603216b1de793539",
		PayURL: "https://api.mch.weixin.qq.com/pay/unifiedorder",
	}
	client.InitWechatClient(wechatapp)
	//	fmt.Printf("initClient:%v\n", ali)
}
func InitKeys(c *client.AliAppClient) error {

	//	log.Println("init rsakeys begin")

	block, _ := pem.Decode(libs.AliAppPrivateKey)
	if block == nil {
		log.Println("rsaSign private_key error")
		return fmt.Errorf("rsaSign pem.Decode error")
	}
	var err error
	c.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println("rsaSign ParsePKIXPublicKey error : %v\n", err)
		return err
	}

	block, _ = pem.Decode(libs.AliAppPublicKey)
	if block == nil {
		log.Println("public key error")
		return err
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println("rsaSign ParsePKIXPublicKey error : %v\n", err)
		return err
	}

	c.PublicKey = pubInterface.(*rsa.PublicKey)
	//	log.Println("init rsakeys success ")
	return err
}
func PayFunc(w rest.ResponseWriter, r *rest.Request) {
	var rstr string
	payType := r.PathParam("PayType")
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	amount, _ := strconv.Atoi(r.PathParam("Amount"))
	if payType == "1" {
		rstr = AliAppPay(uid, int64(amount))
	} else if payType == "2" {
		rstr = WeChatAppPay(uid, int64(amount))
	}
	var err error
	if err != nil {
		fmt.Printf("GetCarBrands error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", rstr)
}

/*func initHandle() {
	http.HandleFunc("callback/aliappcallback", func(w http.ResponseWriter, r *http.Request) {
		aliResult, err := AliAppCallback(w, r)
		if err != nil {
			fmt.Println(err)
			//log.xxx
			return
		}
		selfHandler(aliResult)
	})
}

*/

func selfHandler(i interface{}) {
	fmt.Printf("selfHandler:%v\n", i)
}
