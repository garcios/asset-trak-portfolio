# Podman

## Install via brew
```shell
brew install podman
```

## Start podman
```shell
podman machine init
podman machine start
```

## Verify
```shell
podman info
```

## Run docker image
```shell
podman run hello-world
```

## How install redis
```shell
podman run -p 6379:6379 --name redis -e ALLOW_EMPTY_PASSWORD=yes bitnami/redis:latest
```