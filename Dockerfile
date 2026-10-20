# Build stage
FROM golang:1.27.1-alpine AS build

WORKDIR /src

# Cache module downloads separately from source changes.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/simple-webhook .

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/simple-webhook /simple-webhook

ENV GIN_MODE=release \
    PORT=8080
EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/simple-webhook"]
