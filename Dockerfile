FROM golang:1.21-alpine AS build

# Install CA certificates on the Alpine-based build image
RUN apk add --no-cache ca-certificates

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o go-htmx .

# Final Stage
FROM busybox

# Metadata as described by the OCI image spec
LABEL org.opencontainers.image.title="MyApp" \
      org.opencontainers.image.description="A simple Go application" \
      org.opencontainers.image.version="1.0"

# Copy the binary from the build stage
COPY --from=build /src/go-htmx /go-htmx
COPY --from=build /src/assets /assets

# Copy CA certificates from the build stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV GIN_MODE=release
ENV HOST=0.0.0.0:8069
ENV ENV=production
ENV BASE_URL=https://go-htmx.fly.dev

ENTRYPOINT ["/go-htmx"]
