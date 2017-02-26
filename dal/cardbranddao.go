package dal

import (
	"fca/model"
)

func (dao BaseDao) AddBrand(brand *model.CarBrand) (err error) {

	//err = MDbMap.Insert(brand)
	return

}
func (dao BaseDao) AddCarSet(set *model.CarSet) (err error) {

	//err = MDbMap.Insert(set)
	return

}
func (dao BaseDao) GetBrands(status int) (brands []model.CarBrand, err error) {

	//	_, err = MDbMap.Select(&brands, "select * from CarBrands order by seq")
	return
}
func (dao BaseDao) GetSets(status int) (sets []model.CarSet, err error) {
	//_, err = MDbMap.Select(&sets, "select * from CarSets order by seq")
	return
}
