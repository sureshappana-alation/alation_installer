#!/bin/bash


EXCLUDE_MODULES=$(echo $INPUT_CONTEXT | jq '.EXCLUDE_MODULES')

IFS=',' read -r -a excludeModulesList <<< "$EXCLUDE_MODULES"

FORMATTED_EXCLUDE_MODULES=".EXCLUDE_MODULES,"

# Generate the module list to be excluded
for module in ${excludeModulesList[@]}
do
  module=`echo $module | xargs`
  FORMATTED_EXCLUDE_MODULES+=".\""$module"\","
done

FORMATTED_EXCLUDE_MODULES="${FORMATTED_EXCLUDE_MODULES%,}"


INPUT_CONTEXT_JSON=$(echo $INPUT_CONTEXT | jq .)

# Merge version files and override versions
VERSIONS=$(jq -s add versions-json/*.json $INPUT_CONTEXT_JSON)

# # Override versions
# VERSIONS=$(echo $VERSIONS $INPUT_CONTEXT | jq -s add)

# Remove the excluded modules from final version list
VERSIONS=$(echo $VERSIONS | jq 'del('$FORMATTED_EXCLUDE_MODULES')')

echo "Final versions: $VERSIONS"

VERSIONS=$(echo $VERSIONS | jq @json)

echo ::set-output name=versions::${VERSIONS[@]}
