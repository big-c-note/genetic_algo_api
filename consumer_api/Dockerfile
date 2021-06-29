FROM golang:1.13.15

WORKDIR /app
COPY . /app

RUN cd /app && go build -o consumer_api

CMD ["/app/consumer_api"]
