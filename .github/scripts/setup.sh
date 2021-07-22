#!/bin/bash

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

# set env variables
echo ALATION_OUTPUT_DIR=alation-$VERSION >> $GITHUB_ENV
echo ALATION_OUTPUT=$ALATION_OUTPUT_DIR.tar.gz  >> $GITHUB_ENV
