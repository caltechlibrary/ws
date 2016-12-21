#
# Biuld the project.
#
PROJECT = ws

VERSION = $(shell grep 'Version = ' $(PROJECT).go | cut -d \" -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build: ws.go cmds/ws/ws.go
	env CGO_ENABLE=0 go build -o bin/ws cmds/ws/ws.go

lint:
	gocyclo -over 12 ws.go
	gocyclo -over 12 cmds/ws/ws.go
	gofmt -w ws.go && golint ws.go
	gofmt -w cmds/ws/ws.go && golint cmds/ws/ws.go

install: bin/ws ws.go
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/ws/ws.go

clean: 
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then /bin/rm $(PROJECT)-$(VERSION)-release.zip; fi

test:
	go test

website:
	./mk-website.bash

save:
	./mk-website.bash
	git commit -am "Quick save"
	git push origin $(BRANCH)

publish:
	./mk-website.bash
	./publish.bash

release:
	./mk-release.bash
