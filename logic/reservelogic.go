package logic

import (
	"fca/dal"
	"fca/model"
	"fmt"
)

func (logic LogicBase) CreateReserve(uid, epid int, startTime, endTime int64) *model.Reserve {
	reserve := new(model.Reserve)
	var u model.User
	dal.DB.Where("UID = ?", uid).First(&u)
	fmt.Printf("CreateReserve.User:%v\n", u)
	reserve.User = u
	reserve.UID = uid
	reserve.ElectricPileID = uint(epid)
	reserve.StartTime = startTime
	reserve.EndTime = endTime
	reserve.Status = 1
	b := dal.DB.NewRecord(reserve)
	dal.DB.Create(&reserve)
	fmt.Printf("CreateReserve:%v\n", b)
	return reserve
}
func (logic LogicBase) GetMyReserve(uid int) []model.MyReserveResponse {
	var reservers []model.Reserve
	var tmps []model.MyReserveResponse
	dal.DB.Preload("ElectricPile").Where("UID = ?", uid).Find(&reservers)
	for _, v := range reservers {
		a := model.MyReserveResponse{}
		a.EndTime = v.EndTime
		a.UserID = v.UserID
		a.UID = v.UID
		a.StartTime = v.StartTime
		a.Status = v.Status
		a.EPName = v.ElectricPile.Name
		a.EPAddress = v.ElectricPile.Address
		a.EPNo = v.ElectricPile.No

		tmps = append(tmps, a)
	}
	return tmps
}
