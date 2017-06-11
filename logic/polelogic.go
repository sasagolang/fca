package logic

import (
	"fca/dal"
	"fca/model"
)

func (logic LogicBase) GetPole(uuid string) *model.Pole {
	var pole model.Pole
	dal.DB.Preload("ElectricPile").Where("UUID=?", uuid).First(&pole)
	return &pole
}
func (logic LogicBase) UpdateCharge(uuid, chargeState string, energy int32, amount float32, duration int32) {

	var cs model.ChargeOrder
	if !dal.DB.Where("UUID=?", uuid).Last(&cs).RecordNotFound() {
		dal.DB.Model(&cs).Updates(map[string]interface{}{"State": chargeState, "Duration": duration, "Energy": energy, "Amount": amount})
	}

}
func (logic LogicBase) EndCharge(uuid string) {

	var cs model.ChargeOrder
	if !dal.DB.Where("UUID=?", uuid).Last(&cs).RecordNotFound() {
		dal.DB.Model(&cs).Updates(map[string]interface{}{"Status": 2, "StatusName": "待扣费"})
	}

}
func (logic LogicBase) EndChargeOrderno(uuid string, orderno int64) {

	var cs model.ChargeOrder
	if !dal.DB.Where("UUID=? and no=?", uuid, orderno).Last(&cs).RecordNotFound() {
		dal.DB.Model(&cs).Updates(map[string]interface{}{"Status": 2, "StatusName": "待扣费"})
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
