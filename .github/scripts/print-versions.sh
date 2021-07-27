#!/bin/bash
echo "Printing1"
echo $INPUT_CONTEXT

echo "Printing2"
echo $INPUT_CONTEXT1

echo "Printing2"
echo ${INPUT_CONTEXT1[ALATIONANALYTICS]}

export x=ALATIONANALYTICS

echo "Printing2"
echo ${INPUT_CONTEXT1.ALATIONANALYTICS} | jq ."$x"


