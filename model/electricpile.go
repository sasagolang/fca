package model

import "github.com/jinzhu/gorm"

type EPDirectType struct {
	gorm.Model
	TypeName string
}

type EPType struct {
	gorm.Model
	TypeName string
}
type ElectricPile struct {
	gorm.Model
	//EPID         int
	No             string
	Name           string
	Address        string
	Brand          string
	EPDirectType   EPDirectType
	EPDirectTypeID uint
	Latitude       float32
	Longitude      float32
	Interval       float32
	EPType         EPType
	EPTypeID       uint
	Idle           int
	HadCharage     int
	IsBook         int
	Standard       string
	Seq            int
	ServiceFee     int
	ChargeRule     ChargeRule
	ChargeRuleID   uint
	//ModelBase
}
