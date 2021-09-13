#!/bin/bash

aws ecr get-login-password \
     --region us-east-2 | helm registry login \
     --username AWS \
     --password-stdin ${{env.ECR_URL}}

for i in $modules; do
  module="${i}"

  # Add entry to versions file
  echo "$module=${!module}" >> $VERSIONS_FILE

  # Download files from S3 only if version is not null
  if [[ -z ${!module} ]]; then
    echo "Skipping $module"
  else
    moduleFullName="${module}-${!i}.tar.gz"
    echo "Pulling helm chart"
    # aws s3 cp $S3_DEV_BUCKET_URL/${moduleFullName,,} $MODULES_DIR/${module,,}/
    helm chart pull 248135293344.dkr.ecr.us-east-2.amazonaws.com/alation-analytics:12.17.21-906
    helm chart export 248135293344.dkr.ecr.us-east-2.amazonaws.com/alation-analytics:12.17.21-906
  fi
done
