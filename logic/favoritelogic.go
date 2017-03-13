package logic

import (
	"fca/dal"
	"fca/model"
)

func (logic LogicBase) MyFavorite(uid int) (*[]model.Favorite, error) {

	var myFavorite []model.Favorite
	dal.DB.Preload("ElectricPile").Where("uid=?", uid).Find(&myFavorite)
	return &myFavorite, nil
}
func (logic LogicBase) AddFavorite(uid int, electricPileID int) error {

	var myFavorite model.Favorite
	myFavorite.UID = uid
	myFavorite.ElectricPileID = uint(electricPileID)
	dal.DB.NewRecord(&myFavorite)
	dal.DB.Create(&myFavorite)
	return nil
}
func (logic LogicBase) RemoveFavorite(uid int, electricPileID int) error {

	dal.DB.Where("uid = ? and electric_pile_id=?", uid, electricPileID).Delete(&model.Favorite{})
	return nil
}
