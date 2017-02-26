package dal

import (
	"fca/model"
)

func (dao BaseDao) AddElectricPile(ep *model.ElectricPile) (err error) {
	//	err = MDbMap.Insert(ep)
	return
}
func (dao BaseDao) GetElectricPileByID(id int) (ep *model.ElectricPile, err error) {
	//	obj, err := MDbMap.Get(model.ElectricPile{}, id)
	//	if obj != nil {
	//		ep = obj.(*model.ElectricPile)

	//	}
	return
}
func (dao BaseDao) GetElectricPiles(status int) (eps []model.ElectricPile, err error) {
	//	_, err = MDbMap.Select(&eps, "select * from ElectricPiles order by seq")
	return
}
