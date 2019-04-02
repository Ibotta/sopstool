FROM alpine:latest as build
RUN apk --update add ca-certificates

COPY sopsinstall.sh /tmp/sopsinstall.sh
RUN /tmp/sopsinstall.sh -b /usr/local/bin

##########

FROM scratch

# get the root certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# get sops
COPY --from=build usr/local/bin/sops /usr/local/bin/sops
# get sopstool
COPY sopstool /usr/local/bin/sopstool

WORKDIR /work

ENTRYPOINT ["/usr/local/bin/sopstool"]
