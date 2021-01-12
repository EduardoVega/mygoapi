FROM golang:1.14-alpine AS build

WORKDIR /src

COPY . /src

RUN go build -o /bin/mygoapi cmd/mygoapi/main.go

FROM alpine

COPY --from=build /bin/mygoapi /

ENTRYPOINT ["/mygoapi"]