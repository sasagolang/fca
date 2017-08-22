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
	Admin.AddResource(&model.User{}, &admin.Config{Name: "用户信息"})
	//Admin.AddResource(&model.User{}, &admin.Config{Name: "用户信息"})
	//return
	//users.Name = "用户信息"
	//users.Meta(&admin.Meta{Name: "用户信息"})
	//Admin.AddMenu(&admin.Menu{Name: "客户信息", Link: "/admin/users"})
	//Question
	types := Admin.AddResource(&model.QuestionType{}, &admin.Config{Name: "问题类型"})
	question := Admin.AddResource(&model.Question{}, &admin.Config{Name: "问题"})
	question.Meta(&admin.Meta{Name: "QuestionType", Config: &admin.SelectOneConfig{SelectMode: "select_async", RemoteDataResource: types}})
	//CarBrand
	//question.Meta(&admin.Meta{Name: "QuestionType", Config: &admin.SelectOneConfig{Collection: QuestionType}})
	carBrands := Admin.AddResource(&model.CarBrand{}, &admin.Config{Name: "车型品牌"})
	carSets := Admin.AddResource(&model.CarSet{}, &admin.Config{Name: "车型"})
	carSets.Meta(&admin.Meta{Name: "CarBrand", Config: &admin.SelectOneConfig{SelectMode: "select_async", RemoteDataResource: carBrands}})

	//ElectricPile
	epDirectTypes := Admin.AddResource(&model.EPDirectType{}, &admin.Config{Name: "站点标准"})
	epTypes := Admin.AddResource(&model.EPType{}, &admin.Config{Name: "站点类型"})
	electricPiles := Admin.AddResource(&model.ElectricPile{}, &admin.Config{Name: "站点信息"})
	electricPiles.Meta(&admin.Meta{Name: "EPDirectType", Config: &admin.SelectOneConfig{SelectMode: "select_async", RemoteDataResource: epDirectTypes}})
	electricPiles.Meta(&admin.Meta{Name: "EPType", Config: &admin.SelectOneConfig{SelectMode: "select_async", RemoteDataResource: epTypes}})

	//订单
	//orders := Admin.AddResource(&model.Order{}, &admin.Config{Invisible: true})
	//if orders != nil {
	//}
	//钱包
	wallets := Admin.AddResource(&model.MyWallet{}, &admin.Config{Name: "钱包金额"})
	if wallets != nil {
	}
	walletlogs := Admin.AddResource(&model.MyWalletLog{}, &admin.Config{Name: "钱包变更记录"})
	if walletlogs != nil {
	}
	Admin.AddResource(&model.BaseConfig{}, &admin.Config{Name: "配置信息"})
	Admin.AddResource(&model.Feedback{}, &admin.Config{Name: "反馈信息"})
	Admin.AddResource(&model.Favorite{}, &admin.Config{Name: "收藏信息"})
	Admin.AddResource(&model.ChargeFee{}, &admin.Config{Name: "收费模式"})
	//fees.Meta(&admin.Meta{Name: "StartTime", Type: "date"})
	Admin.AddResource(&model.ChargeRule{}, &admin.Config{Name: "收费规则"})
	Admin.AddResource(&model.Reserve{}, &admin.Config{Name: "预约信息"})
	Admin.AddResource(&model.SMSInfo{}, &admin.Config{Name: "验证码信息"})
	Admin.AddResource(&model.Pole{}, &admin.Config{Name: "电桩信息"})
	Admin.AddResource(&model.Comment{}, &admin.Config{Name: "评论信息"})
	Admin.AddResource(&model.ChargeOrder{}, &admin.Config{Name: "充电订单"})
	Admin.AddResource(&model.Message{}, &admin.Config{Name: "消息记录"})
	Admin.SetAuth(&Auth{})
}
