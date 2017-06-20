package libs

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
)

var Log *log.Logger

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//	fmt.Printf("getCurrentPath:%v\n", dir)
	if err != nil {
		panic(err)
		//beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func NewLogger() *log.Logger {
	if Log != nil {
		return Log
	}

	Log = log.New()
	log.SetOutput(os.Stdout)
	//path := "/var/log/go.log"
	infolog, _ := rotatelogs.New(
		getCurrentPath()+"/log/info.log",
		rotatelogs.WithLinkName(getCurrentPath()+"/log/info.log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	debuglog, _ := rotatelogs.New(
		getCurrentPath()+"/log/debug.log",
		rotatelogs.WithLinkName(getCurrentPath()+"/log/debug.log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	errorlog, _ := rotatelogs.New(
		getCurrentPath()+"/log/error.log",
		rotatelogs.WithLinkName(getCurrentPath()+"/log/error.log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	//writer.LinkName = path
	//writer.RotationTime = time.Duration(86400) * time.Second // rotate once a day
	//	writer.MaxAge = time.Duration(604800) * time.Second      // keep one week of log files
	//writer.Offset = 0

	Log.Hooks.Add(lfshook.NewHook(lfshook.WriterMap{
		log.InfoLevel:  infolog,
		log.DebugLevel: debuglog,
		log.ErrorLevel: errorlog,
	}))
	return Log

}
