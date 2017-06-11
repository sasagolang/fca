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
	APIMap["GetElectricPiles"] = model.API{Name: "GetElectricPiles", URL: "/GetElectricPiles?searchkey=上海&mylat=1.0123&mylng=2&distinct=100000&epTypeID=1", Request: nil, Response: []model.ElectricPile{}}
	APIMap["GetElectricPileByNo"] = model.API{Name: "GetElectricPileByNo", URL: "/GetElectricPileByNo/{No}", Request: nil, Response: model.ElectricPile{}}
	APIMap["SendCode"] = model.API{Name: "SendCode", URL: "/SendCode/:mobile", Request: nil, Response: model.VerrifyImg{}}

	APIMap["CreateElectricPile"] = model.API{Name: "CreateElectricPile", URL: "/CreateElectricPile", Request: model.ElectricPile{}, Response: nil}
	APIMap["RegisterByMobile"] = model.API{Name: "RegisterByMobile", URL: "/RegisterByMobile", Request: model.RegisterUserRequest{}, Response: nil}
	APIMap["GetBaseConfig"] = model.API{Name: "GetBaseConfig", URL: "/GetBaseConfig", Request: nil, Response: model.BaseConfig{}}
	APIMap["Pay"] = model.API{Name: "Pay", URL: "/Pay/:PayType/:uid/:Amount", Request: nil, Response: nil}

	APIMap["GetFeedback"] = model.API{Name: "GetFeedback", URL: "GetFeedback/:uid", Request: nil, Response: model.Feedback{}}
	APIMap["CreateFeedback"] = model.API{Name: "CreateFeedback", URL: "CreateFeedback", Request: model.CreateFeedbackRequest{}, Response: model.Feedback{}}
	APIMap["MyWallet"] = model.API{Name: "MyWallet", URL: "/MyWallet/:uid", Request: nil, Response: model.MyWallet{}}
	APIMap["UpdateUserInfo"] = model.API{Name: "UpdateUserInfo", URL: "UpdateUserInfo", Request: model.UpdateUserInfoRequest{}, Response: nil}
	APIMap["MyFavorite"] = model.API{Name: "MyFavorite", URL: "/MyFavorite/:uid", Request: nil, Response: model.Favorite{}}
	APIMap["AddFavorite"] = model.API{Name: "AddFavorite", URL: "/AddFavorite/:uid/:ElectricPileID", Request: nil, Response: nil}
	APIMap["RemoveFavorite"] = model.API{Name: "RemoveFavorite", URL: "/RemoveFavorite/:uid/:ElectricPileID", Request: nil, Response: nil}
	APIMap["CreateReserve"] = model.API{Name: "CreateReserve", URL: "/CreateReserve/:uid", Request: model.CreateReserveRequest{}, Response: nil}
	APIMap["GetMyReserve"] = model.API{Name: "GetMyReserve", URL: "/GetMyReserve/:uid", Request: nil, Response: []model.MyReserveResponse{}}
	APIMap["CreateComment"] = model.API{Name: "CreateComment", URL: "/CreateComment/:uid", Request: model.CreateCommentRequest{}, Response: nil}
	APIMap["GetComments"] = model.API{Name: "GetComments", URL: "/GetComments/:epid", Request: nil, Response: []model.GetCommentResponse{}}
	APIMap["GetLastCharge"] = model.API{Name: "GetLastCharge", URL: "/GetLastCharge/:uid/", Request: nil, Response: model.ChargeOrder{}}
	APIMap["GetMyMessages"] = model.API{Name: "GetMyMessages", URL: "/GetMyMessages/:uid", Request: nil, Response: []model.Message{}}
	APIMap["SetMessageRead"] = model.API{Name: "SetMessageRead", URL: "/SetMessageRead/:uid/:msgid", Request: nil, Response: nil}
	APIMap["StartCharge"] = model.API{Name: "StartCharge", URL: "/StartCharge/:uid/:no/:mins", Request: nil, Response: nil}
	APIMap["EndCharge"] = model.API{Name: "EndCharge", URL: "/EndCharge/:uid/:no", Request: nil, Response: nil}
	APIMap["StartLight"] = model.API{Name: "StartLight", URL: "/StartLight/:uid/:no", Request: nil, Response: nil}

	APIMap["EndLight"] = model.API{Name: "EndLight", URL: "/EndLight/:uid/:no", Request: nil, Response: nil}
	APIMap["GetChargeStatus"] = model.API{Name: "GetChargeStatus", URL: "/GetChargeStatus/:uid/:orderno", Request: nil, Response: model.ChargeOrder{}}
	APIMap["GetMyCharges"] = model.API{Name: "GetMyCharges", URL: "/GetMyCharges/:uid/", Request: nil, Response: []model.ChargeOrder{}}
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
