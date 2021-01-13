# `mygoapi`

This is a simple API that exposes two endpoints, one to encrypt `/api/encrypt` and another to decrypt `/api/decrypt` a value from a JSON document.

By default a random key will be created to encrypt and decrypt the value. This key will not be persistent, which means that as soon as the app is stopped, new incoming requests to decrypt previous values will not work, since they were encrypted with another key.

To provide persistency, a file containing a key can be created in `/etc/mygoapi`. The key must be a string of `32` characters.

## How to run the application

### Local execution

```
# Run app
go run cmd/mygoapi/main.go
```

### Container execution

```
# Podman build and run example

podman build -t mygoapi:latest .

podman run -p 8081:8081 localhost/mygoapi:latest

podman run -v /tmp/mygoapi:/etc/mygoapi:z -p 8081:8081 localhost/mygoapi:latest # Persistent key

# Docker build and run example

sudo docker build -t mygoapi:latest .

sudo docker run -p 8081:8081 mygoapi:latest

sudo docker run -v /tmp/mygoapi:/etc/mygoapi -p 8081:8081 mygoapi:latest # Persistent key
```

## How to use the application

This example will use `curl` command but any other tool like postman should work.

```
# Send a request to encrypt a value
curl --header "Content-Type: application/json" --request POST --data '{"value": "foo"}'   http://localhost:8081/api/encrypt

# Send a request to decrypt the value encrypted above
curl --header "Content-Type: application/json" --request POST --data '{"value": "MY ENCRYPTED VALUE"}'   http://localhost:8081/api/decrypt
```

## Running the Unit Tests

```
go test -cover ./...
```