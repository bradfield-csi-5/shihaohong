#!/bin/bash

TMP_FILENAME="out.txt"
WIKIPEDIA_API_URL="https://en.wikipedia.org/w/api.php"

remove_html_tags() {
    # sed to search for all <*> and remove them
    echo "$1" | sed 's/<[^>]*>//g'
}

print_first_sentence_article() {
    curl "$WIKIPEDIA_API_URL?format=json&action=query&prop=extracts&exsentences=1&explaintext&titles=$1" -s > section.txt
    search_result=$(grep -o '"extract":".*"' section.txt | awk -F'"' '{print $4}')
    # echo searched results
    # echo "${search_result:0:1000}"

    # echo filtered sentence
    echo "$search_result"
}

print_section_headings_article() {
    # grep to search for the section headings
    ggrep -oP '(?<=<span class="mw-headline")(.*?)(?=</span)' "$1" | cut -d '>' -f2
}

# curl website
curl "https://en.wikipedia.org/wiki/$1" -s > $TMP_FILENAME

if [ ! -f $TMP_FILENAME ]; then
    echo Error fetching web page
    exit 1
fi

# if no third argument:
    # regex and grab the first sentence of the article
    # regex and grab the list of section headings
if  [ -z "$2" ]; then
    echo "---Title of Wikipedia Article:---"
    echo "$1"
    echo "---First sentence:---"
    print_first_sentence_article "$1"
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
