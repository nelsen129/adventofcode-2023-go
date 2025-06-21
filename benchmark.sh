#!/bin/bash

directories=$(find . -maxdepth 1 -type d -regex '\./.*[0-9]' | sort)

for directory in $directories ; do
    echo "$directory"
    go run ./$directory ./$directory/input.txt
done
