package main

import (
	"net/http"

	"strconv"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
)

type Result struct {
	ResultCode    int
	ResultMessage string
	Data          interface{}
}

func WriterResponse(c *gin.Context, code int, errMsg string, data interface{}) {
	result := Result{}
	result.ResultCode = code
	result.ResultMessage = errMsg
	result.Data = data
	c.JSON(200, result)
}
func InitHttp() {
	router := gin.New()
	router.GET("/", welcomeFunc)
	router.GET("/start/:uuid/:mins", StartFunc)
	router.GET("/stop/:uuid", StopFunc)
	router.GET("/startlight/:uuid", StartLightFunc)
	router.GET("/stoplight/:uuid", StopLightFunc)
	router.Run(":8997")
}
func welcomeFunc(c *gin.Context) {
	//firstname := c.DefaultQuery("firstname", "Guest")
	//lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	//	p, _ := ioutil.ReadAll(c.Request.Body)

	c.String(http.StatusOK, "welcome！")
}
func StartFunc(c *gin.Context) {
	uuid := c.Param("uuid")
	mins, _ := strconv.Atoi(c.Param("mins"))
	if uuid == "12344445" || uuid == "12344446" {
		if uuid == "12344445" {
			go Mock9(uuid, t1)
		}
		if uuid == "12344446" {
			go Mock9(uuid, t2)
		}
		WriterResponse(c, 1, "", nil)
		return
	}
	ch := GetChannelByUUID(uuid)
	if ch != nil {

		Logic.Start_cmd(int32(mins), time.Now().Add(time.Duration(int64(mins))*time.Minute), ch.signal)
	} else {
		WriterResponse(c, 800, "设备未找到", nil)
		return
	}

	WriterResponse(c, 1, "", nil)

}
func StopFunc(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "12344445" {
		t1 <- 1
	} else if uuid == "12344446" {
		t2 <- 1
	} else {
		ch := GetChannelByUUID(uuid)
		if ch != nil {

			Logic.Stop_cmd(ch.signal)
		}
	}
	WriterResponse(c, 1, "", nil)
}
func StartLightFunc(c *gin.Context) {
	fmt.Printf("StartLightFunc:%v\n", c)
	uuid := c.Param("uuid")

	ch := GetChannelByUUID(uuid)

	if ch != nil {

		Logic.StartLight_cmd(ch.signal)
	}
	WriterResponse(c, 1, "", nil)
}
func StopLightFunc(c *gin.Context) {
	uuid := c.Param("uuid")

	ch := GetChannelByUUID(uuid)
	if ch != nil {

		Logic.StopLight_cmd(ch.signal)
	}
	WriterResponse(c, 1, "", nil)
}
