package http

import (
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetMyMessagesFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	msgs := Logic.GetMyMessages(uid)
	/*if err != nil {
		fmt.Printf("GetMyMessagesFunc error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}*/
	WriterResponse(w, 1, "", msgs)
}
func SetMessageReadFunc(w rest.ResponseWriter, r *rest.Request) {
	//	uid, _ := strconv.Atoi(r.PathParam("uid"))
	msgid, _ := strconv.Atoi(r.PathParam("msgid"))
	Logic.SetMessageRead(msgid)
	WriterResponse(w, 1, "", nil)
}
