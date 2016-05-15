// Package main provides ...
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

const (
	APPID string = "wx782c26e4c19acffb"
)

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/contacts", getContacts)
	http.HandleFunc("/msg", sendMsg)
	http.HandleFunc("/favicon.ico", handlerICon)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	clearContact()
	uuid := GetUUID(APPID)
	if uuid == "" {
		w.Write([]byte("error"))
	} else {
		t := template.New("main html")
		t, err := t.Parse(MAIN_HTML)
		if err != nil {
			log.Fatalln(err)
		}

		t.Execute(w, uuid)
		go CheckLogin(uuid)
	}
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	f := GetFriends()
	g := GetGroups()
	p := GetPublics()
	res := &struct {
		Ret     int       `json:"ret"`
		Friends []Contact `json:"friends"`
		Groups  []Contact `json:"groups"`
		Publics []Contact `json:"publics"`
	}{}
	if f == nil || g == nil || p == nil || (len(f) == 0 && len(g) == 0 && len(p) == 0) {
		res.Ret = 1
	} else {
		res.Ret = 0
		res.Friends = f
		res.Groups = g
		res.Publics = p
	}
	data, _ := json.Marshal(res)
	w.Write(data)
}

func sendMsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	to := r.Form.Get("to")
	msg := r.Form.Get("msg")
	if to == "" || msg == "" {
		w.Write([]byte("err params"))
		return
	}
	b := SendMsg(to, msg)
	if b {
		w.Write([]byte("success"))
	} else {
		w.Write([]byte("failed"))
	}
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}
