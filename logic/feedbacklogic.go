package logic

import (
	"fca/dal"
	"fca/model"
)

func (logic LogicBase) GetFeedback(uid int) (*[]model.Feedback, error) {
	var feedbacks []model.Feedback
	dal.DB.Where("uid=?", uid).Find(&feedbacks)
	return &feedbacks, nil

}

func (logic LogicBase) CreateFeedback(uid int, feedType int, description string, photo string, questionType string, electricPileName string) (*model.Feedback, error) {

	fb := model.Feedback{}
	fb.FeedType = feedType
	fb.UID = uid
	fb.Description = description
	fb.Photo = photo
	fb.Status = 1
	fb.ElectricPileName = electricPileName
	fb.QuestionType = questionType
	dal.DB.NewRecord(&fb)
	dal.DB.Create(&fb)
	return &fb, nil
}
