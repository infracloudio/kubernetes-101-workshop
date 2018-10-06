FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o goapp

FROM alpine:3.4
WORKDIR /app
RUN apk add --update ca-certificates
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp
