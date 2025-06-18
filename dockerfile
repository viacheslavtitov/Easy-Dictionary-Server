FROM alpine:latest

COPY easydictionary /app/easydictionary
COPY migrations /app/migrations

WORKDIR /app

EXPOSE 8080

CMD ["./easydictionary"]
