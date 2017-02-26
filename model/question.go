package model

import "github.com/jinzhu/gorm"

type QuestionType struct {
	//QTID  int
	gorm.Model
	TypeName string
	Seq      int
	//Questions []Question `gorm:"ForeignKey:QT"`
	//ModelBase
}
type Question struct {
	gorm.Model
	//QID   int
	QuestionType   QuestionType
	Description    string
	Seq            int
	QuestionTypeID uint
	//ModelBase
}
type QuestionResponse struct {
	ID        int
	TypeName  string
	Seq       int
	Questions []Question
}
