FROM golang:alpine AS builder

ENV GIN_MODE=release
ENV PORT=2021

WORKDIR /app/

COPY . /app/

RUN go get -d -v /app/

RUN go build /app/

EXPOSE $PORT

ENTRYPOINT ["/app/pikpik"]
