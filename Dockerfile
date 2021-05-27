FROM ubuntu

ENV GIN_MODE=release

WORKDIR /app

COPY app .
COPY drivers drivers

CMD ["./app"]

EXPOSE 8080