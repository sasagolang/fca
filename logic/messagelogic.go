package logic

import (
	"fca/dal"
	"fca/model"
)

func (logic LogicBase) GetMyMessages(uid int) *[]model.Message {
	var msgs []model.Message

	dal.DB.Where("uid=? and status>0", uid).Find(&msgs)
	return &msgs
}
func (logic LogicBase) SetMessageRead(id int) {
	msg := model.Message{}
	dal.DB.Model(&msg).Where("id=?", id).Update("Status", 2)

}
