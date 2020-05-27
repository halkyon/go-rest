FROM golang:1.14 AS builder
RUN apt-get update && apt-get -y install upx
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-w -s" -o ./app
RUN upx -qq -9 ./app
RUN useradd app

FROM scratch
WORKDIR /bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/app ./app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder --chown=app:app /tmp /tmp
USER app
CMD ["/app"]
