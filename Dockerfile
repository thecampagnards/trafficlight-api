FROM alpine

ENV GIN_MODE=release

RUN apk add --no-cache hidapi linux-headers

WORKDIR /app

COPY app .
COPY drivers drivers

CMD ["./app"]

EXPOSE 8080