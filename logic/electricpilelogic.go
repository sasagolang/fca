package logic

import (
	"errors"
	"fca/dal"
	"fca/model"
	"math"
)

func (logic LogicBase) CreateElectricPile(epName string, address string, epBrand string, epDirectType int, latitude float32, longitude float32, interval float32, epType int, idle int) (err error) {
	/*
		ep := model.ElectricPile{}
		ep.EPName = epName
		ep.Address = address
		ep.EPBrand = epBrand
		ep.EPDirectType = ep.EPDirectType
		ep.Latitude = latitude
		ep.Longitude = longitude
		ep.Interval = interval
		ep.EPType = epType
		ep.Idle = idle
		ep.Seq = 1
		ep.Status = 1
		ep.CreateTime = time.Now().UnixNano()
		ep.LastUpdateTime = time.Now().UnixNano()
		err = Dao.AddElectricPile(&ep)
	*/
	return
}

func (logic LogicBase) GetElectricPileByID(id int) (ep *model.ElectricPile, err error) {
	var eptmp model.ElectricPile
	if dal.DB.Preload("EPDirectType").Preload("EPType").First(&eptmp, id).RecordNotFound() {
		return nil, errors.New("未查找到")
	}

	return &eptmp, nil
}
func (logic LogicBase) GetElectricPileByNo(no string) (ep *model.ElectricPile, err error) {
	var eptmp model.ElectricPile
	if dal.DB.Preload("EPDirectType").Preload("ChargeRule").Preload("ChargeRule.ChargeFees").Preload("EPType").Where("No=?", no).First(&eptmp).RecordNotFound() {
		return nil, errors.New("未查找到")
	}

	return &eptmp, nil
}
func (logic LogicBase) GetElectricPiles(searchkey string, mylat float64, mylng float64, distinct float64, epTypeID int) (eps []model.ElectricPile, err error) {
	var elpstmp []model.ElectricPile
	if mylat != 0 && mylng != 0 && distinct != 0 {

		//myLat := 1.00
		//myLng := 1.00
		ranges := 180.00 / math.Pi * distinct / 6372.797 //里面的 1 就代表搜索 1km 之内，单位km
		lngR := ranges / math.Cos(mylat*math.Pi/180.0)
		maxLat := mylat + ranges
		minLat := mylat - ranges
		maxLng := mylng + lngR
		minLng := mylng - lngR
		if epTypeID > 0 {
			dal.DB.Preload("EPDirectType").Preload("EPType").Where("Name like ? And ((latitude BETWEEN ? AND ?) AND (longitude BETWEEN ? AND ?)) AND EPTypeID=?", "%"+searchkey+"%", minLat, maxLat, minLng, maxLng, epTypeID).Find(&elpstmp)

		} else {
			dal.DB.Preload("EPDirectType").Preload("EPType").Where("Name like ? And ((latitude BETWEEN ? AND ?) AND (longitude BETWEEN ? AND ?)) ", "%"+searchkey+"%", minLat, maxLat, minLng, maxLng).Find(&elpstmp)

		}
		//SELECT * FROM checkinTable WHERE ((lat BETWEEN ? AND ?) AND (lng BETWEEN ? AND ?))

	} else {
		dal.DB.Preload("EPDirectType").Preload("EPType").Where("Name like ?  ", "%"+searchkey+"%").Find(&elpstmp)

	}
	return elpstmp, nil
}
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6371000.00 // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * radius
}
