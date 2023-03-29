FROM golang:alpine3.16

RUN mkdir -p /go-users-api
WORKDIR /go-users-api
COPY . .

RUN go build
RUN chmod +x go-user-api

CMD ["./go-user-api"]
