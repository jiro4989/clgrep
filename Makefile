SRCS := main.go
APPNAME := clgrep

.PHONY: build cross-build archive deploy

build:
	go build -o bin/$(APPNAME) $(SRCS)

cross-build:
	-mkdir dist/
	-rm -rf dist/*
	mkdir -p dist/$(APPNAME)_linux_amd64
	mkdir -p dist/$(APPNAME)_linux_386
	mkdir -p dist/$(APPNAME)_darwin_amd64
	mkdir -p dist/$(APPNAME)_darwin_386
	mkdir -p dist/$(APPNAME)_windows_amd64
	mkdir -p dist/$(APPNAME)_windows_386
	GOOS=linux GOARCH=amd64   go build -o dist/$(APPNAME)_linux_amd64/$(APPNAME)       $(SRCS) 
	GOOS=linux GOARCH=386     go build -o dist/$(APPNAME)_linux_386/$(APPNAME)         $(SRCS) 
	GOOS=darwin GOARCH=amd64  go build -o dist/$(APPNAME)_darwin_amd64/$(APPNAME)      $(SRCS) 
	GOOS=darwin GOARCH=386    go build -o dist/$(APPNAME)_darwin_386/$(APPNAME)        $(SRCS) 
	GOOS=windows GOARCH=amd64 go build -o dist/$(APPNAME)_windows_amd64/$(APPNAME).exe $(SRCS) 
	GOOS=windows GOARCH=386   go build -o dist/$(APPNAME)_windows_386/$(APPNAME).exe   $(SRCS) 

archive: cross-build
	-rm dist/*.tar.gz
	find dist/ -mindepth 1 -type d | while read -r d; do cp ./README.md "$$d"/ ; done
	( cd dist/ && tar czf $(APPNAME)_linux_amd64.tar.gz   $(APPNAME)_linux_amd64 )
	( cd dist/ && tar czf $(APPNAME)_linux_386.tar.gz     $(APPNAME)_linux_386 )
	( cd dist/ && tar czf $(APPNAME)_darwin_amd64.tar.gz  $(APPNAME)_darwin_amd64 )
	( cd dist/ && tar czf $(APPNAME)_darwin_386.tar.gz    $(APPNAME)_darwin_386 )
	( cd dist/ && tar czf $(APPNAME)_windows_amd64.tar.gz $(APPNAME)_windows_amd64 )
	( cd dist/ && tar czf $(APPNAME)_windows_386.tar.gz   $(APPNAME)_windows_386 )

deploy: archive
	ghr v1.0.2 dist/
