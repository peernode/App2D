package main

import (
	"log"
	"net/http"
	"time"

	l4g "github.com/alecthomas/log4go"
	"github.com/julienschmidt/httprouter"
)

var gAppDir = "BestvVR"
var logFilename = "srvLog_"
var gLogger l4g.Logger

var giPhoneURL = "https://itunes.apple.com/cn/app/vr-ying-yuan-lao-si-ji-kanvr/id1120003008?mt=8"
var gAndroidURL = "http://vr.ott.bestv.com.cn:8808/vr/static/BestvVR_guanfang.apk"
var gOtherURL = "http://vr.ott.bestv.com.cn:8808/vr/static/BestvVR_guanfang.apk"

//init for logger
func initLogger() {
	gLogger = make(l4g.Logger)

	gLogger.AddFilter("stdout", l4g.INFO, l4g.NewConsoleLogWriter())
	flw := l4g.NewFileLogWriter(logFilename+time.Now().Format("20060102"), true)
	flw.SetFormat("[%D %T] [%L] %M")
	flw.SetRotateDaily(true)
	gLogger.AddFilter("logfile", l4g.FINEST, flw)

	gLogger.Info("Init logger! The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
}

func initHTTPRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/vr/download", DownloadApp)

	router.GET("/vr/static/download.html", download)
	router.GET("/vr/static/image/top.png", top)
	router.GET("/vr/static/image/button.png", button)
	router.GET("/vr/static/image/bottom.png", bottom)
	router.GET("/vr/static/BestvVR_guanfang.apk", guanfang)
	router.GET("/vr/downloadRedirect", downloadRedirect)

	//router.ServeFiles("/vr/static/*filepath", http.Dir(gAppDir)) //下载相应的媒体文件

	return router
}

func main() {
	initLogger()
	router := initHTTPRouter()

	s := &http.Server{
		Addr:    ":8808",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}
