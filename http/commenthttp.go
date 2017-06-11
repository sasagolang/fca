package http

import (
	"fca/model"
	"fmt"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func CreateCommentFunc(w rest.ResponseWriter, r *rest.Request) {
	//uid, epid int, startTime, endTime time.Time
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	request := model.CreateCommentRequest{}
	err := r.DecodeJsonPayload(&request)
	if err != nil {
		fmt.Sprintf("CreateCommentFunc.err:%v\n", err)
	}
	Logic.CreateComment(uid, request.ElectricPileID, request.Content, request.Title)
	WriterResponse(w, 1, "", "")
}
func GetCommentsFunc(w rest.ResponseWriter, r *rest.Request) {
	// uid int
	epid, _ := strconv.Atoi(r.PathParam("epid"))
	comments := Logic.GetComments(epid)
	WriterResponse(w, 1, "", comments)
}
