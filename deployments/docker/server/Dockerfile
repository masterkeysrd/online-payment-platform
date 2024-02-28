FROM golang:1.21.7-alpine as BUILDER
WORKDIR /app
RUN apk add --no-cache bash
COPY ./go.mod ./go.sum ./
COPY ./cmd/server/main.go ./cmd/server/main.go
COPY ./api ./api
COPY ./internal ./internal
COPY ./scripts ./scripts
RUN go mod download && go mod verify && go mod tidy
RUN  bash ./scripts/server_build.sh

FROM alpine:3.14.2
WORKDIR /app
COPY --from=BUILDER /app/bin/server /app/server
CMD ["/app/server"]
