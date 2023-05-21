##
## Build
##

FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /main

FROM alpine

COPY config.ini . 
COPY --from=builder /main /main

CMD ["/main"]