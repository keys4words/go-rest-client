FROM golang:1.16-alpine AS build
WORKDIR /app
ADD go.mod .
COPY main.go .
RUN go build -o app main.go

FROM alpine
WORKDIR /app
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["./app"]