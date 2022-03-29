FROM golang:alpine

RUN mkdir -p /location-app
WORKDIR /location-app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./app ./main.go

# COPY /api/app.out .
expose 3001
ENTRYPOINT ["./app"]

