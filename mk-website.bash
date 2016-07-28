#!/bin/bash
#

function makePage () {
    page=$1
    nav=$2
    html_page=$3
    echo "Generating $html_page"
    shorthand \
        -e "{{content}} :import-markdown: $page" \
        -e "{{nav}} :import-markdown: $nav" \
        page.shorthand > $html_page
}

# index.html
makePage README.md nav.md index.html

# install.html (slide presentation)
makePage INSTALL.md nav.md install.html

