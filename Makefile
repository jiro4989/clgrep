SRCS := main.go
APPNAME := clgrep
LDFLAGS := -ldflags="-X main.version=$(VERSION)"
BUILD_CMD := go build $(LDFLAGS)

.PHONY: build cross-build archive deploy var-check

build:
	$(BUILD_CMD) -o bin/$(APPNAME) $(SRCS)

install: var-check build
	go install $(LDFLAGS)

cross-build: var-check
	-mkdir dist/
	-rm -rf dist/*
	mkdir -p dist/$(APPNAME)_linux_amd64
	mkdir -p dist/$(APPNAME)_linux_386
	mkdir -p dist/$(APPNAME)_darwin_amd64
	mkdir -p dist/$(APPNAME)_darwin_386
	mkdir -p dist/$(APPNAME)_windows_amd64
	mkdir -p dist/$(APPNAME)_windows_386
	GOOS=linux GOARCH=amd64   $(BUILD_CMD) -o dist/$(APPNAME)_linux_amd64/$(APPNAME)       $(SRCS) 
	GOOS=linux GOARCH=386     $(BUILD_CMD) -o dist/$(APPNAME)_linux_386/$(APPNAME)         $(SRCS) 
	GOOS=darwin GOARCH=amd64  $(BUILD_CMD) -o dist/$(APPNAME)_darwin_amd64/$(APPNAME)      $(SRCS) 
	GOOS=darwin GOARCH=386    $(BUILD_CMD) -o dist/$(APPNAME)_darwin_386/$(APPNAME)        $(SRCS) 
	GOOS=windows GOARCH=amd64 $(BUILD_CMD) -o dist/$(APPNAME)_windows_amd64/$(APPNAME).exe $(SRCS) 
	GOOS=windows GOARCH=386   $(BUILD_CMD) -o dist/$(APPNAME)_windows_386/$(APPNAME).exe   $(SRCS) 

archive: cross-build
	-rm dist/*.tar.gz
	find dist/ -mindepth 1 -type d | while read -r d; do cp ./README.md ./changelog "$$d"/ ; done
	( cd dist/ && tar czf $(APPNAME)_linux_amd64.tar.gz   $(APPNAME)_linux_amd64 )
	( cd dist/ && tar czf $(APPNAME)_linux_386.tar.gz     $(APPNAME)_linux_386 )
	( cd dist/ && tar czf $(APPNAME)_darwin_amd64.tar.gz  $(APPNAME)_darwin_amd64 )
	( cd dist/ && tar czf $(APPNAME)_darwin_386.tar.gz    $(APPNAME)_darwin_386 )
	( cd dist/ && tar czf $(APPNAME)_windows_amd64.tar.gz $(APPNAME)_windows_amd64 )
	( cd dist/ && tar czf $(APPNAME)_windows_386.tar.gz   $(APPNAME)_windows_386 )

deploy: archive
	ghr $(VERSION) dist/

var-check:
	if [ -z $(VERSION) ]; then echo Require VERSION. ; exit 1; fi
