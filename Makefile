SRCS := main.go
APPNAME := clgrep

.PHONY: build cross-build archive

build:
	go build -o bin/$(APPNAME) $(SRCS)

cross-build:
	-mkdir dist/
	-rm -rf dist/*
	mkdir -p dist/linux_amd64
	mkdir -p dist/linux_386
	mkdir -p dist/darwin_amd64
	mkdir -p dist/darwin_386
	mkdir -p dist/windows_amd64
	mkdir -p dist/windows_386
	GOOS=linux GOARCH=amd64   go build -o dist/linux_amd64/$(APPNAME)       $(SRCS) 
	GOOS=linux GOARCH=386     go build -o dist/linux_386/$(APPNAME)         $(SRCS) 
	GOOS=darwin GOARCH=amd64  go build -o dist/darwin_amd64/$(APPNAME)      $(SRCS) 
	GOOS=darwin GOARCH=386    go build -o dist/darwin_386/$(APPNAME)        $(SRCS) 
	GOOS=windows GOARCH=amd64 go build -o dist/windows_amd64/$(APPNAME).exe $(SRCS) 
	GOOS=windows GOARCH=386   go build -o dist/windows_386/$(APPNAME).exe   $(SRCS) 

archive: cross-build
	-rm dist/*.tar.gz
	find dist/ -mindepth 1 -type d | while read -r d; do cp ./README.md "$d"/ ; done
	( cd dist/ && tar czf linux_amd64.tar.gz linux_amd64 )
	( cd dist/ && tar czf linux_386.tar.gz linux_386 )
	( cd dist/ && tar czf darwin_amd64.tar.gz darwin_amd64 )
	( cd dist/ && tar czf darwin_386.tar.gz darwin_386 )
	( cd dist/ && tar czf windows_amd64.tar.gz windows_amd64 )
	( cd dist/ && tar czf windows_386.tar.gz windows_386 )
