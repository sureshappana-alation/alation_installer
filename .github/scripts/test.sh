#!/bin/bash

version=1.0.0-187
major=`echo $version | cut -d. -f1`
minor=`echo $version | cut -d. -f2`

if [[ $version == *"-"* ]]
then
  patch_and_build=`echo $version | cut -d. -f3`
  patch=`echo $patch_and_build | cut -d- -f1`
  build=`echo $patch_and_build | cut -d- -f2`
else
  patch=`echo $version | cut -d. -f3`
  build=`echo $version | cut -d. -f4`
fi

echo $major
echo $minor
echo $patch
echo $build
