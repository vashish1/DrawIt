From golang:latest

RUN mkdir /server

ADD . /server

WORKDIR /server

ENV GO111MODULE=on

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o DrawIt

# EXPOSE 3000

CMD ["./DrawIt"]