#!/bin/sh

cd $(mktemp -d)
curl -o sopstool.tar.gz https://oss-pkg.ibotta.com/sopstool/sopstool_linux.tar.gz
tar -xvzf sopstool.tar.gz
curl -o sops.tar.gz https://oss-pkg.ibotta.com/sops/sops_linux.tar.gz
tar -xvzf sops.tar.gz
mv sops /bin
mv sopstool /bin
rm -r $(pwd)
