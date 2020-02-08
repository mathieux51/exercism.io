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
    # $1 and $2
    # loop over $1 and get value by index on 2
    distance=0
    len=${#1}
    for ((i = 0 ; i <= len-1; i++)); do
        echo "$i ${1:i:1} ${2:i:1}"
        distance+=1

    done
    echo $distance
    # create counter/distance variable
    # print variable

}

# call main with all of the positional arguments
main "$@"

