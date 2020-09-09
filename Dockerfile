# STEP 1 build executable binary
FROM golang:alpine AS builder

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/api

# STEP 2 build a small image
FROM scratch

ENV VERSION "1"

COPY --from=builder /go/bin/api /go/bin/api

ENTRYPOINT ["/go/bin/api"]
