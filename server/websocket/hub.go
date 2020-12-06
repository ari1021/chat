// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool //Clientのpointerがkeyでvalueがbool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub { //新たにHubを作ってそのpointerを返す
	return &Hub{
		broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register: //Hubのregisterというchannelに*Clientが入っているとき
			h.clients[client] = true //clientを登録する
		case client := <-h.unregister: //Hubのunregisterというchannelに*Clientが入っているとき
			if _, ok := h.clients[client]; ok { //そのclientが登録されていれば
				delete(h.clients, client) //削除する
				close(client.Send)        //そのclientのchannelをcloseする
			}
		case message := <-h.broadcast: //Hubのbroadcastというchannelにmessage(byte)が入っているとき
			for client := range h.clients { //登録されているclient全員に対して
				select {
				case client.Send <- message: //messageを送ることができれば送る
				default: //送ることができなければ
					close(client.Send)        //channelをcloseして
					delete(h.clients, client) //Hubからdeleteする
				}
			}
		}
	}
}
