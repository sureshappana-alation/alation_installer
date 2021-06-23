#!/bin/bash
# ALATION_ANALYTICS=1.0.0
# OCF=1.0.0

# SERVICES=("ALATION_ANALYTICS"  "OCF")


# for service in "${SERVICES[@]}"
# do
#    echo "$i"
#    # or do whatever with individual element of the array
# done
# declare -A arr
# IFS='='
# for i in ./versions/*.sh; do
#     while read line || [ -n "$line" ];
#     do
#         # echo $line
#         read -a strarr <<< "$line"
#         echo ${strarr[1] }
#     done < "${i}"
#  done


ARRAY=()
for i in ./versions/*.sh; do
while read line || [ -n "$line" ];
do
    ARRAY+=($(echo $line | awk '{split($0,a,"="); printf a[1]}'))
    # ARRAY+=(${a[1]})
done < "${i}"
done
echo "${ARRAY[@]}"