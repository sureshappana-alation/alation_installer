#!/bin/bash
set -e

EXCLUDE_MODULES=$(echo $INPUT_CONTEXT | jq '.EXCLUDE_MODULES')

# Remove quotes from the beginning and ending
EXCLUDE_MODULES="${EXCLUDE_MODULES%\"}"
EXCLUDE_MODULES="${EXCLUDE_MODULES#\"}"

IFS=',' read -r -a excludeModulesList <<< "$EXCLUDE_MODULES"

FORMATTED_EXCLUDE_MODULES=".\"EXCLUDE_MODULES\","

echo "Formatted modules list: $FORMATTED_EXCLUDE_MODULES"
# Generate the module list to be excluded
for module in ${excludeModulesList[@]}
do
  echo "[$module]: started processing"
  module=`echo $module | xargs`
  FORMATTED_EXCLUDE_MODULES+=".\""$module"\","
  echo "Formatted modules list: $FORMATTED_EXCLUDE_MODULES"
  echo "[$module]: Added to excluded list"
done

FORMATTED_EXCLUDE_MODULES="${FORMATTED_EXCLUDE_MODULES%,}"

echo "Reading version json files"
# Merge version files and override versions
MODULE_VERSIONS=$(jq -s add versions/*.json)
echo "Merged version json files: $MODULE_VERSIONS"

echo "Applying override version on base versions"
# Override versions
MODULE_VERSIONS=$(echo $MODULE_VERSIONS $INPUT_CONTEXT | jq -s add)
echo "Override versions applied: $MODULE_VERSIONS"

echo "Removing excluded modules list from modules list"
# Remove the excluded modules from final version list
MODULE_VERSIONS=$(echo $MODULE_VERSIONS | jq 'del('$FORMATTED_EXCLUDE_MODULES')')
echo "Excluded modules from EXCLUDE_MODULES: $MODULE_VERSIONS"

MODULE_VERSIONS=$(echo $MODULE_VERSIONS | jq @json)

echo ::set-output name=modules::${MODULE_VERSIONS[@]}
