#!/bin/bash
#
# Make releases for Linux/amd64, Linux/ARM7 (Raspberry Pi), Windows, and Mac OX X (darwin)
#
PROJECT=ws
VERSION=$(grep 'Version = ' $PROJECT.go | cut -d \" -f 2)
RELEASE_NAME="$PROJECT-$VERSION"
echo "Preparing $RELEASE_NAME-release.zip"
for PROGNAME in ws; do
    env CGO_ENABLED=0 GOOS=linux   GOARCH=amd64       go build -o dist/linux-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    env CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64       go build -o dist/macosx-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    env CGO_ENABLED=0 GOOS=windows GOARCH=amd64       go build -o dist/windows-amd64/$PROGNAME.exe cmds/$PROGNAME/$PROGNAME.go
    env CGO_ENABLED=0 GOOS=linux   GOARCH=arm GOARM=6 go build -o dist/raspberrypi-arm6/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    env CGO_ENABLED=0 GOOS=linux   GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
done
mkdir -p dist
for FNAME in README.md LICENSE INSTALL.md; do
    cp -v $FNAME dist/
done
echo "Zipping $RELEASE_NAME"
zip -r $RELEASE_NAME-release.zip dist/*
