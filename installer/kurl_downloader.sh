#!/bin/bash

# this is a temporarily solution to download needed packages for the installer - build pipeline will replace this

# download kurl.sh package
curl https://alation-public-kurl-package.s3.amazonaws.com/kurl-d2b213e.tar.gz > res/kurl-d2b213e.tar.gz

# download nginx container
mkdir res/modules/nginx-ingress-controller/images
curl https://alation-public-kurl-package.s3.amazonaws.com/nginx-ingress-1.11.3.tar > res/modules/nginx-ingress-controller/images/nginx-ingress-1.11.3.tar