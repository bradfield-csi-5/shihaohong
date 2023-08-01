#!/bin/bash
WIKIPEDIA_API_URL="https://en.wikipedia.org/w/api.php"
SECTION_FILENAME="section.txt"
SECTION_DATA_FILENAME="section_data.txt"
SECTION_LIST_FILENAME="section_list.txt"

print_first_sentence_article() {
    curl "$WIKIPEDIA_API_URL?format=json&action=query&prop=extracts&exsentences=1&explaintext&titles=$1" -s > $SECTION_FILENAME
    if [ ! -f $SECTION_FILENAME ]; then
        echo Error fetching web page
        exit 1
    fi

    grep -o '"extract":".*"' $SECTION_FILENAME | awk -F'"' '{print $4}'
}

print_first_sentence_section_and_subsection() {
    curl "$WIKIPEDIA_API_URL?action=parse&format=json&page=$1&prop=sections&disabletoc=1" -s | tr '}' '\n' > $SECTION_LIST_FILENAME
    if [ ! -f $SECTION_LIST_FILENAME ]; then
        echo Error fetching web page
        exit 1
    fi

    local EXPR
    EXPR=("\"line\":\"$2\"")
    local SECTION_INDEX
    SECTION_INDEX=$(grep "${EXPR[@]}" $SECTION_LIST_FILENAME | grep -o '"index":"[0-9]*"' | sed -e 's/"index":"//' -e 's/"$//')

    curl "$WIKIPEDIA_API_URL?action=parse&format=json&servedby=1&page=$1&prop=text&section=$SECTION_INDEX&disabletoc=1" > $SECTION_DATA_FILENAME -s
    if [ ! -f $SECTION_DATA_FILENAME ]; then
        echo Error fetching web page
        exit 1
    fi

    local RESP
    RESP=$(awk -F'<p>' '{print $2}' $SECTION_DATA_FILENAME | sed -e 's/<[^>]*>//g' -e 's/\..*//').
    echo "$RESP"

    echo "---Subsections:---"
    local NUMBER
    NUMBER=$(grep "${EXPR[@]}" $SECTION_LIST_FILENAME | grep -o '"number":"[0-9]*"' | sed -e 's/"number":"//' -e 's/"$//')
    EXPR=("\"number\":\"$NUMBER\.[0-9]*\"")
    RESP=$(grep "${EXPR[@]}" $SECTION_LIST_FILENAME | grep -o '"line":".*"' | sed -e 's/"line":"//' -e 's/",".*//')

    if [ -z "$RESP" ]; then
        echo No Subsections
    else
        echo "$RESP"
    fi
}

print_section_headings_article() {
    curl "$WIKIPEDIA_API_URL?action=parse&format=json&page=$1&prop=sections&disabletoc=1" -s | tr '}' '\n' > $SECTION_LIST_FILENAME
    if [ ! -f $SECTION_LIST_FILENAME ]; then
        echo Error fetching web page
        exit 1
    fi

    grep 'toclevel":1' $SECTION_LIST_FILENAME | grep -o '"line":.*,' |  sed -e 's/"line":"//'  -e 's/","number":".*//'
}

# TODO(shihaohong): Error handling, empty responses or invalid string values
if  [ -z "$2" ]; then
    echo "---Title of Wikipedia Article:---"
    echo "$1"
    echo "---First sentence:---"
    print_first_sentence_article "$1"
    echo "---List of Sections:---"
    print_section_headings_article "$1"
elif [ -z "$3" ]; then
    echo "---Title of Wikipedia Article:---"
    echo "$1"
    echo "---Title of Wikipedia Section:---"
    echo "$2"
    echo "---First sentence:---"
    print_first_sentence_section_and_subsection "$1" "$2"
elif [ $# -gt 3 ]; then
    echo Too many arguments. Please pass at most two arguments
    echo Current args:
    printf '"%s"\n' "$@"
else
    echo Pass a Wikipedia URL into the first argument of the program
    echo Optionally, pass a second argument with the section to display
fi
