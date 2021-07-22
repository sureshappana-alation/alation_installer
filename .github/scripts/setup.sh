#!/bin/bash

# set env variables
echo ALATION_OUTPUT_DIR=alation-$VERSION >> $GITHUB_ENV
echo ALATION_OUTPUT=$ALATION_OUTPUT_DIR.tar.gz  >> $GITHUB_ENV
echo S3_DEV_BUCKET_URL: "s3://${{ secrets.S3_DEV_BUCKET_NAME }}" >> $GITHUB_ENV
echo S3_RELEASE_BUCKET_URL: "s3://${{ secrets.S3_RELEASE_BUCKET_NAME }}" >> $GITHUB_ENV
echo BASE_DIR: ./alation >> $GITHUB_ENV
echo RESOURCE_DIR: ./alation/res >> $GITHUB_ENV
echo MODULES_DIR: ./alation/res/modules >> $GITHUB_ENV
echo KURL_PATCH_DIR: ./alation/res/kurl_patch >> $GITHUB_ENV
echo VERSIONS_FILE: ./alation/versions.txt >> $GITHUB_ENV
echo INSTALL_CONFIG_FILE: ./alation/alation_install.yaml >> $GITHUB_ENV

echo "Create required directories"

mkdir -p $BASE_DIR
mkdir -p $MODULES_DIR
mkdir -p $KURL_PATCH_DIR

echo "Copy res content from source code to archiving directory"
cp -a $GITHUB_WORKSPACE/installer/res/. $RESOURCE_DIR/

echo "Create versions and install config file"
touch $VERSIONS_FILE $INSTALL_CONFIG_FILE


# Get current date
NOW=$(date +'%Y%m%d')

# Get branch name
BRANCH=${GITHUB_REF##*/}

VERSION=alation-k8s-$branch-$NOW.$GITHUB_RUN_NUMBER
