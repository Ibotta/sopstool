# fresh root certs straight from alpine at build time
FROM alpine:latest AS certs
RUN apk --update add ca-certificates

# upstream sops image; only the binary is used
FROM ghcr.io/getsops/sops:v3.13.3-alpine AS sops

FROM scratch
ARG TARGETARCH

# get the root certs
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# get sops
COPY --from=sops /usr/local/bin/sops /usr/local/bin/sops
# get sopstool binary from dist
COPY dist/sopstool_linux_${TARGETARCH}*/sopstool /usr/local/bin/sopstool

WORKDIR /work

ENTRYPOINT ["/usr/local/bin/sopstool"]
