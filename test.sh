#!/bin/bash

merged=$(jq -s add versions-json/*.json)
echo $merged

override={\"alation-analytics\":\"1.0\"}

overridejson=$(echo $override | jq .)

echo $overridejson

echo $merged $overridejson | jq -s add | jq 'del(.sds)' 

