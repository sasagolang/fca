package dal

import "fca/model"

func (dao BaseDao) GetQuestionTypes(status int) (questionTypes []model.QuestionType, err error) {

	//	_, err = MDbMap.Select(&questionTypes, "select * from QuestionTypes order by seq")
	return
}
func (dao BaseDao) GetQuestions(status int) (questions []model.Question, err error) {
	//	_, err = MDbMap.Select(&questions, "select * from Questions order by seq")
	return
}
