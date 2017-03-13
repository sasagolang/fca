package http

import (
	"encoding/base64"
	"fca/model"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetFeedbackFunc(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	fbs, err := Logic.GetFeedback(uid)
	if err != nil {
		fmt.Printf("GetFeedbackFunc error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", fbs)
}
func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("getCurrentPath:%v\n", dir)
	if err != nil {
		panic(err)
		//beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func CreateFeedbackFunc(w rest.ResponseWriter, r *rest.Request) {
	fb := model.CreateFeedbackRequest{}
	err := r.DecodeJsonPayload(&fb)
	if err != nil {
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	if fb.Photo != "" {
		//fmt.Printf("CreateFeedbackFunc.fb.Photo :%v\n", fb.Photo)
		//	ddd, _ := base64.StdEncoding.DecodeString(fb.Photo)
		filename := "/img/feedback/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "." + fb.PhotoSuffix //成图片文件并把文件写入到buffer

		ddd, err3 := base64.StdEncoding.DecodeString(fb.Photo) //成图片文件并把文件写入到buffer
		if err3 != nil {
			panic(err3)
		}
		//fmt.Printf("CreateFeedbackFunc.ddd:%v\n", ddd)
		err2 := ioutil.WriteFile(getCurrentPath()+filename, ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		//err2 := ioutil.WriteFile(filename, ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		if err2 != nil {
			//panic(err2)
			fmt.Printf("CreateFeedbackFunc:%v\n", err2)
		} else {
			fb.Photo = filename
		}
	}
	feedback, err := Logic.CreateFeedback(fb.UID, fb.FeedType, fb.Description, fb.Photo, fb.QuestionType, fb.ElectricPileName)
	if err != nil {
		fmt.Printf("GetFeedbackFunc error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", feedback)
}
