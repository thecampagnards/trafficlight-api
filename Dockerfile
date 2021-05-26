FROM scratch

ENV GIN_MODE=release

WORKDIR /app

COPY app .
COPY drivers .

CMD ["./app"]

EXPOSE 8080