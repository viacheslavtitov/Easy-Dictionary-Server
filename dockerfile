FROM alpine:latest

COPY easydictionary /app/easydictionary

WORKDIR /app

EXPOSE 8080

CMD ["./easydictionary"]
