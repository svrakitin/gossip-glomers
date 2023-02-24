# fly.io Gossip Glomers

## Build

```sh
docker build . -t maelstrom-node:latest
```

## Run

```sh
docker run -it \
    -v $(pwd)/maelstrom-store:/opt/gossip-glomers/maelstrom/store \
    maelstrom-node:latest \
    unique-ids \
    -w unique-ids \
    --time-limit 30 \
    --rate 1000 \
    --node-count 3 \
    --availability total \
    --nemesis partition
```

## Serve results

```sh
docker run -it \
    -p 38080:8080 \
    -v $(pwd)/maelstrom-store:/opt/gossip-glomers/maelstrom/store \
    --entrypoint /bin/sh \
    maelstrom:latest \
    maelstrom serve
```