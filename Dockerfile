ARG GO_VERSION=1.17
ARG APP_NAME="conduit"
ARG PORT=8777

FROM golang:${GO_VERSION} AS builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -ldflags="-s -w" \
    -installsuffix 'static' \
    -o /app .

FROM scratch AS final

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app /app

EXPOSE ${PORT}

ENTRYPOINT ["/app"]
