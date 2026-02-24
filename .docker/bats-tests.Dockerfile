FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X github.com/tx3stn/pair/cmd.Version=e2e-test" -o pair

FROM bats/bats:1.13.0

RUN apk add --no-cache \
	curl \
	git \
	musl-dev \
	expect

COPY --from=builder /app/pair /usr/bin/pair

ENTRYPOINT [ "bash" ]
