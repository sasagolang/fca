package logic

import (
	"fca/dal"
	"fca/model"
	"fmt"
	"time"
)

func (logic LogicBase) CreateComment(uid, epid int, content, title string) *model.Comment {
	comment := new(model.Comment)
	var u model.User
	dal.DB.Where("UID = ?", uid).First(&u)
	fmt.Printf("CreateComment.User:%v\n", u)
	comment.User = u
	comment.UID = uid
	comment.ElectricPileID = uint(epid)

	comment.Status = 1
	comment.Content = content
	comment.Title = title
	comment.CreateTime = time.Now().Unix()
	b := dal.DB.NewRecord(comment)
	dal.DB.Create(&comment)
	fmt.Printf("CreateComment:%v\n", b)
	return comment
}
func (logic LogicBase) GetComments(epid int) []model.GetCommentResponse {
	var comments []model.Comment
	var tmps []model.GetCommentResponse
	dal.DB.Preload("User").Preload("ElectricPile").Where("electric_pile_id = ?", epid).Order("id desc").Find(&comments)
	for _, v := range comments {
		a := model.GetCommentResponse{}
		a.UID = v.UID
		a.Name = v.User.Name
		a.NickName = v.User.NickName
		a.HeadImg = v.User.HeadImg
		a.Content = v.Content
		a.CreateTime = v.CreateTime
		tmps = append(tmps, a)
	}
	return tmps
}
