package model

import "github.com/jinzhu/gorm"

type Feedback struct {
	gorm.Model
	FeedType         int
	UID              int
	Description      string
	Photo            string
	Status           int
	Reply            string
	QuestionType     string
	ElectricPileName string
}

type CreateFeedbackRequest struct {
	UID              int
	FeedType         int
	Description      string
	Photo            string
	PhotoSuffix      string
	QuestionType     string
	ElectricPileName string
}
