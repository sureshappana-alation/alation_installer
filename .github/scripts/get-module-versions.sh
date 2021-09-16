#!/bin/bash

# modulesList=()
# shopt -s nullglob

# # Read version files
# for file in ./versions/*.env; do
#   while IFS='=' read -r module version
#   do
#     echo "$module=$version" >> $GITHUB_ENV
#     modulesList+=($module)
#   done < "$file"
# done

# # Override versions information
# echo ${INPUT_CONTEXT} | jq -r 'to_entries[] | "\(.key)=\(.value)"' >> $GITHUB_ENV

IFS=',' read -r -a excludeModulesList <<< "$EXCLUDE_MODULES_STRING"

FORMATTED_EXCLUDE_STRING=""
# Remove the modules in exclude list for processing
for module in ${excludeModulesList[@]}
do
  module=`echo $module | xargs`
  FORMATTED_EXCLUDE_STRING+=".\""$module"\","
done
FORMATTED_EXCLUDE_STRING="${FORMATTED_EXCLUDE_STRING%,}"
echo $FORMATTED_EXCLUDE_STRING
# versions=$(jq -s add versions-json/*.json | jq 'del('$FORMATTED_EXCLUDE_STRING')' | jq @json)
versions=$(jq -s add versions-json/*.json)
if [ ! -z "${FORMATTED_EXCLUDE_STRING}" ]
then
  versions=$(echo $versions | jq 'del('$FORMATTED_EXCLUDE_STRING')')
fi
versions=$(echo $versions | jq @json)

echo $versions

echo ::set-output name=versions::${versions[@]}
