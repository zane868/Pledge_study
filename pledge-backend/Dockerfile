FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /bin/pledge-api ./api && CGO_ENABLED=0 go build -trimpath -o /bin/pledge-sync ./schedule
FROM alpine:3.22
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=build /bin/pledge-api /bin/pledge-sync /usr/local/bin/
USER 65532:65532
EXPOSE 8080
CMD ["pledge-api"]
