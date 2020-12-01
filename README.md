# websocket

## 使い方

サーバーの起動方法

1. docker-compose build
1. docker-compose up

## ディレクトリ構成

```:text
.
├── Dockerfile
├── README.md
├── config
│   └── config.go
├── controller  // modelを使ってデータの操作をするファイル
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
├── model  // structだけが書かれたファイル
├── server
│   ├── api
│   │   └── router.go
│   └── websocket
│       ├── client.go
│       └── hub.go
└── view  // 画面表示用のファイル
    └── index.html
```
