FROM golang:1.19 AS builder

WORKDIR /go/src/BWG/

ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/http/main.go

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/BWG/main .
CMD ["./main"]