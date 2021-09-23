#!/bin/bash

# Exit on error
set -e

echo "Create required directories"

mkdir -p $BASE_DIR
mkdir -p $MODULES_DIR

echo "Copy res content from source code to archiving directory"
cp -a $GITHUB_WORKSPACE/installer/res/. $RESOURCE_DIR/

echo "Create versions and install config file"
touch $VERSIONS_FILE $INSTALL_CONFIG_FILE

# Creating modules entry in alation-install.yaml
echo "modules:" >> $INSTALL_CONFIG_FILE

VERSION_FILE=$GITHUB_WORKSPACE/version.env

# check if file exist in repo, exit 1 if file does not exist
if [[ ! -f "$VERSION_FILE" ]]
then
    echo "ERROR: $VERSION_FILE does not exist."
    exit 1
fi

# Load version file to environment
. $VERSION_FILE

echo MAJOR=$MAJOR >> $GITHUB_ENV
echo MINOR=$MINOR >> $GITHUB_ENV
echo PATCH=$PATCH >> $GITHUB_ENV

BUILD_VERSION=$MAJOR.$MINOR.$PATCH.$GITHUB_RUN_NUMBER

# set env variables
echo BUILD_VERSION=$BUILD_VERSION >> $GITHUB_ENV
echo BUILD_OUTPUT=$BUILD_VERSION.tar.gz  >> $GITHUB_ENV
