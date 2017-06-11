package http

import (
	"fca/model"
	"strconv"

	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
)

func CreateReserveFunc(w rest.ResponseWriter, r *rest.Request) {
	//uid, epid int, startTime, endTime time.Time
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	request := model.CreateReserveRequest{}
	err := r.DecodeJsonPayload(&request)
	if err != nil {
		fmt.Sprintf("CreateReserveFunc.err:%v\n", err)
	}
	Logic.CreateReserve(uid, request.ElectricPileID, request.StartTime, request.EndTime)
	WriterResponse(w, 1, "", "")
}
func GetMyReserveFunc(w rest.ResponseWriter, r *rest.Request) {
	// uid int
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	reserves := Logic.GetMyReserve(uid)
	WriterResponse(w, 1, "", reserves)
}
