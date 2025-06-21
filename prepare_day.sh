#!/bin/bash

TMPDIR=TMPDIR

if [ -z "$1" ]
then
  echo "No argument supplied. Please add the number of the day you'd like to
  create, such as \"02\""
  exit 1
fi

DAY=$1

cp -rT template $DAY
mv $DAY/template.go $DAY/$DAY.go
mv $DAY/template_test.go $DAY/$DAY_test.go
find $DAY -type f -print0 | xargs -0 sed -i "s/template/$DAY/g"

curl -b .session "https://adventofcode.com/2023/day/$((10#$DAY))/input" -o "$DAY/input.txt"
