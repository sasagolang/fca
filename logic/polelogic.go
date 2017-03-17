package logic

import (
	"fca/model"
	"fca/dal"
)

func (logic LogicBase) GetPol(uuid string)* model.Pole {
	var pole model.Pole
	dal.DB.Preload("ElectricPile").Where("UUID=?", uuid).First(&pole)
	return &pole
}
