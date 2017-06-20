package logic

import (
	"encoding/json"
	"fca/dal"
	"fca/model"
	"fmt"
	"io/ioutil"
	"net/http"
)

var DeviceUrl = "http://127.0.0.1:8997"

func (logic LogicBase) GetPole(uuid string) *model.Pole {
	var pole model.Pole
	dal.DB.Preload("ElectricPile").Where("UUID=?", uuid).First(&pole)
	return &pole
}
func (logic LogicBase) UpdatePoleStatus(uuid string, status int32, statusName string) {
	var pole model.Pole
	if !dal.DB.Where("UUID=? ", uuid).Last(&pole).RecordNotFound() {

		dal.DB.Model(&pole).Updates(map[string]interface{}{"Status": status, "StatusName": statusName})
	}

}
func (logic LogicBase) UpdateCharge(uuid, chargeState string, energy int32, amount float32, duration int32) {

	var cs model.ChargeOrder
	if !dal.DB.Where("UUID=?", uuid).Last(&cs).RecordNotFound() {
		if cs.Status >= 2 {
			return
		}
		wallet, _ := logic.MyWallet(cs.UID)

		var statusname string
		var status int
		status = 1
		if float32(wallet.Balance)/100.00 <= amount {
			statusname = "金额不足，停止充电"
			status = 2
			logic.EndChargeTCP(uuid, cs.NO)
		} else {
			switch chargeState {
			case "0":
				statusname = "充电异常"
				status = 2
			case "1":
				statusname = "充电中"
			case "2":
				statusname = "已充满,待扣费"
				status = 2
			case "101":
				statusname = "设备断网"
				status = 2
			case "102":
				statusname = "充电结束"
				status = 2
			}
		}
		fmt.Println("UpdateCharge:%v\n", cs)
		if cs.Status == 2 {
			statusname = cs.StatusName
		}
		if chargeState != "101" && chargeState != "102" {
			dal.DB.Model(&cs).Updates(map[string]interface{}{"State": chargeState, "Status": status, "StatusName": statusname, "Duration": duration, "Energy": energy, "Amount": amount})
		} else {
			dal.DB.Model(&cs).Updates(map[string]interface{}{"State": chargeState, "Status": status, "StatusName": statusname})

		}
	}

}
func (logic LogicBase) EndCharge(uuid string) {

	var cs model.ChargeOrder
	if !dal.DB.Where("UUID=?", uuid).Last(&cs).RecordNotFound() {
		if cs.Status < 2 {
			dal.DB.Model(&cs).Updates(map[string]interface{}{"Status": 2, "StatusName": "待扣费"})
		}
	}

}
func (logic LogicBase) EndChargeOrdernoAndStatus(uuid string, orderno int64, status int) {

	var cs model.ChargeOrder
	if !dal.DB.Where("UUID=? and no=?", uuid, orderno).Last(&cs).RecordNotFound() {
		if cs.Status < 2 {
			dal.DB.Model(&cs).Updates(map[string]interface{}{"Status": status, "StatusName": "待扣费"})
		}
	}

}
func (logic LogicBase) GetPoleByNo(no int) *model.Pole {
	var pole model.Pole
	dal.DB.Preload("ElectricPile").Where("no=?", no).First(&pole)
	return &pole
}

func (logic LogicBase) GetChargeOrderByNo(no int64) *model.ChargeOrder {
	var co model.ChargeOrder
	dal.DB.Where("no=?", no).First(&co)
	return &co
}

type ResultHttp struct {
	ResultCode    int
	ResultMessage string
	Data          interface{}
}

func (logic LogicBase) EndChargeTCP(uuid string, orderno int64) (resultCode int, err1 string) {

	corder := logic.GetLastChargeByUUID(uuid)
	if corder != nil && corder.NO != orderno {
		logic.EndChargeOrdernoAndStatus(uuid, orderno, 2)
		return 1, ""
	}
	logic.EndChargeOrdernoAndStatus(uuid, orderno, 1)
	response, err := http.Get(fmt.Sprintf("%s/stop/%s", DeviceUrl, uuid))
	if response != nil {
		defer response.Body.Close()
	}
	if err == nil && response != nil {

		body, _ := ioutil.ReadAll(response.Body)
		var result ResultHttp
		json.Unmarshal(body, &result)
		//	Logic.EndCharge(pole.UUID)

		resultCode = result.ResultCode
		err1 = result.ResultMessage
		return
		//WriterResponse(w, result.ResultCode, result.ResultMessage, nil)

	} else {
		return 3, err.Error()
		//	WriterResponse(w, 3, err.Error(), nil)
	}
}
