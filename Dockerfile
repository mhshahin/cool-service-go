FROM golang:1.20.6-alpine3.18 AS build

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN GOPATH=/usr/go GODEBUG=madvdontneed=1 CGO_ENABLED=0 go build -buildvcs=false -o cool-service .

FROM alpine:3.18

COPY --from=build /app/cool-service /app/

RUN apk --no-cache update && \
   apk add --no-cache --update tzdata bash ca-certificates curl gcompat bind-tools yq jq busybox-extras zip unzip && \
   chmod +x /app/cool-service


WORKDIR /app
CMD ["/app/cool-service", "-c", "/app/config.yaml", "serve"]