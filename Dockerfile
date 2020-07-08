FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN cd cmd/server && go build -o main

# Remove SSH keys
RUN rm -rf /root/.ssh/

EXPOSE 8080

CMD ["./cmd/server/main"]