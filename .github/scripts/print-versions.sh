#!/bin/bash
echo "Printing1"
echo $INPUT_CONTEXT

echo "Printing2"
echo $INPUT_CONTEXT1

echo "Printing3"
echo ${INPUT_CONTEXT1[ALATIONANALYTICS]}

export x=ALATIONANALYTICS

echo "Printing4"
echo ${INPUT_CONTEXT} | jq ."$x"


