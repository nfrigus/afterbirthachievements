FROM golang:1.11.5-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
EXPOSE 9090
RUN apk --update add git gettext
RUN go get github.com/gorilla/mux && go build -o main .
CMD envsubst < config.json.example > config.json && ./main
