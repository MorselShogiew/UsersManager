FROM golang:1.18
WORKDIR /go/src/github.com/MorselShogiew/UsersManager

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./proto/. ./proto
COPY ./packages/. ./packages
WORKDIR /go/src/github.com/MorselShogiew/UsersManagerservices/server
COPY ./services/server/. .
RUN go build -o app .

ENTRYPOINT [ "./app" ]

# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=0 /go/src/github.com/MorselShogiew/UsersManager/services/server/app ./
# CMD ["./app"]