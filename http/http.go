package http

import (
	"fmt"
	"log"
	"net/http"

	"fca/libs"
	"fca/logic"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/qor/admin"
)

var Logic logic.LogicBase
var Admin *admin.Admin

type Result struct {
	ResultCode    int
	ResultMessage string
	Data          interface{}
}

func Init() {
	Logic = logic.LogicBase{}
	logic.InitLogic()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	//fmt.Printf("db:%v\n", dal.DB)

	//mux := http.NewServeMux()
	//fmt.Printf("mux:%v\n", mux)
	//Admin.MountTo("/admin", mux)
	InitQorAdmin()

	router, err := rest.MakeRouter(
		rest.Get("/", Index),
		rest.Get("/Hello", Hello),
		rest.Get("/GetAPIs", GetAPIs),
		rest.Get("/GetAPIByName/:Name", GetAPIByName),
		rest.Get("/InitData", InitData),
		rest.Get("/GetUser/:uid", GetUser),
		rest.Post("/RegisterUser", RegisterUser),
		rest.Post("/RegisterByMobile", RegisterByMobile),
		rest.Post("/UserLogin", UserLogin),
		rest.Post("/CreateCarBrand", CreateCarBrand),
		rest.Get("/GetCarBrands", GetCarBrands),
		rest.Get("/GetQuestions", GetQuestions),
		rest.Get("/GetElectricPile/:id", GetElectricPile),
		rest.Get("/GetElectricPileByNo/:No", GetElectricPileByNo),
		rest.Get("/GetElectricPiles", GetElectricPiles),
		rest.Post("/CreateElectricPile", CreateElectricPile),
		rest.Get("/VerifyCode/:keystr/:valuestr", VerifyCode),
		rest.Get("/SendCode/:mobile", SendCode),
		rest.Get("/Pay/:PayType/:uid/:Amount", PayFunc),
		rest.Get("/GetBaseConfig", GetBaseConfigFunc),
		rest.Get("/GetFeedback/:uid", GetFeedbackFunc),
		rest.Post("/CreateFeedback", CreateFeedbackFunc),
		rest.Get("/MyWallet/:uid", MyWalletFunc),
		rest.Post("/UpdateUserInfo", UpdateUserInfoFunc),
		rest.Get("/MyFavorite/:uid", MyFavoriteFunc),
		rest.Get("/AddFavorite/:uid/:ElectricPileID", AddFavoriteFunc),
		rest.Get("/RemoveFavorite/:uid/:ElectricPileID", RemoveFavoriteFunc),
		rest.Get("/StartCharge/:uid/:no/:mins", StartChargeFunc),
		rest.Get("/EndCharge/:uid/:no", EndChargeFunc),
		rest.Get("/StartLight/:uid/:no", StartLightFunc),
		rest.Get("/EndLight/:uid/:no", EndLightFunc),
		rest.Get("/GetChargeStatus/:uid/:orderno", GetChargeStatusFunc),
		rest.Get("/GetMyCharges/:uid", GetMyChargesFunc),
		rest.Get("/GetLastCharge/:uid", GetLastChargeFunc),
		//预约
		rest.Post("/CreateReserve/:uid", CreateReserveFunc),
		rest.Get("/GetMyReserve/:uid", GetMyReserveFunc),
		//评论
		rest.Post("/CreateComment/:uid", CreateCommentFunc),
		rest.Get("/GetComments/:epid", GetCommentsFunc),
		rest.Get("/GetMyMessages/:uid", GetMyMessagesFunc),
		rest.Get("/SetMessageRead/:uid/:msgid", SetMessageReadFunc),
	//rest.Delete("/countries/:code", DeleteCountry),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))
	http.HandleFunc("/verifyimg", GetVerifyImg)
	http.HandleFunc("/admin/login", ExecuteHttp)
	http.HandleFunc("/admin/loginpost", LoginPost)
	Admin.MountTo("/admin", http.DefaultServeMux)
	http.HandleFunc("/callback/aliappcallback", func(w http.ResponseWriter, r *http.Request) {
		aliResult, err := AliAppCallback(w, r)
		if err != nil {
			fmt.Println(err)
			//log.xxx
			return
		}
		selfHandler(aliResult)
	})
	//InitAuth()
	if libs.Debug {
		go log.Fatal(http.ListenAndServe(":8999", nil))
	} else {
		err1 := http.ListenAndServeTLS(":443", "214079740930545.pem", "214079740930545.key", nil)
		fmt.Printf("err1:%v\n", err1)
	}

	/*	router := httprouter.New()
		router.GET("/", Index)
		router.GET("/hello/:name", Hello)
		router.GET("/UserLogin", UserLogin)
		router.POST("RegisterUser", RegisterUser)
		log.Fatal(http.ListenAndServe(":8080", router))*/
}
func Index(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson("Welcome111111111!")

}

func Hello(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson("Welcome111111111!")

}

func WriterResponse(w rest.ResponseWriter, code int, errMsg string, data interface{}) {
	result := Result{}
	result.ResultCode = code
	result.ResultMessage = errMsg
	result.Data = data
	w.WriteJson(result)
}
