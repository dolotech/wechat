package main

import (
	"github.com/golang/glog"
	"github.com/songtianyi/wechat-go/plugins/wxweb/cleaner"
	"github.com/songtianyi/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/songtianyi/wechat-go/plugins/wxweb/forwarder"
	"github.com/songtianyi/wechat-go/plugins/wxweb/gifer"
	"github.com/songtianyi/wechat-go/plugins/wxweb/joker"
	"github.com/songtianyi/wechat-go/plugins/wxweb/laosj"
	"github.com/songtianyi/wechat-go/plugins/wxweb/replier"
	"github.com/songtianyi/wechat-go/plugins/wxweb/revoker"
	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/plugins/wxweb/system"
	"github.com/songtianyi/wechat-go/plugins/wxweb/youdao"
	"github.com/songtianyi/wechat-go/wxweb"
	"time"
)

func main() {
	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		glog.Error(err)
		return
	}
	// load plugins for this session
	faceplusplus.Register(session)
	replier.Register(session)
	switcher.Register(session)
	gifer.Register(session)
	cleaner.Register(session)
	laosj.Register(session)
	joker.Register(session)
	revoker.Register(session)
	forwarder.Register(session)
	system.Register(session)
	youdao.Register(session)

	// enable by type example
	if err := session.HandlerRegister.EnableByType(wxweb.MSG_SYS); err != nil {
		glog.Error(err)
		return
	}

	for {
		if err := session.LoginAndServe(false); err != nil {
			glog.Error("session exit, %s", err)
			for i := 0; i < 3; i++ {
				glog.Info("trying re-login with cache")
				if err := session.LoginAndServe(true); err != nil {
					glog.Error("re-login error or session down, %s", err)
				}
				time.Sleep(3 * time.Second)
			}
			if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.TERMINAL_MODE); err != nil {
				glog.Error("create new sesion failed, %s", err)
				break
			}
		} else {
			glog.Info("closed by user")
			break
		}
	}
}
