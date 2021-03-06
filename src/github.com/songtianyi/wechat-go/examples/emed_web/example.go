package main

import (
	"github.com/golang/glog"
	"github.com/songtianyi/wechat-go/wxweb"
	"net/http"
	"os"
	"path/filepath"
	"flag"
)

func main() {
	flag.Parse()
	// get web server root path
	cur_dir := filepath.Dir(os.Args[0])
	public_dir := filepath.Join(cur_dir, "public")

	// create session and put qrcode image to webroot
	session, err := wxweb.CreateWebSessionWithPath(nil, nil, public_dir)
	if err != nil {
		glog.Error(err)
		return
	}

	// serve and wait for wechat msg
	go session.LoginAndServe(true)

	// serve http
	http.ListenAndServe(":8080", http.FileServer(http.Dir(public_dir)))

	// then visit http://target:8080/public/  + session.QrcodePath
}
