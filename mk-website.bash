#!/bin/bash

function checkApp() {
    APP_NAME=$1
    if [ "$APP_NAME" = "" ]; then
        echo "Missing $APP_NAME"
        exit 1
    fi
}

function softwareCheck() {
    for APP_NAME in mkpage; do
        checkApp $APP_NAME
    done
}

function mkPage () {
    nav="$1"
    content="$2"
    html="$3"

    echo "Rendering $html"
    mkpage \
        "nav=$nav" \
        "content=$content" \
        page.tmpl > $html
}

echo "Checking necessary software is installed"
softwareCheck
echo "Generating website index.html"
mkPage nav.md README.md index.html
echo "Generating install.html"
mkPage nav.md INSTALL.md install.html
echo "Generasting license.html"
mkPage nav.md "markdown:$(cat LICENSE)" license.html
