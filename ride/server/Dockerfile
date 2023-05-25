FROM golang:1.20.1

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /ride-server

CMD [ "/ride-server" ]
