package http

import (
	"fmt"

	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func CreateElectricPile(w rest.ResponseWriter, r *rest.Request) {
	/*	ep := model.ElectricPile{}
		err := r.DecodeJsonPayload(&ep)
		err = Logic.CreateElectricPile(ep.EPName, ep.Address, ep.EPBrand, ep.EPDirectType, ep.Latitude, ep.Longitude, ep.Interval, ep.EPType, ep.Idle)
		if err != nil {
			fmt.Printf("CreateElectricPile error(%v)\n", err)
		}
		WriterResponse(w, 1, err)
	*/
}
func GetElectricPile(w rest.ResponseWriter, r *rest.Request) {
	id, _ := strconv.Atoi(r.PathParam("id"))
	ep, err := Logic.GetElectricPileByID(id)
	if err != nil {
		fmt.Printf("GetElectricPile error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", ep)
}
func GetElectricPileByNo(w rest.ResponseWriter, r *rest.Request) {
	no := r.PathParam("No")
	ep, err := Logic.GetElectricPileByNo(no)
	if err != nil {
		fmt.Printf("GetElectricPileByNo error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", ep)
}
func GetElectricPiles(w rest.ResponseWriter, r *rest.Request) {
	searchkey := r.Request.FormValue("searchkey")
	mylat, _ := strconv.ParseFloat(r.Request.FormValue("mylat"), 64)
	mylng, _ := strconv.ParseFloat(r.Request.FormValue("mylng"), 64)
	distinct, _ := strconv.ParseFloat(r.Request.FormValue("distinct"), 64)
	//fmt.Printf("searchkey:%v", searchkey)
	eps, err := Logic.GetElectricPiles(searchkey, mylat, mylng, distinct)
	if err != nil {
		fmt.Printf("GetElectricPiles error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", eps)
}
