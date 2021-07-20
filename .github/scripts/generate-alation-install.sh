#!/bin/bash
for dir in $(find $MODULES_DIR -mindepth 1 -maxdepth 1 -type d); do
  directory=${dir%*/}

  echo "[$directory]: Started processing"
  files=$(find $directory -mindepth 1 -maxdepth 1 -iname "*.tar.gz")

  if [ "${#files}" -gt 1 ]; then
    moduleFile=${files[0]}

    echo "[$directory]: Found $moduleFile. Extracting..."
    tar -xzf $moduleFile -C $directory  --strip-components=1
    
    echo "[$directory]: Deleting $moduleFile"
    rm -f $moduleFile
  else
    echo "[$directory]: No module tar files found"
  fi

  if test -f "$directory/install.yaml"; then
    echo "[$directory]: Found install.yaml"
    echo "[$directory]: Appending install.yaml content to $INSTALL_CONFIG_FILE file"
    
    # Creating any entry for module in alation_install.yaml
    echo "${directory##*/}:" >> $INSTALL_CONFIG_FILE
    
    # Adding contents of install.yaml to alation_install.yaml with indentation
    sed -e 's/^/  /' $directory/install.yaml >> $INSTALL_CONFIG_FILE
  else
    echo "[$directory]: $directory/install.yaml" doesn\'t exists. Ignoring...
  fi
done
