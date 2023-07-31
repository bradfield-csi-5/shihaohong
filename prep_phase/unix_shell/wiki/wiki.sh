#!/bin/bash

TMP_FILENAME="out.txt"

remove_html_tags() {
    # sed to search for all <*> and remove them
    echo "$1" | sed 's/<[^>]*>//g'
}

print_title_article() {
    # grep to search for the main title
    search_result=$(grep '<span class="mw-page-title-main">.*</span>' "$1")
    # echo title result
    # echo $search_result

    filtered_result=$(remove_html_tags "$search_result")

    # trim whitespaces
    read -r filtered_result <<< "$filtered_result"
    # echo filtered result
    echo "$filtered_result"
}

print_first_sentence_article() {
    # Assume the first <p> followed by no other tags is the first sentence of the paragraph
    # TODO(shihaohong): this filter doesn't actually work all the time. Examples:
    # - <p><b> Malaysia ...
    # - <p> tags in tables preceding the first sentence, usually on the right side of the wiki article
    # idea: loosely look for <b> within a <p> tag? that's usually the first sentence highlights the search term
    search_result=$(grep '<p>[a-zA-Z].*[.!?] ' "$1")
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

print_section_headings_article() {
    # grep to search for the section headings
    ggrep -oP '(?<=<span class="mw-headline")(.*?)(?=</span)' "$1" | cut -d '>' -f2
}

# curl website
curl "$1" -s > $TMP_FILENAME

if [ ! -f $TMP_FILENAME ]; then
    echo Error fetching web page
    exit 1
fi

# if no third argument:
    # regex and grab the first sentence of the article
    # regex and grab the list of section headings
if  [ -z "$2" ]; then
    echo "---Title of Wikipedia Article:---"
    print_title_article "$TMP_FILENAME"
    echo "---First sentence:---"
    print_first_sentence_article "$TMP_FILENAME"
    echo "---List of Sections:---"
    print_section_headings_article "$TMP_FILENAME"
# if third argument:
    # regex and grab the first sentence of the section
    # grab all subsection headings
elif [ -z "$3" ]; then
    echo TODO
else
    echo Pass a Wikipedia URL into the first argument of the program
    echo Optionally, pass a second argument with the section to display
fi
