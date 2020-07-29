FROM golang:latest as builder

WORKDIR /go/src/app

COPY . .

RUN cd cmd && go build -o main

FROM ubuntu:18.04

COPY --from=builder /go/src/app/cmd/main /app/main

EXPOSE 8080

CMD ["/app/main"]