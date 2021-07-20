#!/bin/bash

modulesList=()
shopt -s nullglob

# Read version files
for file in ./versions/*.env; do
  while IFS='=' read -r key val
  do
    # Ignore KURL from modules list as KURL has special treatment
    if [[ $key != *"KURL"* ]]; then
      modulesList+=($key)
    fi
  done < "$file"
  cat "$file"
  echo
done  >> $GITHUB_ENV
echo ::set-output name=modulesList::${modulesList[@]}
