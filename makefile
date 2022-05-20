.PHONY: build
build:
	rm -f build/converter.exe build/logs/log.txt build/*.mp4
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/converter.exe cmd/main.go

release: build
	rm -f converter-win.zip
	zip -p -r converter-win.zip build/

run:
	go run cmd/main.go