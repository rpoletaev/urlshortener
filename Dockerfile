FROM golang:alpine as builder
RUN apk update && apk add --no-cache git make ca-certificates tzdata && update-ca-certificates
RUN mkdir /urlshortener
WORKDIR /urlshortener
COPY . .
RUN make build

FROM alpine
ENV USER=runner USER_ID=1002 USER_G=runner USER_G_ID=1002
RUN addgroup -g ${USER_G_ID} ${USER_G} && \
    adduser -D --home /app -u ${USER_ID} -G ${USER_G} ${USER}
WORKDIR /app
COPY --from=builder /urlshortener/bin/urlshortener /app
# added for future if ssl will needed
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app/urlshortener"]