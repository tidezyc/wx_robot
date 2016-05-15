package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func SendMsg(to string, msg string) bool {
	if LoginCookie == nil {
		return false
	}
	l := fmt.Sprintf("https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxsendmsg?lang=en_US&pass_ticket=%s", LoginCookie.Pass_ticket)
	msgId := fmt.Sprintf("%d", time.Now().UnixNano()/100)
	req := LoginCookie.BaseQuquest()
	req["Msg"] = map[string]interface{}{
		"ClientMsgId":  msgId,
		"Content":      msg,
		"FromUserName": LoginCookie.CurrentUsername,
		"LocalID":      msgId,
		"ToUserName":   to,
		"Type":         1,
	}
	data := httpPostBytes(l, req)
	if data == nil {
		return false
	}
	var v struct {
		BaseResponse BaseResponse `json:"BaseResponse"`
		LocalID      string       `json:"LocalID"`
		MsgID        string       `json:"MsgID"`
	}
	json.Unmarshal(data, &v)
	return v.BaseResponse.Ret == 0
}
