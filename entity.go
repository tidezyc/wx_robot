package main

import "encoding/xml"

type WxCookie struct {
	XMLName         xml.Name `xml:"error"`
	Ret             int      `xml:"ret"`
	Message         string   `xml:"message"`
	Skey            string   `xml:"skey"`
	Wxsid           string   `xml:"wxsid"`
	Wxuin           int      `xml:"wxuin"`
	Pass_ticket     string   `xml:"pass_ticket"`
	Isgrayscale     bool     `xml:"isgrayscale"`
	DeviceID        string
	CurrentUsername string
}

func (cookie *WxCookie) BaseQuquest() map[string]interface{} {
	return map[string]interface{}{
		"BaseRequest": map[string]interface{}{
			"DeviceID": cookie.DeviceID,
			"Sid":      cookie.Wxsid,
			"Uin":      cookie.Wxuin,
			"Skey":     cookie.Skey,
		},
	}
}

type Contact struct {
	UserName   string `json:"UserName"`
	NickName   string `json:"NickName"`
	HeadImgUrl string `json:"HeadImgUrl"`
	Statues    int    `json:"Statues"`    // 0:单聊和公众号; 1:群聊
	AttrStatus int    `json:"AttrStatus"` // 0:公众号和群聊; >0:单聊
}

type BaseResponse struct {
	Ret    int    `json:"Ret"`
	ErrMsg string `json:"ErrMsg"`
}

type SyncCheckResult struct {
	Retcode  string `json:"retcode"`
	Selector string `json:"selector"`
}
