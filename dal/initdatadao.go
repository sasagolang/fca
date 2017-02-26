package dal

func (dao BaseDao) InitData() {
	/*dao.InitBrandType()
	dao.InitBrandSet()*/
}

func (dao BaseDao) InitBrandType() {

	/*brands := []*model.CarBrand{}
		brands = append(brands, &model.CarBrand{BrandName: "奥迪", Logo: "aodi.png", Seq: 1})
		brands = append(brands, &model.CarBrand{BrandName: "宝马", Logo: "baoma.png", Seq: 1})
		brands = append(brands, &model.CarBrand{BrandName: "北京新能源", Logo: "beifangxinnengyuan.png", Seq: 1})
		brands = append(brands, &model.CarBrand{BrandName: "奔驰", Logo: "benchi.png", Seq: 1})
		brands = append(brands, &model.CarBrand{BrandName: "比亚迪", Logo: "biyadi.png", Seq: 1})
		i := 1
		for _, v := range brands {
			v.CreateTime = time.Now().UnixNano()
			v.LastUpdateTime = time.Now().UnixNano()
			v.Seq = i
			v.Status = 1
			i++
		}
		for _, v := range brands {
			dao.AddBrand(v)
		}
	}
	func (dao BaseDao) InitBrandSet() {
		carSets := []*model.CarSet{}
		carSets = append(carSets, &model.CarSet{SetName: "A1", BID: 1})
		carSets = append(carSets, &model.CarSet{SetName: "其他", BID: 1})
		carSets = append(carSets, &model.CarSet{SetName: "I3", BID: 2})
		carSets = append(carSets, &model.CarSet{SetName: "3系", BID: 2})
		carSets = append(carSets, &model.CarSet{SetName: "其他", BID: 2})
		carSets = append(carSets, &model.CarSet{SetName: "E150EV", BID: 3})
		carSets = append(carSets, &model.CarSet{SetName: "其他", BID: 3})
		i := 1
		for _, v := range carSets {
			v.CreateTime = time.Now().UnixNano()
			v.LastUpdateTime = time.Now().UnixNano()
			v.Seq = i
			v.Status = 1
			i++
		}
		for _, v := range carSets {
			dao.AddCarSet(v)
		}
	*/
}
