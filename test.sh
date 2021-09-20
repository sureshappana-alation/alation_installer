#!/bin/bash

override={\"alation-analytics\":\"1.0\"}

overridejson=$(echo $override | jq .)

echo $overridejson

merged=$(jq -s add versions-json/*.json $overridejson)
echo $merged


echo $merged $overridejson | jq -s add | jq 'del(.sds)' 

