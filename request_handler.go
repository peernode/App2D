package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strings"
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
