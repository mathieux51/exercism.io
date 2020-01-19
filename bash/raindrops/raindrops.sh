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
set -o errexit
set -o nounset


main () {
    (( $1 % 3 )) || str+=Pling
    (( $1 % 5 )) || str+=Plang
    (( $1 % 7 )) || str+=Plong

    echo ${str:-$1}
}

main "$@"
