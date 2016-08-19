package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//按平台返回不同的URL
func DownloadApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	agentS := r.Header.Get("User-Agent")
	fmt.Println(agentS)
	if strings.Index(agentS, "Android") > -1 || strings.Index(agentS, "Adr") > -1 {
		fmt.Println("android")
		http.Redirect(w, r, "http://vr.ott.bestv.com.cn:8808/vr/static/BestvVR_guanfang.apk", http.StatusFound)
	} else {
		fmt.Println("iphone")
		http.Redirect(w, r, "https://itunes.apple.com/cn/app/vr-ying-yuan-lao-si-ji-kanvr/id1120003008?mt=8", http.StatusMovedPermanently)
	}
}

// 响应静态页面
func download(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	agentS := r.Header.Get("User-Agent")
	fmt.Println(agentS)
	// if strings.Index(agentS, "Android") > -1 || strings.Index(agentS, "Adr") > -1 {
	// 	fmt.Println("android")
	// } else {
	// 	fmt.Println("iphone")
	// }

	req := fmt.Sprintf("dowload|%s|%s|%s", getIP(r), agentS, r.RequestURI)
	gLogger.Info(req)

	fp := path.Join("BestvVR", "download.html")
	fmt.Println("fp: " + fp)
	http.ServeFile(w, r, fp)
}

func top(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fp := path.Join("BestvVR/image", "top.png")
	fmt.Println("fp: " + fp)
	http.ServeFile(w, r, fp)
}

func button(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fp := path.Join("BestvVR/image", "button.png")
	fmt.Println("fp: " + fp)
	http.ServeFile(w, r, fp)
}

func bottom(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fp := path.Join("BestvVR/image", "bottom.png")
	fmt.Println("fp: " + fp)
	http.ServeFile(w, r, fp)
}

func guanfang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fp := path.Join("BestvVR", "BestvVR_guanfang.apk")
	fmt.Println("fp: " + fp)
	http.ServeFile(w, r, fp)
}

// 中转页面，用于统计是否有人点击下载按钮
func downloadRedirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	agentS := r.Header.Get("User-Agent")
	fmt.Println(agentS)

	req := fmt.Sprintf("clickRedirect|%s|%s", getIP(r), agentS)
	gLogger.Info(req)

	if strings.Index(agentS, "iPhone") > -1 || strings.Index(agentS, "iOS") > -1 {
		fmt.Println("iphone")
		http.Redirect(w, r, giPhoneURL, http.StatusFound)
	} else if strings.Index(agentS, "Android") > -1 || strings.Index(agentS, "Adr") > -1 {
		fmt.Println("android")
		http.Redirect(w, r, gAndroidURL, http.StatusMovedPermanently)
	} else {
		fmt.Println("other")
		http.Redirect(w, r, gOtherURL, http.StatusMovedPermanently)
	}
}
