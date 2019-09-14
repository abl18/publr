FROM golang:1.12 AS base
WORKDIR /go/src/github.com/prksu/publr
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
COPY vendor vendor
RUN go mod verify

FROM base AS build
ENV GO111MODULE=on
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /sites cmd/service/sites/main.go

FROM alpine:latest
COPY --from=build /sites /sites
EXPOSE 9000
ENTRYPOINT [ "/sites" ]