package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var Friends []Contact
var Groups []Contact
var Publics []Contact

func GetFriends() []Contact {
	return Friends
}

func GetGroups() []Contact {
	return Groups
}

func GetPublics() []Contact {
	return Publics
}

func LoadContact(cookie *WxCookie) {
	clearContact()
	l := fmt.Sprintf("https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxgetcontact?pass_ticket=%s&r=%d&seq=0&skey=%s", cookie.Pass_ticket, time.Now().UnixNano()/int64(time.Millisecond), cookie.Skey)
	data := httpGetBytes(l)
	if data == nil {
		return
	}
	var v struct {
		BaseResponse BaseResponse `json:"BaseResponse"`
		MemberCount  int          `json:"MemberCount"`
		MemberList   []Contact    `json:"MemberList"`
	}
	json.Unmarshal(data, &v)
	if v.BaseResponse.Ret != 0 {
		log.Println("get contact res err: " + v.BaseResponse.ErrMsg)
		return
	}
	if v.MemberCount <= 0 {
		log.Println("get contact size <=0")
		return
	}
	parseContact(v.MemberList)
}

func parseContact(list []Contact) {
	for _, c := range list {
		if c.Statues == 0 {
			if c.AttrStatus == 0 {
				Publics = append(Publics, c)
			} else {
				Friends = append(Friends, c)
			}
		} else {
			Groups = append(Groups, c)
		}
	}
}

func clearContact() {
	Friends = []Contact{}
	Groups = []Contact{}
	Publics = []Contact{}
}
