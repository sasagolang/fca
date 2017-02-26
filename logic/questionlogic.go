package logic

import (
	"fca/dal"
	"fca/model"
)

func (logic LogicBase) GetQuestions() (questions []model.QuestionResponse, err error) {
	//var carBrands []model.CarBrandResponse
	var qts []model.QuestionType
	var qs []model.Question
	dal.DB.Find(&qts)
	dal.DB.Find(&qs)

	for _, b := range qts {
		qt := model.QuestionResponse{}
		qt.ID = int(b.ID)
		qt.TypeName = b.TypeName
		qt.Seq = b.Seq
		qt.Questions = []model.Question{}
		for _, s := range qs {
			if b.ID == s.QuestionTypeID {
				q := model.Question{}
				q.ID = s.ID
				q.QuestionTypeID = s.QuestionTypeID
				q.Description = s.Description
				qt.Questions = append(qt.Questions, s)
			}
		}
		questions = append(questions, qt)
	}
	return
}
