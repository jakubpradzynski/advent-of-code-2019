#!/bin/bash

DIR_NAME=$1
mkdir "${DIR_NAME}"
mkdir "${DIR_NAME}/src"
mkdir "${DIR_NAME}/src/main"
mkdir "${DIR_NAME}/src/resources"
cp ~/Downloads/input.txt "${DIR_NAME}/src/resources/input.txt"
touch "${DIR_NAME}/README.md"
