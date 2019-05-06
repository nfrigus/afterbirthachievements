FROM golang:1.11.5-alpine as builder
WORKDIR /build
COPY . . 
RUN apk --update add git
RUN go get github.com/gorilla/mux && go build -o main .

FROM alpine
EXPOSE 9090
WORKDIR /root
RUN apk --update add ca-certificates gettext
COPY tmpl tmpl
COPY data data
COPY config.json.example .
COPY --from=builder /build/main .
CMD envsubst < config.json.example > config.json && ./main
