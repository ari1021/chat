module github.com/ari1021/websocket

go 1.15

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0 // Need to use github.com/oxequa/realize, used in ./Dockerfile

require (
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	// github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.17
)
