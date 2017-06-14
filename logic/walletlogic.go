package logic

import (
	"fca/dal"
	"fca/model"

	"github.com/jinzhu/gorm"
)

func (logic LogicBase) MyWallet(uid int) (*model.MyWallet, error) {

	var myWallet model.MyWallet
	if dal.DB.Where("uid=?", uid).First(&myWallet).RecordNotFound() {
		myWallet = model.MyWallet{}

	}
	return &myWallet, nil
}

func (logic LogicBase) ChargeFee(uid int, balance int) bool {
	var myWallet model.MyWallet
	if !dal.DB.Where("uid=?", uid).First(&myWallet).RecordNotFound() {
		dal.DB.Model(&myWallet).Update("Balance", gorm.Expr("Balance - ?", balance))
		return true
	}
	return false
}
