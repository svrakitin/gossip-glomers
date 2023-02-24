FROM golang:1.20 AS builder

WORKDIR /build

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o maelstrom-node

FROM eclipse-temurin:19

RUN apt update && \
    apt install -y bzip2 git graphviz gnuplot

WORKDIR /opt/gossip-glomers

RUN curl -sLO https://github.com/jepsen-io/maelstrom/releases/download/v0.2.2/maelstrom.tar.bz2 && \
    tar -xvf maelstrom.tar.bz2 && \
    rm maelstrom.tar.bz2

COPY --from=builder /build/maelstrom-node /opt/gossip-glomers/maelstrom-node

WORKDIR /opt/gossip-glomers/maelstrom

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "maelstrom", "test", "--bin", "/opt/gossip-glomers/maelstrom-node"]