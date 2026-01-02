# ---------- BUILD STAGE ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# install git (needed for go mod)
RUN apk add --no-cache git

# copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o foodos-api cmd/api/main.go


# ---------- RUNTIME STAGE ----------
FROM alpine:latest

WORKDIR /app

# copy binary from builder
COPY --from=builder /app/foodos-api .

# expose port
EXPOSE 8080

# run binary
CMD ["./foodos-api"]
