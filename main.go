package main

import (
	"log"
	"net/http"
	"time"

	l4g "github.com/alecthomas/log4go"
	"github.com/julienschmidt/httprouter"
)

var gAppDir = "BestvVR"
var logFilename = "srvLog.txt"
var gLogger l4g.Logger

//init for logger
func initLogger() {
	gLogger = make(l4g.Logger)

	gLogger.AddFilter("stdout", l4g.INFO, l4g.NewConsoleLogWriter())
	flw := l4g.NewFileLogWriter(logFilename, true)
	flw.SetFormat("[%D %T] [%L] %M")
	flw.SetRotateDaily(true)
	gLogger.AddFilter("logfile", l4g.FINEST, flw)

	gLogger.Info("Init logger! The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
}

func initHTTPRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/vr/downloadAPP", DownloadApp)

	router.ServeFiles("/vr/static/*filepath", http.Dir(gAppDir)) //下载相应的媒体文件

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
