package logic

import (
	"fca/dal"
	"fca/model"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func (logic LogicBase) CreateOrder(uid int, orderAmout int64, fromSource string) *model.Order {
	order := new(model.Order)
	var u model.User
	dal.DB.Where("UID = ?", uid).First(&u)
	fmt.Printf("CreateOrder.User:%v\n", u)
	order.User = u
	order.OrderNo = time.Now().UnixNano()
	order.PaymentAmount = orderAmout
	order.OrderType = "充值"
	order.FromSource = fromSource
	b := dal.DB.NewRecord(order)
	dal.DB.Create(&order)
	fmt.Printf("CreateOrder:%v\n", b)
	return order
}
func (logic LogicBase) DepositFunc(uid int, depositAmout int64, fromSource string, orderno int64) {
	//检查订单
	var order model.Order
	//var user *model.User
	//dal.DB.Model(&order).Association("users").First(&user)

	if dal.DB.Preload("User").Where("order_no=? and state=''", orderno).First(&order).RecordNotFound() {

	} else {
		uid = order.User.UID
		fmt.Printf("DepositFunc.&order:%v\n", &order)
		var user model.User

		dal.DB.First(&user, order.UserID)
		fmt.Printf("DepositFunc.user:%v\n", &user)
		if false && int(user.UID) != uid {
			fmt.Printf("DepositFunc.order1:%v\n", &order)
		} else {
			//order.SetState("已付款")
			//order.State = "dfdfd"
			//order.PayFrom = fromSource
			//db.Model(&user).Updates(User{Name: "hello", Age: 18})
			fmt.Printf("DepositFunc.order2:%v\n", order)
			dal.DB.Model(&order).Updates(map[string]interface{}{"State": "已付款", "PayFrom": fromSource})

		}
	}
	fmt.Printf("DepositFunc.RecordNotFound:%v\n", order)
	var myWallet model.MyWallet
	b := dal.DB.Where("uid=?", uid).First(&myWallet).RecordNotFound()

	if b {
		myWallet = *new(model.MyWallet)
		myWallet.UID = uid
		myWallet.Balance = depositAmout
		myWallet.FreezeBalance = 0
		r := dal.DB.NewRecord(&myWallet)
		fmt.Printf("DepositFunc1:%v,%v\n", myWallet, r)
		dal.DB.Create(&myWallet)
	} else {
		dal.DB.Model(&myWallet).Update("Balance", gorm.Expr("Balance + ?", depositAmout))
	}

	//记录日志
	walletLog := new(model.MyWalletLog)
	walletLog.UID = uid
	walletLog.Type = "充值"
	walletLog.Info = "充值:" + strconv.FormatFloat(float64(depositAmout), 'f', -4, 32) + "订单号：" + strconv.FormatInt(orderno, 10) + "日期：" + time.Now().Format("2006-01-02 15:04:05")
	dal.DB.NewRecord(walletLog)
	dal.DB.Create(walletLog)
	//fmt.Printf("DepositFunc:%v,%v\n", myWallet, b)
	//if myWallet != nil {

	//}
}
