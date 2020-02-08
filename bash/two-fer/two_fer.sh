#!/usr/bin/env bash

# Read this https://www.computerhope.com/unix/bash/read.htm
# Get argument from command line
# override default
set -o errexit
set -o nounset

echo "One for ${1:-you}, one for me."
