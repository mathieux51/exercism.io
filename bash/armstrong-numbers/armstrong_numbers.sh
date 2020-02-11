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
main () {
  number=$1
  nLength=${#number}
  an=0
  for (( i=0; i<nLength; i++ )); do
    digit=${number:$i:1}
    an=$(($an + $digit**$nLength ))
  done

  if [ $an -eq $1 ]
  then
    echo 'true'
    return
  fi

  echo 'false'
  # your main function code here
}
#
#   # call main with all of the positional arguments
main "$@"
