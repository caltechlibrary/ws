#!/bin/bash

function checkApp() {
    APP_NAME=$1
    if [ "$APP_NAME" = "" ]; then
        echo "Missing $APP_NAME"
        exit 1
    fi
}

function softwareCheck() {
    for APP_NAME in shorthand; do
        checkApp $APP_NAME
    done
}

function mkPage () {
    nav="$1"
    content="$2"
    html="$3"

    echo "Rendering $html from $content and $nav"
    shorthand \
        -e "{{navContent}} :import-markdown: $nav" \
        -e "{{pageContent}} :import-markdown: $content" \
        page.shorthand > $html
}

echo "Checking necessary software is installed"
softwareCheck
echo "Generating website index.html with shorthand"
mkPage nav.md README.md index.html
echo "Generating install.html with shorthand"
mkPage nav.md INSTALL.md installation.html
