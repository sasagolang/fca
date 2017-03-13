package http

import (
	"fca/model"

	"github.com/ant0ine/go-json-rest/rest"

	"fca/libs"
	"fmt"
)

func CreateCarBrand(w rest.ResponseWriter, r *rest.Request) {
	brand := model.CarBrand{}
	err := r.DecodeJsonPayload(&brand)
	err = Logic.CreateCarBrand(brand.Logo, brand.BrandName, 1)
	if err != nil {
		libs.Log.Error(fmt.Sprintf("CreateCarBrand error(%v)\n", err))

	}
	WriterResponse(w, 1, err.Error(), err)
}
func GetCarBrands(w rest.ResponseWriter, r *rest.Request) {
	carBrands, err := Logic.GetCarBrands()
	if err != nil {
		libs.Log.Error(fmt.Sprintf("GetCarBrands error(%v)\n", err))

		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", carBrands)
}
