#
# Copyright © 2024 Eric Culley <https://github.com/ericulley>
#

fmt:
	gofmt -w -l .

build:
	go build -o bin/ascii main.go

update-permissions:
	chmod +x bin/ascii	

move-binary:
	sudo cp bin/ascii /usr/local/bin

install: build update-permissions move-binary
	@echo "Success! Installed ascii to /usr/local/bin. Run \`ascii --help\` to get started"
