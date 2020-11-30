FROM golang:1.15.4-alpine

WORKDIR /go/src

COPY . .

RUN go mod download

RUN apk add --no-cache git \
    && go get github.com/oxequa/realize

CMD [ "realize", "start"]