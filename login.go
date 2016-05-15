package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

var LoginCookie *WxCookie

func GetUUID(appid string) string {
	jslogin := fmt.Sprintf("https://login.weixin.qq.com/jslogin?appid=%s&fun=%s", appid, "new")
	res := httpGetString(jslogin)
	if res != "" {
		parts := strings.Split(res, ";")
		if len(parts) == 3 {
			parts = strings.Split(parts[1], " = ")
			if len(parts) == 2 && strings.Contains(parts[0], "uuid") {
				return parts[1][1 : len(parts[1])-1]
			}
		}
	}
	return ""
}

func CheckLogin(uuid string) {
	for {
		res := httpGetString(fmt.Sprintf("https://login.weixin.qq.com/cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=%s&tip=0", uuid))
		log.Println("------- do check login: " + res)
		if res != "" {
			parts := strings.Split(res, ";")
			if len(parts) >= 2 {
				if parts[0] == "window.code=200" {
					sep := "window.redirect_uri="
					redirectUrl := parts[1][strings.Index(parts[1], sep)+len(sep)+1 : len(parts[1])-1]
					go doLogin(redirectUrl)
					return
				}
			}
		}
	}
}

func doLogin(url string) {
	cookie := getCookie(url)
	log.Printf("get cookie:%v\n", *cookie)
	if cookie != nil {
		wxInit(cookie)
		if LoginCookie == nil {
			log.Println("init faild, get cookie is nil")
			return
		}
		LoadContact(cookie)
	}
}

func wxInit(cookie *WxCookie) {
	l := fmt.Sprintf("https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxinit?pass_ticket=%s", cookie.Pass_ticket)
	data := httpPostBytes(l, cookie.BaseQuquest())

	var v struct {
		Res  BaseResponse `json:"BaseResponse"`
		User struct {
			UserName string `json:"UserName"`
		} `json:"User"`
		SyncKey struct {
			List []SyncKey `json:"List"`
		} `json:"SyncKey"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Println(err)
		return
	}
	if v.Res.Ret != 0 {
		log.Println("wx init return err" + v.Res.ErrMsg)
		return
	}
	cookie.CurrentUsername = v.User.UserName
	go DoSync(v.SyncKey.List)

	LoginCookie = cookie
}
func getCookie(url string) *WxCookie {
	res := httpGetBytes(url)
	cookie := WxCookie{
		DeviceID: "e723568210709414",
	}
	xml.Unmarshal(res, &cookie)
	if cookie.Ret != 0 {
		return nil
	}
	return &cookie
}
