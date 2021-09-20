#!/bin/bash

EXCLUDE_MODULES=$(echo ${INPUT_CONTEXT} | jq '.EXCLUDE_MODULES')

IFS=',' read -r -a excludeModulesList <<< "$EXCLUDE_MODULES"

FORMATTED_EXCLUDE_MODULES=".EXCLUDE_MODULES,"

# Remove the modules in exclude list from modules list
for module in ${excludeModulesList[@]}
do
  module=`echo $module | xargs`
  FORMATTED_EXCLUDE_MODULES+=".\""$module"\","
done

FORMATTED_EXCLUDE_MODULES="${FORMATTED_EXCLUDE_MODULES%,}"

echo $FORMATTED_EXCLUDE_MODULES

# Merge version files
versions=$(jq -s add versions-json/*.json)

echo "base versions from versions/*.json: $versions"

# Override versions
versions=$(echo $versions $INPUT_CONTEXT | jq -s add)

echo "versions after overriding: $versions"

if [ ! -z "${FORMATTED_EXCLUDE_MODULES}" ]
then
  versions=$(echo $versions | jq 'del('$FORMATTED_EXCLUDE_MODULES')')
fi

# Final versions
echo "final versions: $versions"

versions=$(echo $versions | jq @json)


echo $versions

echo ::set-output name=versions::${versions[@]}
