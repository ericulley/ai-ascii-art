fmt:
	gofmt -w -l .

build:
	go build -o bin/ascii main.go

update-permissions:
	chmod +x bin/ascii	

move-binary:
	sudo cp bin/ascii /usr/local/bin

install: build update-permissions move-binary
	@echo "Installed ascii to /usr/local/bin"
