FROM golang:1.10.3-alpine as gobuild
WORKDIR /
ENV GOPATH="/"
RUN apk update && apk add pkgconfig build-base bash autoconf automake libtool gettext openrc git libzmq zeromq-dev mercurial
#COPY . . if you update the libs below build with --no-cache
RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/toshbrown/lib-go-databox
COPY . .
RUN addgroup -S databox && adduser -S -g databox databox
RUN GGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w' -o app /src/*.go

FROM amd64/alpine:3.8
COPY --from=gobuild /etc/passwd /etc/passwd
RUN apk update && apk add libzmq
USER databox
WORKDIR /
COPY --from=gobuild /app .
COPY --from=gobuild /ui/dist /www
LABEL databox.type="app"
EXPOSE 8080

CMD ["./app"]
