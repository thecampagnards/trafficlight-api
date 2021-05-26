FROM scratch

ENV GIN_MODE=release

WORKDIR /app

COPY app .

CMD app

EXPOSE 8080