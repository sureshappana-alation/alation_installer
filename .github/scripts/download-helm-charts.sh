#!/bin/bash

export HELM_EXPERIMENTAL_OCI=1 

aws ecr get-login-password \
     --region us-east-2 | helm registry login \
     --username AWS \
     --password-stdin $ECR_URL
echo $modules | jq -r 'fromjson | to_entries[] | .key +" " + .value' | while IFS=' ' read -r key value; do 
  module="${key}"
  moduleVersion="${value}"

  # Add entry to versions file
  echo "$key=$value" >> $VERSIONS_FILE

  # Download files from S3 only if version is not null
  if [[ -z $value ]]; then
    echo "Skipping $module"
  else
    echo "Pulling helm chart $module"
    # aws s3 cp $S3_DEV_BUCKET_URL/${moduleFullName,,} $MODULES_DIR/${module,,}/
    helm pull $HELM_REGISTRY_URL/$module/helm --version $moduleVersion --untar --untardir $MODULES_DIR/$module/charts
    # move helm chart files to parent and remove the module directory
    mv $MODULES_DIR/$module/charts/$module/*  $MODULES_DIR/$module/charts && rm -rf $MODULES_DIR/$module/charts/$module
  fi
done
