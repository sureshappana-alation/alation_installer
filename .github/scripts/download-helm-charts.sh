#!/bin/bash

echo "Printing env"
# printenv
bash --version

# export HELM_EXPERIMENTAL_OCI=1

# aws ecr get-login-password \
#      --region us-east-2 | helm registry login \
#      --username AWS \
#      --password-stdin $ECR_URL

# for i in $modules; do
#   module="${i}"

#   echo "aa version is: " ${alation-analytics}
#   # Add entry to versions file
#   echo "$module=${!module}" >> $VERSIONS_FILE

#   # Download files from S3 only if version is not null
#   if [[ -z ${!module} ]]; then
#     echo "Skipping $module"
#   else
#     moduleFullName="${module}-${!i}.tar.gz"
#     echo "Pulling helm chart $module"
#     # aws s3 cp $S3_DEV_BUCKET_URL/${moduleFullName,,} $MODULES_DIR/${module,,}/
#     helm chart pull 248135293344.dkr.ecr.us-east-2.amazonaws.com/$module:${!module}
#     echo "Exporting helm chart $module"
#     helm chart export 248135293344.dkr.ecr.us-east-2.amazonaws.com/$module:${!module}
#   fi
# done


declare -A acronyms
acronyms[ACK]=Acknowledgement
acronyms[BGP]="Border Gateway Protocol"
acronyms[CIDR]="Classless Inter-Domain Routing"
acronyms[DHCP]="Dynamic Host Configuration Protocol"
acronyms[EOF]="End of Frame"
echo ${acronyms[ACK]}
echo ${acronyms[DHCP]}

echo ::set-output name=acronyms::${acronyms[@]}
