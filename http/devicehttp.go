package http

import (
	"io/ioutil"
	"strconv"

	"fmt"
	"net/http"

	"encoding/json"

	"github.com/ant0ine/go-json-rest/rest"
)

var DeviceUrl string

func StartChargeFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	no, _ := strconv.Atoi(r.PathParam("no"))
	mins, _ := strconv.Atoi(r.PathParam("mins"))
	pole := Logic.GetPoleByNo(no)
	wallet, _ := Logic.MyWallet(uid)
	if pole != nil {
		if wallet.Balance <= 0 {
			WriterResponse(w, 5, "余额不足，请先充值!", nil)
			return
		}
		if pole.Status != 2004 {
			WriterResponse(w, 4, "请连接充电枪后开始充电！", nil)
			return
		}

		response, err := http.Get(fmt.Sprintf("%s/start/%s/%d", DeviceUrl, pole.UUID, mins))
		fmt.Printf(fmt.Sprintf("err:%v\n", err))
		if response != nil {
			defer response.Body.Close()
		}
		if err == nil && response != nil {
			var orderno int64
			body, _ := ioutil.ReadAll(response.Body)
			var result Result
			json.Unmarshal(body, &result)
			fmt.Printf(fmt.Sprintf("result:%v\n", result))
			if result.ResultCode == 1 {
				order := Logic.CreateChargeOrder(uid, no, pole.UUID, pole.ElectricPile.Name, pole.ElectricPile.Address)
				orderno = order.NO
			}
			WriterResponse(w, result.ResultCode, result.ResultMessage, orderno)
			fmt.Printf(fmt.Sprintf("%s/start/%s/%d", DeviceUrl, pole.UUID, mins))
		} else {
			WriterResponse(w, 3, err.Error(), nil)
		}
	} else {
		WriterResponse(w, 2, "充电桩不存在", nil)
	}
	fmt.Printf("StartChargeFunc:%v,%v", uid, no)
}
func EndChargeFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	no, _ := strconv.ParseInt(r.PathParam("no"), 10, 64)
	fmt.Printf("EndChargeFunc:%v,%v,%s", uid, no, r.PathParam("no"))
	pole := Logic.GetChargeOrderByNo(no)
	if pole != nil && len(pole.UUID) > 0 && pole.Status == 1 {
		code, msg := Logic.EndChargeTCP(pole.UUID, no)
		WriterResponse(w, code, msg, nil)
	} else {
		WriterResponse(w, 2, "充电桩不存在", nil)
	}
}

func StartLightFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	no, _ := strconv.Atoi(r.PathParam("no"))
	fmt.Printf("StartLightFunc:%v,%v", uid, no)
	pole := Logic.GetPoleByNo(no)
	if pole != nil {
		response, err := http.Get(fmt.Sprintf("%s/startlight/%s", DeviceUrl, pole.UUID))
		if response != nil {
			defer response.Body.Close()
		}
		if err == nil && response != nil {

			body, _ := ioutil.ReadAll(response.Body)
			var result Result
			json.Unmarshal(body, &result)

			WriterResponse(w, result.ResultCode, result.ResultMessage, nil)

		} else {
			WriterResponse(w, 3, err.Error(), nil)
		}
	} else {
		WriterResponse(w, 2, "充电桩不存在", nil)
	}
}
func EndLightFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	no, _ := strconv.Atoi(r.PathParam("no"))
	fmt.Printf("EndLightFunc:%v,%v", uid, no)
	pole := Logic.GetPoleByNo(no)
	if pole != nil {
		response, err := http.Get(fmt.Sprintf("%s/stoplight/%s", DeviceUrl, pole.UUID))
		if response != nil {
			defer response.Body.Close()
		}
		if err == nil && response != nil {

			body, _ := ioutil.ReadAll(response.Body)
			var result Result
			json.Unmarshal(body, &result)
			WriterResponse(w, result.ResultCode, result.ResultMessage, nil)

		} else {
			WriterResponse(w, 3, err.Error(), nil)
		}
	} else {
		WriterResponse(w, 2, "充电桩不存在", nil)
	}
}

func GetChargeStatusFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	no, _ := strconv.ParseInt(r.PathParam("orderno"), 10, 64)
	//no, _ := strconv.pa(r.PathParam("orderno"))
	cs := Logic.GetChargeStatus(uid, no)
	WriterResponse(w, 1, "", cs)
	fmt.Printf("GetChargeStatusFunc:%v,%v", uid, no)
}
func GetMyChargesFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))

	//no, _ := strconv.pa(r.PathParam("orderno"))
	cs := Logic.GetMyChargeOrders(uid)
	WriterResponse(w, 1, "", cs)

}
func GetLastChargeFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))

	//no, _ := strconv.pa(r.PathParam("orderno"))
	cs := Logic.GetLastCharge(uid)
	WriterResponse(w, 1, "", cs)
}
