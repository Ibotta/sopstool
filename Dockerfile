FROM busybox

# install sops, which is required. Do it in the image because it needs to match the platform.
# happily, busybox has everything we need!
COPY sopsinstall.sh /tmp/sopsinstall.sh
RUN /tmp/sopsinstall.sh -b /bin && rm /tmp/sopsinstall.sh

COPY sopstool /bin/sopstool

ENTRYPOINT ["/bin/sopstool"]
