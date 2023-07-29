#!/bin/sh
file() {
    echo $#
}

# Do not change anything
if [ ! $# -lt 1 ]; then
    file "$@"
    exit 0
fi
