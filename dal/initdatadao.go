package dal

import (
	"fca/model"
	"time"
)

func (dao BaseDao) InitData() {
	/*dao.InitBrandType()
	dao.InitBrandSet()*/
}

func (dao BaseDao) InitBrandType() {

	brands := []*model.CarBrand{}
	brands = append(brands, &model.CarBrand{BrandName: "奥迪", Logo: "aodi.png", Seq: 1})
	brands = append(brands, &model.CarBrand{BrandName: "宝马", Logo: "baoma.png", Seq: 1})
	brands = append(brands, &model.CarBrand{BrandName: "北京新能源", Logo: "beifangxinnengyuan.png", Seq: 1})
	brands = append(brands, &model.CarBrand{BrandName: "奔驰", Logo: "benchi.png", Seq: 1})
	brands = append(brands, &model.CarBrand{BrandName: "比亚迪", Logo: "biyadi.png", Seq: 1})
	i := 1
	for _, v := range brands {
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		v.Seq = i

		i++
	}
	for _, v := range brands {
		dao.AddBrand(v)
	}
}
func (dao BaseDao) InitBrandSet() {
	carSets := []*model.CarSet{}
	carSets = append(carSets, &model.CarSet{SetName: "A1"})
	carSets = append(carSets, &model.CarSet{SetName: "其他"})
	carSets = append(carSets, &model.CarSet{SetName: "I3"})
	carSets = append(carSets, &model.CarSet{SetName: "3系"})
	carSets = append(carSets, &model.CarSet{SetName: "其他"})
	carSets = append(carSets, &model.CarSet{SetName: "E150EV"})
	carSets = append(carSets, &model.CarSet{SetName: "其他"})
	i := 1
	for _, v := range carSets {
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		v.Seq = i

		i++
	}
	for _, v := range carSets {
		dao.AddCarSet(v)
	}

}
