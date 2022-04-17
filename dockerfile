FROM golang:alpine
WORKDIR /go/src/news
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o news main.go

FROM alpine:latest
EXPOSE 7000

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
WORKDIR /app/
COPY --from=0 /go/src/news  .

CMD ["./news"]
