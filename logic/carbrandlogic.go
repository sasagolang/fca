package logic

import (
	"fca/dal"
	"fca/model"
)

func (logic LogicBase) CreateCarBrand(logo string, brandName string, seq int) (err error) {
	/*	brand := model.CarBrand{}
		brand.BrandName = brandName
		brand.Logo = logo
		brand.Seq = seq
		brand.CreateTime = time.Now().UnixNano()
		brand.LastUpdateTime = time.Now().UnixNano()
		err = Dao.AddBrand(&brand)*/
	return
}
func (logic LogicBase) GetCarBrands() (carBrands []model.CarBrandResponse, err error) {
	var sets []model.CarSet
	var brands []model.CarBrand
	//dal.DB.Preload("CarBrand").Find(&sets)
	dal.DB.Find(&sets)
	dal.DB.Find(&brands)
	//var carBrands []model.CarBrandResponse

	for _, b := range brands {
		carb := model.CarBrandResponse{}

		carb.BrandName = b.BrandName
		carb.Logo = b.Logo
		carb.Seq = b.Seq
		carb.CarSets = []model.CarSet{}
		for _, s := range sets {
			if b.ID == s.CarBrandID {
				set := model.CarSet{}
				set.ID = s.ID
				set.CarBrandID = s.CarBrandID
				set.SetName = s.SetName
				set.Seq = s.Seq
				carb.CarSets = append(carb.CarSets, set)

			}
		}
		carBrands = append(carBrands, carb)
	}

	return
}
