FROM --platform=linux/arm64 golang:1.17 AS build

ENV GO111MODULE=on

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

WORKDIR /go/src/github.com/devcamp-2022-snd

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

FROM alpine
COPY --from=build /go/src/github.com/devcamp-2022-snd/service go/bin/devcamp-2022-snd/

EXPOSE 9000
WORKDIR go/bin/devcamp-2022-snd/

CMD ["./service"]