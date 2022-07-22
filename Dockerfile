FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY api ./api
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /admin .

##
## Deploy
##
FROM alpine
WORKDIR /
COPY --from=build /admin .
ENV PORT=8080
ENTRYPOINT [ "./admin"]
