package rrutils

import (
	"github.com/golang/glog"
	"net/http"
	_ "net/http/pprof"
)

func StartProfiling() {
	go func() {
		if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
			glog.Error("Start profiling fail, %s", err)
			return
		}
	}()
}
