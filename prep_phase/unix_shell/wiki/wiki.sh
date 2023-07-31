#!/bin/bash

remove_html_tags() {
    # sed to search for all <*> and remove them
    echo "$1" | sed 's/<[^>]*>//g'
}

print_title_article() {
    # grep to search for the main title
    search_result=$(echo "$PAGE" | ggrep -P '(?<=<span class="mw-page-title-main">)(.*?)(?=</span>)')
    # echo title result
    # echo $search_result

    filtered_result=$(remove_html_tags "$search_result")

    # trim whitespaces
    read -r filtered_result <<< "$filtered_result"
    # echo filtered result
    echo "$filtered_result"
}

print_first_sentence_article() {
    PAGE=$1

    # Assume the first <p> followed by no other tags is the first sentence of the paragraph
    # TODO(shihaohong): what happened to -P flag in newer versions of OS X?
    search_result=$(echo "$PAGE" | ggrep -P '(?<=<p>[a-zA-Z])(.*?)(?=[.!?]\s)')
    # echo searched results
    # echo "${search_result:0:1000}"
    filtered_result=$(remove_html_tags "$search_result")

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
    echo Error fetching web page
    exit 1
fi

# if no third argument:
    # regex and grab the first sentence of the article
    # regex and grab the list of section headings
if  [ -z "$2" ]; then
    print_title_article "$PAGE"
    print_first_sentence_article "$PAGE"
# if third argument:
    # regex and grab the first sentence of the section
    # grab all subsection headings
elif [ -z "$3" ]; then
    echo TODO
else
    echo Pass a Wikipedia URL into the first argument of the program
    echo Optionally, pass a second argument with the section to display
fi
