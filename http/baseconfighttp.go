package http

import "github.com/ant0ine/go-json-rest/rest"

func GetBaseConfigFunc(w rest.ResponseWriter, r *rest.Request) {
	bc := Logic.GetBaseConfig()
	WriterResponse(w, 1, "", bc)
}
