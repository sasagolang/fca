package http

import (
	"bytes"
	"encoding/base64"
	"fca/model"
	"fmt"
	"net/http"
	"strconv"

	"fca/libs"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetVerifyImg(w http.ResponseWriter, r *http.Request) {
	d := make([]byte, 4)
	s := NewLen(4)
	ss := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		ss += strconv.FormatInt(int64(d[v]), 32)
	}
	w.Header().Set("Content-Type", "image/png")
	NewImage(d, 100, 40).WriteTo(w)
	fmt.Println(ss)

}

const aesTable = "abcdefghijklmnopkrstuvwsyz012345"

func GetVerifyImgData(w rest.ResponseWriter, r *rest.Request) {
	widthStr := r.PathParam("width")
	heightStr := r.PathParam("height")
	width := 100
	height := 40
	if widthStr != "" {
		width, _ = strconv.Atoi(widthStr)
	}
	if heightStr != "" {
		height, _ = strconv.Atoi(heightStr)
	}

	d := make([]byte, 4)
	s := NewLen(4)
	ss := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		ss += strconv.FormatInt(int64(d[v]), 32)
	}

	b1 := new(bytes.Buffer)
	NewImage(d, width, height).WriteTo(b1)
	vy := model.VerrifyImg{}
	vy.Data = base64.StdEncoding.EncodeToString(b1.Bytes())
	dats, _ := libs.AesEncrypt([]byte(ss), []byte(aesTable))

	fmt.Println(ss)
	vy.Key = base64.StdEncoding.EncodeToString(dats)

	//if err != nil {
	//	fmt.Printf("GetQuestions error(%v)\n", err)
	//	WriterResponse(w, 2, err)
	//	return
	//}
	WriterResponse(w, 1, "", vy)
}
func VerifyCode(w rest.ResponseWriter, r *rest.Request) {
	keyStr := r.PathParam("keystr")
	valueStr := r.PathParam("valuestr")
	keyStrByte, _ := base64.StdEncoding.DecodeString(keyStr)
	keyStrdata, _ := libs.AesDecrypt(keyStrByte, []byte(aesTable))
	b := false
	if valueStr == string(keyStrdata) {
		b = true
	}
	fmt.Println(valueStr)
	fmt.Println(string(keyStrdata))
	WriterResponse(w, 1, "", b)
}
