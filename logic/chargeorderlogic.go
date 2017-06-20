package logic

import (
	"fca/dal"
	"fca/model"
	"strconv"
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
	order.StatusName = "准备充电中"
	order.PoleNO = no
	dal.DB.NewRecord(order)
	dal.DB.Create(&order)

	return order
}
func (logic LogicBase) GetChargeStatus(uid int, no int64) *model.ChargeOrder {
	var charge model.ChargeOrder
	dal.DB.Where("UID = ? and no=?", uid, no).First(&charge)
	return &charge
}
func (logic LogicBase) GetLastChargeByUUID(uuid string) *model.ChargeOrder {
	var charge model.ChargeOrder
	dal.DB.Where("UUID = ? ", uuid).Last(&charge)
	return &charge
}
func (logic LogicBase) GetLastCharge(uid int) *model.ChargeOrder {
	var charge model.ChargeOrder
	dal.DB.Where("UID = ? and Status=?", uid, 1).Last(&charge)
	return &charge
}
func (logic LogicBase) GetMyChargeOrders(uid int) *[]model.ChargeOrder {
	var charges []model.ChargeOrder
	dal.DB.Where("UID = ? and Status>? ", uid, 0).Order("id desc").Find(&charges)
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
				dal.DB.Model(&myWallet).Update("Balance", gorm.Expr("Balance - ?", v.Amount*100))
				dal.DB.Model(&order).Where("ID = ?", v.ID).Updates(map[string]interface{}{"Status": 3, "StatusName": v.StatusName + ",扣款完成"})
				walletLog := new(model.MyWalletLog)
				walletLog.UID = v.UID
				walletLog.Type = "充电"
				walletLog.Info = "充电:" + strconv.FormatFloat(float64(v.Amount), 'f', -4, 32) + "订单号：" + strconv.FormatInt(v.NO, 10) + "日期：" + time.Now().Format("2006-01-02 15:04:05")
				dal.DB.NewRecord(walletLog)
				dal.DB.Create(walletLog)
			}

			tx.Commit()
		}
	}
}
