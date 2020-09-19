FROM golang:1.12
COPY . /src
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o DrawIt .

FROM heroku/heroku:18
WORKDIR /app
COPY --from=0 /src/DrawIt /app
CMD ["./DrawIt"]