package main

import (
	"container/list"
	"fmt"
	"strconv"
)

var db = list.New()

func store(a *SS) {
	db.PushBack(a)
}

func readAll() {
	var guiConfig GuiConfig
	guiConfig.Index = 0
	guiConfig.Random = false
	guiConfig.Global = false
	guiConfig.Enabled = true
	guiConfig.ShareOverLan = false
	guiConfig.IsDefault = false
	guiConfig.LocalPort = 1080
	guiConfig.PacUrl = nil
	guiConfig.UseOnlinePac = false
	guiConfig.ReconnectTimes = 3
	guiConfig.RandomAlgorithm = 0
	guiConfig.TTL = 0
	guiConfig.ProxyEnable = false
	guiConfig.ProxyType = 0
	guiConfig.ProxyHost = nil
	guiConfig.ProxyPort = 0
	guiConfig.ProxyAuthUser = nil
	guiConfig.ProxyAuthPass = nil
	guiConfig.AuthUser = nil
	guiConfig.AuthPass = nil

	for e := db.Front(); e != nil; e = e.Next() {
		var configEle ConfigEle
		configEle.Enable = false
		configEle.Password = e.Value.(*SS).password
		configEle.Method = e.Value.(*SS).method
		configEle.Remarks = ""
		configEle.Server = e.Value.(*SS).ip


		port, _ := strconv.Atoi(e.Value.(*SS).port)
		configEle.ServerPort = port



	}


	fmt.Println(db.Len())
}
