FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/server ./cmd/server
COPY ./assets ./assets
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build-stage /server /server
COPY --from=build-stage /app/assets /assets

EXPOSE 8080

USER nonroot:nonroot

ENV GIN_MODE=release
ENTRYPOINT ["/server"]
