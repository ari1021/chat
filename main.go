// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"os"

	"github.com/ari1021/websocket/server/api"
	"github.com/ari1021/websocket/server/websocket"
)

func main() {
	flag.Parse()
	hub := websocket.NewHub()
	go hub.Run()
	e := api.NewEcho(hub)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
