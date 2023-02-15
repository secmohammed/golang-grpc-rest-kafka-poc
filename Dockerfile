#Build binary
FROM golang:alpine as builder
LABEL maintainer = "Mohammed Osama"

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main


# Build container to run binary
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

EXPOSE 8001
EXPOSE 8002
ENTRYPOINT [ "./main" ]
