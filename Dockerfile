FROM golang:alpine

WORKDIR /app

COPY ./go.mod ./

RUN go get -u -v /app/cmd/fatt/

CMD [ "tail", "-F", "/dev/null" ]
