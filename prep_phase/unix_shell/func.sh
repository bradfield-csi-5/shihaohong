#!/bin/bash

# In this exercise, you will need to write a function called
# ENGLISH_CALC which can process sentences such as:

# '3 plus 5', '5 minus 1' or '4 times 6' and print the results
# as: '3 + 5 = 8', '5 - 1 = 4' or '4 * 6 = 24' respectively.

ENGLISH_CALC() {
    op=$2
    arg1=$1
    arg2=$3

    if [ "$op" = "plus" ] ; then
        echo $((arg1+arg2))
    elif [ "$op" = "minus" ] ; then
        echo $((arg1-arg2))
    elif [ "$op" = "times" ] ; then
        echo $((arg1*arg2))
    fi
}

# testing code
ENGLISH_CALC 3 plus 5
ENGLISH_CALC 5 minus 1
ENGLISH_CALC 4 times 6