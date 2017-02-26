package http

import (
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetQuestions(w rest.ResponseWriter, r *rest.Request) {
	qs, err := Logic.GetQuestions()
	if err != nil {
		fmt.Printf("GetQuestions error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", qs)
}
