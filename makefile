.PHONY: build
build: clear
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/converter.exe cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/converter cmd/main.go

release: build
	rm -f converter.zip
	zip -p -r converter.zip build/

run:
	go run cmd/main.go

clear:
	rm -f build/converter.exe build/converter build/logs/*.txt build/*.mp4
	rm -f *.mp4 logs/*.txt
