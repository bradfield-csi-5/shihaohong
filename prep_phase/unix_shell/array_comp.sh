#!/bin/bash

# In this exercise, you will need to compare three list of arrays and write the common elements of all the three arrays:
a=(3 5 8 10 6)
b=(6 5 4 12)
c=(14 7 5 7)

for i in "${a[@]}"
do
    for j in "${b[@]}"
    do
        for k in "${c[@]}"
        do
            if [[ i -eq j && j -eq k ]] ; then
                echo "$i"
            fi
        done
    done
done
