FROM golang:alpine

WORKDIR /fatt

COPY . .

RUN go build -o /fatt/fatt /fatt/cmd/fatt/*

RUN cp /fatt/fatt /usr/local/bin

ENTRYPOINT [ "fatt" ]

# Usage:
# > docker run --m --name fatt [ARGS]
