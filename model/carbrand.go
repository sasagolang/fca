package model

import "github.com/jinzhu/gorm"

type CarBrand struct {
	gorm.Model
	//BID       int
	Logo      string
	BrandName string
	Seq       int
	//ModelBase
}

type CarSet struct {
	gorm.Model
	//SID     int
	//BID     int
	CarBrandID uint
	CarBrand   CarBrand
	SetName    string
	Seq        int
	//ModelBase
}

type CarBrandResponse struct {
	ID        int
	Logo      string
	BrandName string
	Seq       int
	CarSets   []CarSet
}
