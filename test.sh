#!/bin/bash

applicationList=()
shopt -s nullglob
for i in ./versions1/*.sh; do
while read line || [ -n "$line" ];
do
    set -f; IFS="="; application=($line)
    if [[ "${application[0],,}" != *"KURL"* ]];then
        applicationList+=(${application[0]})
    fi
done < "${i}"
cat "$i"
echo
done
echo $applicationList

# OCF=1000
# INPUT_OCF=100
# AA=10
# arr=(OCF AA)
# for i in "${arr[@]}"
# do
#     # echo $i
#     # echo "${!i}"
#     # val=$i
#     # echo "${!val}"

#     # input_val="INPUT_$i"
#     # echo "$i=${!input_val:-${!i}}"
#     input_key="INPUT_$i"
#     input_val=${!input_key}
#     if [ ! -z $input_val ]; then
#         echo "$i=$input_val"
#     fi
# done

# IFS='='
# shopt -s nullglob
# for i in ./versions-not-usin1g/*.sh; do
#     cat "$i"
# done

# exists=$(aws s3api list-objects-v2 --bucket unified-installer-build-pipeline-dev --query "contains(Contents[].Key, 'kurl-d2b213e.tar.gz')" || false)
# if $exists; then
# echo "Kurl package kurl-d2b213e.tar.gz exists in S3. Downloading..."
# # aws s3 cp ${{ env.S3_DEV_BUCKET_URL }}/${{env.KURL_FILE_NAME}} ${{ env.RESOURCE_DIR }}/${{ env.KURL_FILE_NAME }}            
# else
# echo "Kurl package kurl-d2b213e.tar.gz does not exists in S3. Fetching it from Kurl.sh repo..."
# # curl -LO https://k8s.kurl.sh/bundle/${{ env.KURL }}
# # mv ${{ env.KURL }} ${{ env.BASE_DIR }}/${{env.KURL_FILE_NAME}}
# # echo "Uploading Kurl Package ${{ env.BASE_DIR}}/${{ env.KURL_FILE_NAME }} to S3"
# # aws s3 cp ${{ env.BASE_DIR }}/${{ env.KURL_FILE_NAME }} ${{ env.S3_DEV_BUCKET_URL }}/${{ env.KURL_FILE_NAME }}
# fi
# # # exists=$(aws s3api list-objects-v2 --bucket unified-installer-build-pipeline-dev --query "contains(Contents[].Key, 'alastion_analytics-0.1.0.tar.gz')" || false)
# # # exists=$(aws s3api head-object --bucket unified-installer-build-pipeline-dev --key kurl-d2b213e.tar.gz || false)
# # # exists=$(aws s3api head-object --bucket unified-installer-build-pipeline-dev --key alation_analystics-0.1.0.tar.gz 2>&1)
# # exists=$(aws s3 cp s3://unified-installer-build-pipeline-dev/kurl-d2b213e.tar.gz . 2>&1)
# # # echo $exists
# # if [[ $exists == *"404"* ]]; then
# # # if $exists; then
# #     echo "not exists"
# #     echo "Kurl Package ${{ env.KURL_FILE_NAME }} does not exists in S3. Downloading it"
# #     # cd ${{ env.RESOURCE_DIR }}
# #     # curl -LO https://k8s.kurl.sh/bundle/${{ env.KURL_FILE_NAME }}
# #     # echo "Uploading Kurl Package ${{ env.KURL_FILE_NAME }} to S3"
# #     # aws s3 cp ${{env.KURL_FILE_NAME}} ${{ env.S3_DEV_BUCKET_URL }}/${{env.KURL_FILE_NAME}}
# # else
# #     echo "exists"
# # fi


# # # ALATION_ANALYTICS=1.0.0
# # # OCF=1.0.0

# # # SERVICES=("ALATION_ANALYTICS"  "OCF")


# # # for service in "${SERVICES[@]}"
# # # do
# # #    echo "$i"
# # #    # or do whatever with individual element of the array
# # # done
# # # declare -A arr
# # # IFS='='
# # # for i in ./versions/*.sh; do
# # #     while read line || [ -n "$line" ];
# # #     do
# # #         # echo $line
# # #         read -a strarr <<< "$line"
# # #         echo ${strarr[1] }
# # #     done < "${i}"
# # #  done
# # # export ALATION_ANALYTICS=0.1.0

# # # ARRAY=()
# # # for i in ./versions/*.sh; do
# # # while read line || [ -n "$line" ];
# # # do
# # #     # ARRAY+=($(echo $line | awk '{split($0,a,"="); printf a[1]}'))
# # #     # IFS="=" read -a myarray <<< $line
    
# # #     set -f; IFS="="; arr=($line)
# # #     echo "${arr[@]}"
# # # done < "${i}"
# # # done
# # # echo "${ARRAY[@]}"

# # # for i in "${!ARRAY[@]}"; do
# # #     # echo "aaa $i version is: $($i)"
# # #     echo "index: $i, value: ${ARRAY[$i]}"
# # # done

# # # hashmap["key"]="value"
# # # echo "${hashmap["key"]}"
# # # hashmap["key2"]="value2"
# # # echo "${hashmap["key"]}"
# # # echo "${hashmap["key2"]}"
# # # echo "${hashmap[@]}"
# # # # for key in ${!hashmap[@]}; do echo $key; done
# # # # for value in ${hashmap[@]}; do echo $value; done
# # # # echo hashmap has ${#hashmap[@]} elements
# # # for key in "${!hashmap[@]}"; do
# # #     echo "$key ${hashmap[$key]}"
# # # done

# # # arr["key1"]=val1

# # # arr+=( ["key2"]=val2 ["key3"]=val3 )
# # # for key in ${!arr[@]}; do
# # #     echo ${key} ${arr[${key}]}
# # # done


# # # applicationList=()
# # # applicationVersions=()
# # # for i in ./versions/*.sh; do
# # # while read line || [ -n "$line" ];
# # # do
# # #     set -f; IFS="="; service=($line)
# # #     applicationList+=(${service[0]})
# # #     applicationVersions+=(${service[1]})
# # # done < "${i}"
# # # cat "$i"
# # # echo
# # # done
# # # echo "Values"
# # # echo "${applicationList[@]}"
# # # echo "${applicationVersions[@]}"
# # # BASE_DIR="/Users/suresh.appana/Desktop/Temp/merge"
# # # touch $BASE_DIR/install-config.yaml
# # # # cd $BASE_DIR
# # # # for dir in $(ls -d  $BASE_DIR/*/ ); do
# # # for dir in $(find $BASE_DIR  -mindepth 1 -maxdepth 1 -type d ); do
# # # directory=${dir%*/}
# # # echo "dir is $directory"
# # # # pattern="$directory-*.tar.gz"
# # # # files=( $pattern )

