FROM golang:1.19-alpine as build
LABEL maintainer "HDO"

RUN apk --no-cache add ca-certificates \
     && apk --no-cache add --virtual build-deps git gcc musl-dev

COPY ./ /go/src/gitlab.hdo.no/fg-telefoni/core-exporter
WORKDIR /go/src/gitlab.hdo.no/fg-telefoni/core-exporter

RUN go get \
 && go test ./... \
 && go build -o /bin/main

FROM alpine:3

RUN apk --no-cache add ca-certificates \
     && addgroup exporter \
     && adduser -S -G exporter exporter
USER exporter
COPY --from=build /bin/main /bin/main
ENV LISTEN_PORT=9172
EXPOSE 9172
ENTRYPOINT [ "/bin/main" ]
