package http

import (
	"fmt"
	"log"
	"net/http"

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
		rest.Get("/GetVerifyImgData/:width/:height", GetVerifyImgData),
		rest.Get("/Pay/:PayType/:uid/:Amount", PayFunc),
	//rest.Delete("/countries/:code", DeleteCountry),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))
	http.HandleFunc("/verifyimg", GetVerifyImg)
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
	log.Fatal(http.ListenAndServe(":8999", nil))
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
