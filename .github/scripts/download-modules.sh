#!/bin/bash

for i in $modules; do
  module="${i}"

  # Add entry to versions file
  echo "$module=${!module}" >> $VERSIONS_FILE

  # Download files from S3 only if version is not null and skip
  if [[ ${!module} = "skip" ]] || [[ -z ${!module} ]]; then
    echo "Skipping $module"
  else
    moduleFullName="${module}-${!i}.tar.gz"
    echo "Downloading $moduleFullName from S3"
    aws s3 cp $S3_DEV_BUCKET_URL/${moduleFullName,,} $MODULES_DIR/${module,,}/
  fi
done
