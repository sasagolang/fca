package logic

import (
	"fca/dal"
	"fca/model"
	"time"

	"github.com/jinzhu/gorm"
)

func (logic LogicBase) CreateChargeOrder(uid int, no int, uuid, epname, epaddres string) *model.ChargeOrder {
	order := new(model.ChargeOrder)
	var u model.User
	dal.DB.Where("UID = ?", uid).First(&u)
	//fmt.Printf("CreateOrder.User:%v\n", u)
	order.UID = uid
	order.NO = time.Now().UnixNano()
	order.Amount = 0
	order.Duration = 0
	order.Status = 0
	order.UUID = uuid
	order.EPName = epname
	order.EPAddress = epaddres
	order.StartTime = time.Now().Unix()
	order.StatusName = "充电中"
	dal.DB.NewRecord(order)
	dal.DB.Create(&order)

	return order
}
func (logic LogicBase) GetChargeStatus(uid int, no int64) *model.ChargeOrder {
	var charge model.ChargeOrder
	dal.DB.Where("UID = ? and no=?", uid, no).First(&charge)
	return &charge
}
func (logic LogicBase) GetLastCharge(uid int) *model.ChargeOrder {
	var charge model.ChargeOrder
	dal.DB.Where("UID = ? and Status=?", uid, 0).First(&charge)
	return &charge
}
func (logic LogicBase) GetMyChargeOrders(uid int) *[]model.ChargeOrder {
	var charges []model.ChargeOrder
	dal.DB.Where("UID = ? ", uid).Find(&charges)
	return &charges
}
func (logic LogicBase) GetKFOrders() *[]model.ChargeOrder {
	var charges []model.ChargeOrder
	dal.DB.Where("Status=? ", 2).Find(&charges)
	return &charges
}
func (logic LogicBase) KF() {
	orders := logic.GetKFOrders()
	if orders != nil {
		for _, v := range *orders {
			tx := dal.DB.Begin()
			var myWallet model.MyWallet
			var order model.ChargeOrder
			if !dal.DB.Where("uid=?", v.UID).First(&myWallet).RecordNotFound() {
				dal.DB.Model(&myWallet).Update("Balance", gorm.Expr("Balance - ?", v.Amount))
				dal.DB.Model(&order).Where("ID = ?", v.ID).Updates(map[string]interface{}{"Status": 3, "StatusName": "扣款完成"})
			}

			tx.Commit()
		}
	}
}
