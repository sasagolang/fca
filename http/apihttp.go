package http

import (
	"fca/model"

	"github.com/ant0ine/go-json-rest/rest"
)

var APIMap map[string]model.API

func APIinit() {
	APIMap = make(map[string]model.API)
	APIMap["Hello"] = model.API{Name: "Hello", URL: "/Hello", Request: nil, Response: nil}
	APIMap["GetUser"] = model.API{Name: "GetUser", URL: "/GetUser/{uid}", Request: nil, Response: model.User{}}
	APIMap["RegisterUser"] = model.API{Name: "RegisterUser", URL: "/RegisterUser", Request: model.User{}, Response: nil}
	APIMap["UserLogin"] = model.API{Name: "UserLogin", URL: "/UserLogin", Request: model.UserLoginRequest{}, Response: model.User{}}
	APIMap["CreateCarBrand"] = model.API{Name: "CreateCarBrand", URL: "/CreateCarBrand", Request: model.CarBrand{}, Response: nil}
	APIMap["GetCarBrands"] = model.API{Name: "GetCarBrands", URL: "/GetCarBrands", Request: nil, Response: []model.CarBrandResponse{}}
	APIMap["GetQuestions"] = model.API{Name: "GetQuestions", URL: "/GetQuestions", Request: nil, Response: []model.QuestionResponse{}}
	APIMap["GetElectricPile"] = model.API{Name: "GetElectricPile", URL: "/GetElectricPile/{id}", Request: nil, Response: model.ElectricPile{}}
	APIMap["GetElectricPiles"] = model.API{Name: "GetElectricPiles", URL: "/GetElectricPiles", Request: nil, Response: []model.ElectricPile{}}
	APIMap["GetElectricPileByNo"] = model.API{Name: "GetElectricPileByNo", URL: "/GetElectricPileByNo/{No}", Request: nil, Response: model.ElectricPile{}}

	APIMap["CreateElectricPile"] = model.API{Name: "CreateElectricPile", URL: "/CreateElectricPile", Request: model.ElectricPile{}, Response: nil}
	APIMap["RegisterByMobile"] = model.API{Name: "RegisterByMobile", URL: "/RegisterByMobile", Request: model.RegisterUserRequest{}, Response: nil}

}

func GetAPIs(w rest.ResponseWriter, r *rest.Request) {
	APIinit()
	lst := []string{}
	for _, key := range APIMap {
		lst = append(lst, key.Name)
	}
	result := Result{}
	result.ResultCode = 1
	result.ResultMessage = ""
	result.Data = lst
	w.WriteJson(result)
}
func GetAPIByName(w rest.ResponseWriter, r *rest.Request) {
	APIinit()
	name := r.PathParam("Name")
	result := Result{}
	result.ResultCode = 1
	result.ResultMessage = ""
	result.Data = APIMap[name]
	w.WriteJson(result)
}
func InitData(w rest.ResponseWriter, r *rest.Request) {
	Logic.InitData()
	result := Result{}
	result.ResultCode = 1
	result.ResultMessage = ""
	result.Data = ""
	w.WriteJson(result)
}
