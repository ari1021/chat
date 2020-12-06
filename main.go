// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"os"

	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	flag.Parse()
	hub := websocket.NewHub()
	go hub.Run()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./view/index.html")
	e.GET("/ws", func(c echo.Context) error {
		websocket.ServeWs(hub, c)
		return nil
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
