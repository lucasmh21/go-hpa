FROM golang:alpine AS builder
COPY ./src/main /go/main
WORKDIR /go/main
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /go /go
ENTRYPOINT ["/go/main/main"]