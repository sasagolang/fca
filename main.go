package main

import (
	"encoding/base64"
	"fca/dal"
	htp "fca/http"
	"fmt"
)

func main() {
	//fmt.Println("fffdf")
	//Dao.GetUserByUID(123)
	//dal.GetUserByUID(123)
	//test()
	//htp.TestPay()
	dal.InitDb()
	htp.Logic.RegisterUser("11232323", "", "", "123")
	order := htp.Logic.CreateOrder(192858680, 200, "aliapp")
	htp.Logic.DepositFunc(192858680, 100, "dfdfdf", order.OrderNo)
	htp.InitClient()
	htp.Init()

}
func test() {
	var data []byte
	var err error
	var decrypted string
	//decrypted = "polaris@studygolang.com"
	if decrypted != "" {
		//data, err = base64.StdEncoding.DecodeString(decrypted)
		///if err != nil {
		//	panic(err)
		//}
	} else {
		//	data, err = libs.RsaEncrypt([]byte("polaris@studygolang.com"))
		if err != nil {
			panic(err)
		}
		fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	}
	//	origData, err := libs.RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	//	fmt.Println(string(origData))
}
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
