#!/bin/bash

echo "Create required directories"

mkdir -p $BASE_DIR
mkdir -p $MODULES_DIR

echo "Copy res content from source code to archiving directory"
cp -a $GITHUB_WORKSPACE/installer/res/. $RESOURCE_DIR/

echo "Create versions and install config file"
touch $VERSIONS_FILE $INSTALL_CONFIG_FILE

# Creating modules entry in alation-install.yaml
echo "modules:" >> $INSTALL_CONFIG_FILE

# Get current date
NOW=$(date +'%Y%m%d')

# Get branch name
BRANCH=${GITHUB_REF##*/}

ALATION_OUTPUT=alation-k8s-$BRANCH-$NOW.$GITHUB_RUN_NUMBER

# set env variables
echo ALATION_OUTPUT_DIR=$ALATION_OUTPUT >> $GITHUB_ENV
echo ALATION_OUTPUT=$ALATION_OUTPUT.tar.gz  >> $GITHUB_ENV
