FROM golang:1.19.4-alpine3.17 as builder-stage
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o go_network_service cmd/main.go


FROM builder-stage AS deploy-stage
FROM alpine:3.17.0
RUN apk --no-cache add ca-certificates
COPY --from=deploy-stage /app/go_network_service .
ENTRYPOINT [ "./go_network_service" ]