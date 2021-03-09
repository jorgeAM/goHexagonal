FROM golang:alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/JorgeAM/goHexagonal

COPY . .

RUN go mod tidy

RUN go build -o /go/bin/goHexagonal cmd/api/main.go

FROM scratch

COPY --from=build /go/bin/goHexagonal /go/bin/goHexagonal

ENTRYPOINT ["/go/bin/goHexagonal"]