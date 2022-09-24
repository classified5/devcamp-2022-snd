FROM --platform=linux/arm64 golang:1.17 AS build

ENV GO111MODULE=on

WORKDIR /go/src/devcamp-2022-snd/

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o service .

FROM alpine
COPY --from=build /go/src/devcamp-2022-snd/ go/bin/devcamp-2022-snd/

EXPOSE 9000
WORKDIR go/bin/devcamp-2022-snd/

CMD ["./service"]