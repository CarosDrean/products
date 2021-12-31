FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN cd cmd/api && \
    go build -o ./api

FROM alpine:3.14.2
WORKDIR /app
COPY --from=build /app/cmd/api .
EXPOSE 8000
CMD /app/api
