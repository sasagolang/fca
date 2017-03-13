package http

import (
	"fmt"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func MyWalletFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	myWallet, err := Logic.MyWallet(uid)
	if err != nil {
		fmt.Printf("MyWalletFunc error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", myWallet)
}
