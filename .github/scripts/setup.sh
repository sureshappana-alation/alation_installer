#!/bin/bash

echo "Create required directories"

mkdir -p ${{ env.BASE_DIR }}
mkdir -p ${{ env.MODULES_DIR }}
mkdir -p ${{ env.KURL_PATCH_DIR }}

echo "Copy res content from source code to archiving directory"
cp -a ${{ github.workspace }}/installer/res/. ${{ env.RESOURCE_DIR }}/

echo "Create versions and install config file"
touch ${{ env.VERSIONS_FILE }} ${{ env.INSTALL_CONFIG_FILE }}
