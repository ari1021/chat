// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" { // "/"以外であればNot found
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" { // GET以外であればMethod not allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html") // home.htmlを呼ぶ
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)                                       // "/"を叩くとserveHomeが呼ばれる
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { // "/ws"を叩くとserveWsが呼ばれる
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil) //　errorを受け取って処理する
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
