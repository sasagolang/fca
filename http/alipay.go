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
	charge.CallbackURL = "http://218.17.28.42:8999/callback/aliappcallback"
	fmt.Printf("tradnum:%v,AliAppPay:%v\n", string(order.OrderNo), charge)
	fdata, err := logic.Pay(charge)
	if err != nil {
		fmt.Printf("AliAppPay:%v\n", err)
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
func InitClient() {
	ali := &client.AliAppClient{
		PartnerID:  "2088102169368862",
		SellerID:   "xuyanmei@glelec.com",
		AppID:      "2017010804929917",
		PrivateKey: nil,
		PublicKey:  nil,
	}
	client.InitAliAppClient(ali)
	InitKeys(ali)
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
	//payType := r.PathParam("PayType")
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	amount, _ := strconv.Atoi(r.PathParam("Amount"))
	s := AliAppPay(uid, int64(amount))
	var err error
	if err != nil {
		fmt.Printf("GetCarBrands error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", s)
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
