################################
# STEP 1 build executable binary
################################

FROM golang:1.20.3-alpine AS builder

COPY . /app/src/
WORKDIR /app/src/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/grpc_server/main.go

############################
# STEP 2 build a small image
############################

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/src/bin/auth_server .

CMD [ "./auth_server" ]
