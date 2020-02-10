#!/usr/bin/env bash

# The following comments should help you get started:
# - Bash is flexible. You may use functions or write a "raw" script.
#
# - Complex code can be made easier to read by breaking it up
#   into functions, however this is sometimes overkill in bash.
#
# - You can find links about good style and other resources
#   for Bash in './README.md'. It came with this exercise.
#
#   Example:
#   # other functions here
#   # ...
#   # ...
#
set -o errexit
set -o nounset

main () {
  words=$(echo ${1//[^a-zA-Z\']/ } | tr -s ' ')
  a=''
  for w in $words
  do
    a=$a${w:0:1}
  done
  echo ${a^^}

}
#
#   # call main with all of the positional arguments
main "$@"