# # # # done
# # # files=$(find $directory  -mindepth 1 -maxdepth 1 -iname "*.tar.gz")
# # # # files=$(ls $directory/*.tar.gz)
# # # # files=( $directory/*.tar.gz )
# # # if [ "${#files}" -gt 1 ]; then
# # #     echo "Found ${files[0]}"
# # #     moduleFile=${files[0]}
# # #     echo "Extracting module $moduleFile"
# # #     tar -xzf $moduleFile -C $directory  --strip-components=2
# # #     if test -f "$directory/install.yaml"; then
# # #       echo "$directory/install.yaml" t exists
# # #       echo "${directory##*/}:" >> $BASE_DIR/install-config.yaml
# # #       sed -e 's/^/  /' $directory/install.yaml >> $BASE_DIR/install-config.yaml
# # #     else
# # #       echo "$directory/install.yaml" doesn\'t exists
# # #     fi
# # # fi

# # # # if [[ ${files[0]} =~ [*] ]] ; then
# # # #    echo "No files not found"
# # # # fi
# # # # echo "${files[0]##*/}"
# # # module="${files[0]##*/}"
# # # # moduleFile="/Users/suresh.appana/Desktop/Temp/merge/$module"

# # # # if test -f "$directory/$module"; then
# # # #     echo "Extracting module $moduleFile"
# # # #     tar -tf $directory/$module
# # # #     # if test -f "$moduleFile/install.yaml"; then

# # # #     # else
# # # #     # echo "$moduleFile/install.yaml" doesn't exists
# # # #     # fi
# # # # else
# # # #     echo "$module not exists"
# # # # fi
# # # done
