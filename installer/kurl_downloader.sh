#!/bin/bash

# this is a temporarily solution to download needed packages for the installer - build pipeline will replace this
touch alation-install.yaml
mkdir res/modules

# download kurl.sh package
curl https://alation-public-kurl-package.s3.amazonaws.com/kurl-d2b213e.tar.gz > res/kurl.tar.gz
