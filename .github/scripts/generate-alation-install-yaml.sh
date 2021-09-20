#!/bin/bash

# Exit on error
set -e

for dir in $(find $MODULES_DIR -mindepth 1 -maxdepth 1 -type d); do
  directory=${dir%*/}

  echo "[$directory]: Started processing"
  # Processing install.yaml file
  if test -f "$directory/charts/install.yaml"; then
    echo "[$directory]: Found install.yaml"
    echo "[$directory]: Appending install.yaml content to $INSTALL_CONFIG_FILE file"
    
    # Creating any entry for module in alation_install.yaml
    echo "${directory##*/}:" >> $INSTALL_CONFIG_FILE
    
    # Adding contents of install.yaml to alation_install.yaml with indentation
    sed -e 's/^/  /' $directory/charts/install.yaml >> $INSTALL_CONFIG_FILE
  else
    echo "[$directory]: $directory/charts/install.yaml" doesn\'t exists. Ignoring...
  fi
done
