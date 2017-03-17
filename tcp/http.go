package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func InitHttp(){
	router := gin.New()
	router.GET("/", welcomeFunc)
}
func welcomeFunc(c *gin.Context) {
	//firstname := c.DefaultQuery("firstname", "Guest")
	//lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	//	p, _ := ioutil.ReadAll(c.Request.Body)

	c.String(http.StatusOK, "welcome！")
}
