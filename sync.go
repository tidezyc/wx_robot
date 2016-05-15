package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type SyncKey struct {
	Key int `json:"Key"`
	Val int `json:"Val"`
}

type SyncKeys []SyncKey

func (keys SyncKeys) String() string {
	keyStr := ""
	for _, key := range keys {
		if keyStr != "" {
			keyStr = keyStr + "|"
		}
		keyStr = keyStr + fmt.Sprintf("%d_%d", key.Key, key.Val)
	}
	return keyStr
}

func DoSync(keys []SyncKey) {
	syncKeys := SyncKeys(keys)
	for {
		log.Println("=========== check sync =============")
		if syncKeys == nil {
			log.Println("syncKeys is nil")
			return
		}
		r := syncCheck(syncKeys)
		if r == nil || r.Retcode != "0" {
			log.Printf("sync check return nil or code not 0 : %v\n", r)
			return
		}
		if r.Selector == "2" {
			log.Println("new msg, do sync")
			syncKeys = sync(syncKeys)
		}
	}
}

func sync(keys SyncKeys) SyncKeys {
	l := fmt.Sprintf("https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxsync?sid=%s&skey=%s&pass_ticket=%s", LoginCookie.Wxsid, LoginCookie.Skey, LoginCookie.Pass_ticket)
	req := LoginCookie.BaseQuquest()
	req["SyncKey"] = map[string]interface{}{
		"Count": len(keys),
		"List":  keys,
	}
	data := httpPostBytes(l, req)
	if data == nil {
		return nil
	}
	var v struct {
		BaseResponse BaseResponse `json:"BaseResponse"`
		SyncKey      struct {
			List SyncKeys `json:"List"`
		} `json:"SyncKey"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil
	}
	return v.SyncKey.List
}

func syncCheck(keys SyncKeys) *SyncCheckResult {
	l := fmt.Sprintf("https://webpush.weixin.qq.com/cgi-bin/mmwebwx-bin/synccheck?r=%d&skey=%s&sid=%s&uin=%d&deviceid=%s&synckey=%s", time.Now().UnixNano()/int64(time.Millisecond), LoginCookie.Skey, LoginCookie.Wxsid, LoginCookie.Wxuin, LoginCookie.DeviceID, keys)
	res := httpGetString(l)
	if res == "" {
		return nil
	}
	parts := strings.Split(res, "=")
	if parts != nil && len(parts) == 2 {
		parts = strings.Split(parts[1][1:len(parts[1])-1], ",")
		if parts != nil && len(parts) == 2 {
			v := &SyncCheckResult{}
			v.Retcode = parts[0][9 : len(parts[0])-1]
			v.Selector = parts[1][10 : len(parts[1])-1]
			return v
		}
	}
	return nil
}
