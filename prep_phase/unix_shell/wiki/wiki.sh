#!/bin/bash
WIKIPEDIA_API_URL="https://en.wikipedia.org/w/api.php"
SECTION_FILENAME="section.txt"
SECTION_LIST_FILENAME="section_list.txt"

remove_html_tags() {
    # sed to search for all <*> and remove them
    echo "$1" | sed 's/<[^>]*>//g'
}

print_first_sentence_article() {
    curl "$WIKIPEDIA_API_URL?format=json&action=query&prop=extracts&exsentences=1&explaintext&titles=$1" -s > $SECTION_FILENAME
    if [ ! -f $SECTION_FILENAME ]; then
        echo Error fetching web page
        exit 1
    fi

    search_result=$(grep -o '"extract":".*"' $SECTION_FILENAME | awk -F'"' '{print $4}')
    # echo searched results
    # echo "${search_result:0:1000}"

    echo "$search_result"
}

print_section_headings_article() {
    curl "$WIKIPEDIA_API_URL?action=parse&format=json&page=$1&prop=sections&disabletoc=1" -s | tr '}' '\n' > $SECTION_LIST_FILENAME
    if [ ! -f $SECTION_LIST_FILENAME ]; then
        echo Error fetching web page
        exit 1
    fi
    # grep to search for the section headings
    grep 'toclevel":1' $SECTION_LIST_FILENAME | grep -o '"line":.*,' |  sed -e 's/"line":"//'  -e 's/","number":".*//'
}

# if no third argument:
    # regex and grab the first sentence of the article
    # regex and grab the list of section headings
if  [ -z "$2" ]; then
    echo "---Title of Wikipedia Article:---"
    echo "$1"
    echo "---First sentence:---"
    print_first_sentence_article "$1"
    echo "---List of Sections:---"
    print_section_headings_article "$1"
# if third argument:
    # regex and grab the first sentence of the section
    # grab all subsection headings
elif [ -z "$3" ]; then
    echo TODO
else
    echo Pass a Wikipedia URL into the first argument of the program
    echo Optionally, pass a second argument with the section to display
fi
