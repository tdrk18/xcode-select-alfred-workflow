TESTDIR := testdir

$(TESTDIR):
	mkdir -p $(TESTDIR)/{data,cache}

install-lint:
	@if [ ! -f ./bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.26.0; \
	fi

.PHONY: lint
lint: install-lint
	./bin/golangci-lint run

.PHONY: test
test: $(TESTDIR)
	env alfred_workflow_bundleid=testid \
		alfred_workflow_cache=$(TESTDIR)/cache \
		alfred_workflow_data=$(TESTDIR)/data \
		go test ./... -v -cover

.PHONY: build
build:
	mkdir -p dist
	GOOS=darwin GOARCH=amd64 go build -o dist/xcode-select-alfred-workflow

.PHONY: clean
clean:
	cd dist && \
	rm -rf xcode-select-alfred-workflow xcode-select-alfred.alfredworkflow resources

.PHONY: distribute
distribute: clean build
	cd dist && \
	zip -r xcode-select-alfred.alfredworkflow .

.PHONY: version
version:
	@if [ -z ${VERSION} ]; then \
		echo "usage: make version VERSION='0.1.2'"; \
		exit 1; \
	fi
	plutil -replace version -string ${VERSION} dist/Info.plist
