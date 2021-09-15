#!/bin/bash

export HELM_EXPERIMENTAL_OCI=1 

aws ecr get-login-password \
     --region us-east-2 | helm registry login \
     --username AWS \
     --password-stdin $ECR_URL
echo $modules | jq -r 'fromjson | to_entries[] | .key +" " + .value' | while IFS=' ' read -r key value; do 
  module="${key}"

  # Add entry to versions file
  echo "$key=$value" >> $VERSIONS_FILE

  # Download files from S3 only if version is not null
  if [[ -z $value ]]; then
    echo "Skipping $module"
  else
    echo "Pulling helm chart $module"
    # aws s3 cp $S3_DEV_BUCKET_URL/${moduleFullName,,} $MODULES_DIR/${module,,}/
    helm chart pull 248135293344.dkr.ecr.us-east-2.amazonaws.com/$module:${value}
    echo "Exporting helm chart $module"
    helm chart export 248135293344.dkr.ecr.us-east-2.amazonaws.com/$module:${value}
  fi
done
