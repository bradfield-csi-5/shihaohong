#!/bin/bash

print_first_sentence_article() {
  PAGE=$1

    # Assume the first <p> followed by no other tags is the first sentence of the paragraph
    search_result=$(echo "$PAGE" | ggrep -P '(?<=<p>[a-zA-Z])(.*?)(?=[.!?]\s)')
    # echo searched results
    # echo "${search_result:0:1000}"
    filtered_result=$(echo "$search_result" | sed 's/<[^>]*>//g')

    # echo filtered results
    # echo "${filtered_result:0:1000}"
    substring=". "

    search=${filtered_result%%"$substring"*}
    search_idx=$((${#search} + 1))

    # echo filtered sentence
    echo "${filtered_result:0:search_idx}"
}

# curl website
PAGE=$(curl "$1" -s)
if [ -z "${PAGE}" ]; then
  echo "Error fetching web page"
  exit 1
fi

# if no third argument:
    # regex and grab the first sentence of the article
    # regex and grab the list of section headings
if  [ -z $2 ]; then

  print_first_sentence_article "$PAGE"
# if third argument:
    # regex and grab the first sentence of the section
    # grab all subsection headings
elif [ -z $3 ]; then
  echo TODO
fi
