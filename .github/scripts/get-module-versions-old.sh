#!/bin/bash
modulesList=()
shopt -s nullglob

# Read version files
for file in ./versions/*.env; do
  while IFS='=' read -r module version
  do
    # Override application versions key name
    input_module="INPUT_$module"
    echo "$module=${!input_module:-$version}" >> $GITHUB_ENV
    modulesList+=($module)
  done < "$file"
done

IFS=',' read -r -a excludeModulesList <<< "$EXCLUDE_MODULES_STRING"

# Remove the modules in exclude list for processing
for del in ${excludeModulesList[@]}
do
   modulesList=("${modulesList[@]/$del}")
done

# Remove KURL from final modules list as kurl has special processing
modulesList=("${modulesList[@]/KURL}")
echo ${modulesList[@]}

echo ::set-output name=modulesList::${modulesList[@]}
