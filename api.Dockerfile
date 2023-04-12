# # Build
# FROM golang:1.19 AS build

# WORKDIR /app

# # Download Go modules
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY . ./

# RUN go build -o /priceapi


# # Deploy
# FROM alpine:latest
# WORKDIR /app

# COPY .env .
# COPY --from=build /priceapi .
# EXPOSE 8080

# ENTRYPOINT ["/app/priceapi"]

# Build
# FROM golang:alpine3.17 AS build
FROM golang:1.19.8-alpine3.17 AS build
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build ./cmd/priceapi

# Deploy
FROM alpine:latest
WORKDIR /app

# COPY .env .
COPY --from=build /app/priceapi .

ENTRYPOINT ["/app/priceapi"]