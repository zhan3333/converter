.PHONY: build
build:
	#CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o build/converter.exe cmd/gui/main.go
	go build -o build/converter cmd/gui/main.go

release:
	rm -f converter.zip build/logs/*.txt build/*.mp4
	zip -p -r converter.zip build/

run:
	go run cmd/gui/main.go

clear:
	rm -f build/converter.exe build/converter build/logs/*.txt build/*.mp4
	rm -f *.mp4 logs/*.txt
