package dal

import (
	"fca/libs"
	"fca/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type BaseDao struct {
}

func InitDb() {
	libs.Log.Info("InitDb")
	var err error
	//DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(localhost:2954)/%v?parseTime=True&loc=Local", "project", "project", "fc"))
	DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v?parseTime=True&loc=Local", "project", "project", "fc"))

	//db, err := sql.Open("mysql", "project:project@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Printf("%v\n", err)
	} else {
		DB.LogMode(true)
	}
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.QuestionType{}, &model.Question{})
	DB.AutoMigrate(&model.CarBrand{}, &model.CarSet{})
	DB.AutoMigrate(&model.EPDirectType{}, &model.EPType{}, &model.ElectricPile{})
	DB.AutoMigrate(&model.Order{})
	DB.AutoMigrate(&model.MyWallet{}, &model.MyWalletLog{})
	DB.AutoMigrate(&model.BaseConfig{})
	DB.AutoMigrate(&model.Feedback{})
	DB.AutoMigrate(&model.Favorite{})
	DB.AutoMigrate(&model.ChargeFee{})
	DB.AutoMigrate(&model.ChargeRule{})
	DB.AutoMigrate(&model.Reserve{})
	DB.AutoMigrate(&model.SMSInfo{})
	DB.AutoMigrate(&model.Pole{})
	DB.AutoMigrate(&model.Comment{})
	DB.AutoMigrate(&model.ChargeOrder{})
	DB.AutoMigrate(&model.Message{})
	//MDbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	//MDbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	//MDbMap.AddTableWithName(model.User{}, "Users").SetKeys(true, "UID")
	//MDbMap.AddTableWithName(model.CarBrand{}, "CarBrands").SetKeys(true, "BID")
	//MDbMap.AddTableWithName(model.CarSet{}, "CarSets").SetKeys(true, "SID")
	//MDbMap.AddTableWithName(model.QuestionType{}, "QuestionTypes").SetKeys(true, "QTID")
	//MDbMap.AddTableWithName(model.Question{}, "Questions").SetKeys(true, "QID")
	//	MDbMap.AddTableWithName(model.ElectricPile{}, "ElectricPiles").SetKeys(true, "EPID")
}
