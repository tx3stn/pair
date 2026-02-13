FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X github.com/tx3stn/pair/cmd.Version=vhs-demo" -o pair

FROM ghcr.io/charmbracelet/vhs:v0.10.0

RUN rm -rf /var/lib/apt/lists/* && \
	apt-get update --allow-releaseinfo-change && \
	apt-get -y install --no-install-recommends git && \
	git config --global --add safe.directory /vhs

COPY --from=builder --chmod=755 /app/pair /usr/bin/pair

ENTRYPOINT ["/usr/bin/vhs"]
