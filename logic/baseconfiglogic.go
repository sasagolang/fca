package logic

import (
	"fca/dal"
	"fca/model"
)

func (logic LogicBase) GetBaseConfig() *model.BaseConfig {
	var bc model.BaseConfig
	dal.DB.First(&bc)
	return &bc
}
