FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY api api
COPY models models
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /admin .

FROM alpine
WORKDIR /usr/app
COPY --from=build /admin .
ENV PORT=8080
ENTRYPOINT [ "./admin"]
