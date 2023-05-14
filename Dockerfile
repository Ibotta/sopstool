FROM alpine:latest as build
ARG TARGETARCH

RUN apk --update add ca-certificates

# download appropriate sops (script gets latest)
COPY sopsinstall.sh /tmp/sopsinstall.sh
RUN sh /tmp/sopsinstall.sh -b /usr/local/bin -a $TARGETARCH

# grab appropriate sopstool binary from dist
COPY dist/sopstool_linux_$TARGETARCH/sopstool /usr/local/bin/sopstool

##########

FROM scratch

# get the root certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# get sops
COPY --from=build usr/local/bin/sops /usr/local/bin/sops
# get sopstool
COPY --from=build usr/local/bin/sopstool /usr/local/bin/sopstool

WORKDIR /work

ENTRYPOINT ["/usr/local/bin/sopstool"]
