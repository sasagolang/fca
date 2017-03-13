package http

import (
	"fmt"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func MyFavoriteFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	myFavorite, err := Logic.MyFavorite(uid)
	if err != nil {
		fmt.Printf("MyFavoriteFunc error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", myFavorite)
}
func AddFavoriteFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	electricPileID, _ := strconv.Atoi(r.PathParam("ElectricPileID"))
	err := Logic.AddFavorite(uid, electricPileID)
	if err != nil {
		fmt.Printf("AddFavoriteFunc error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", "")
}
func RemoveFavoriteFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	electricPileID, _ := strconv.Atoi(r.PathParam("ElectricPileID"))
	err := Logic.RemoveFavorite(uid, electricPileID)
	if err != nil {
		fmt.Printf("RemoveFavoriteFunc error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", "")
}
