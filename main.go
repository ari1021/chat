// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
)

var addr = flag.String("addr", ":"+os.Getenv("PORT"), "http service address")

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
	http.ServeFile(w, r, "view/index.html") // index.htmlを呼ぶ
}

func main() {
	flag.Parse()
	hub := websocket.NewHub()
	go hub.Run()
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Static("/", "./view/index.html")
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, world!")
	// })
	e.GET("/ws", func(c echo.Context) error {
		websocket.ServeWs(hub, c)
		return nil
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	// e.Logger.Fatal(e.Start(":1323"))
	// http.HandleFunc("/", serveHome)                                       // "/"を叩くとserveHomeが呼ばれる
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { // "/ws"を叩くとserveWsが呼ばれる
	// 	websocket.ServeWs(hub, w, r)
	// })
	// err := http.ListenAndServe(*addr, nil) //　errorを受け取って処理する
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
