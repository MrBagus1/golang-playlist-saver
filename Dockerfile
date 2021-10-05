#stage 1
FROM golang:1.17-alpine AS golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache

#stage 2
RUN go build -o main
FROM alpine:3.14
WORKDIR /root/
COPY --from=golang /app/.env .
COPY --from=golang /app/main .
EXPOSE 8080
CMD ["./main"]
