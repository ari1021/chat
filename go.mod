module github.com/ari1021/websocket

go 1.15

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0 // Need to use github.com/oxequa/realize, used in ./Dockerfile

require (
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/mock v1.4.4
	github.com/google/go-cmp v0.5.4
	github.com/gorilla/websocket v1.4.2
	// github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.17
	github.com/leodido/go-urn v1.2.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
)
