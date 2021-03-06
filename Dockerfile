FROM node:alpine as build
WORKDIR /app
COPY ./frontend/package.json ./
COPY ./frontend/package-lock.json ./
RUN npm ci
COPY ./frontend ./
RUN npm run build

FROM golang:alpine
RUN mkdir -p /location-app
WORKDIR /location-app

COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN go mod download

COPY ./backend ./
COPY --from=build /app/build /location-app/view

RUN go build -o ./app ./main.go


ENTRYPOINT ["./app"]

