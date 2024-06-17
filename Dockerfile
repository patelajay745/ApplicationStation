FROM golang:1.22.1-alpine AS BuildStage
WORKDIR /app
COPY . .
RUN go mod download \
    go build -o /webapp main.go
EXPOSE 3000

FROM alpine:latest
WORKDIR /
COPY --from=BuildStage /webapp /webapp
EXPOSE 3000
ENTRYPOINT [ "/webapp" ]