# hello 👋

This is a minimal web server written in Go. Used for testing purposes in containerized environments.

## Usage

```sh
$ go run main.go
```

## Build docker image

```sh
# Build and tag the image
$ docker build -t juliantroeps/hello .
$ docker tag juliantroeps/hello juliantroeps/hello:latest
# Push the image to docker hub
$ docker push juliantroeps/hello

# Or build for a different architecture
$ docker buildx  build --platform linux/arm64 -t juliantroeps/hello .
$ docker tag juliantroeps/hello juliantroeps/hello:latest
$ docker push juliantroeps/hello
```

Run docker container with:

```sh
$ docker run -p 8080:8080 juliantroeps/hello
```
