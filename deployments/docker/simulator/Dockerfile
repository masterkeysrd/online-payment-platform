FROM golang:1.21.7-alpine as BUILDER
WORKDIR /app
RUN apk add --no-cache bash
COPY ./go.mod ./go.sum ./
COPY ./cmd/simulator/main.go ./cmd/simulator/main.go
COPY ./scripts ./scripts
RUN bash ./scripts/simulator_build.sh

FROM alpine:3.14.2
WORKDIR /app
COPY --from=BUILDER /app/bin/simulator /app/simulator
CMD ["/app/simulator"]