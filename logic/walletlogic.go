package logic

import (
	"fca/dal"
	"fca/model"
)

func (logic LogicBase) MyWallet(uid int) (*model.MyWallet, error) {

	var myWallet model.MyWallet
	if dal.DB.Where("uid=?", uid).First(&myWallet).RecordNotFound() {
		myWallet = model.MyWallet{}

	}
	return &myWallet, nil
}
