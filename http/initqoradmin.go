package http

import (
	"fca/dal"
	"fca/model"

	"github.com/qor/admin"
	"github.com/qor/qor"
)

func InitQorAdmin() {
	Admin = admin.New(&qor.Config{DB: dal.DB})
	Admin.SetSiteName("大源充电后台管理系统")
	//, &admin.Config{Name: "客户信息", Link: "/User"}
	Admin.AddResource(&model.User{})
	//Admin.AddMenu(&admin.Menu{Name: "客户信息", Link: "/admin/users"})
	//Question
	types := Admin.AddResource(&model.QuestionType{})
	question := Admin.AddResource(&model.Question{})
	question.Meta(&admin.Meta{Name: "QuestionType", Config: &admin.SelectOneConfig{SelectMode: "select_async", RemoteDataResource: types}})
	//CarBrand
	//question.Meta(&admin.Meta{Name: "QuestionType", Config: &admin.SelectOneConfig{Collection: QuestionType}})
	carBrands := Admin.AddResource(&model.CarBrand{})
	carSets := Admin.AddResource(&model.CarSet{})
	carSets.Meta(&admin.Meta{Name: "CarBrand", Config: &admin.SelectOneConfig{SelectMode: "select_async", RemoteDataResource: carBrands}})

	//ElectricPile
	epDirectTypes := Admin.AddResource(&model.EPDirectType{})
	epTypes := Admin.AddResource(&model.EPType{})
	electricPiles := Admin.AddResource(&model.ElectricPile{})
	electricPiles.Meta(&admin.Meta{Name: "EPDirectType", Config: &admin.SelectOneConfig{SelectMode: "select_async", RemoteDataResource: epDirectTypes}})
	electricPiles.Meta(&admin.Meta{Name: "EPType", Config: &admin.SelectOneConfig{SelectMode: "select_async", RemoteDataResource: epTypes}})

	//订单
	orders := Admin.AddResource(&model.Order{})
	if orders != nil {
	}
	//钱包
	wallets := Admin.AddResource(&model.MyWallet{})
	if wallets != nil {
	}
	walletlogs := Admin.AddResource(&model.MyWalletLog{})
	if walletlogs != nil {
	}
	Admin.AddResource(&model.BaseConfig{})
	Admin.AddResource(&model.Feedback{})
	Admin.AddResource(&model.Favorite{})
	Admin.AddResource(&model.ChargeFee{})
	//fees.Meta(&admin.Meta{Name: "StartTime", Type: "date"})
	Admin.AddResource(&model.ChargeRule{})
	Admin.AddResource(&model.Reserve{})
	Admin.AddResource(&model.SMSInfo{})
}
