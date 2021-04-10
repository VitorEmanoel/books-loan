FROM golang:1.16.3-alpine as compiler
WORKDIR /app
COPY . /app
RUN go mod download && go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat curl
WORKDIR /app
COPY --from=compiler /app/app .
CMD ["./app"]
HEALTHCHECK --interval=1m --timeout=3s CMD curl --fail http://localhost:$PORT/api/health || exit 1
