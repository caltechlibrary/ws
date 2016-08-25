#
# Biuld the project.
#
build:
	env CGO_ENABLE=0 go build -o bin/ws cmds/ws/ws.go

lint:
	gofmt -w ws.go && golint ws.go
	gofmt -w cmds/ws/ws.go && golint cmds/ws/ws.go
	gocyclo -over 20 .

install: bin/ws ws.go
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/ws/ws.go

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

test:
	go test

website:
	./mk-website.bash

save:
	./mk-website.bash
	git commit -am "Quick save"
	git push origin master

publish:
	./mk-website.bash
	./publish.bash

release:
	./mk-release.bash
