FROM golang:1.16-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get -t .
# RUN go build -o proxy/proxy

EXPOSE 8083

CMD [ "go run app/proxy/proxy.go" ]