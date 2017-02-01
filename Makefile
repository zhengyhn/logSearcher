PKGS = $(shell go list ./... | grep -v /vendor/)
PKGS_DELIM = $(shell echo $(PKGS) | sed -e 's/ /,/g')

save:
	godep save
	cp -r $(GOPATH)/src/github.com/yanyiwu/gojieba/ vendor/github.com/yanyiwu/gojieba/

build:
	go build -v -ldflags "-X model.MONGO_LOG=$(MONGO_LOG)"

dev:
	MONGO_LOG=mongodb://localhost:27017/log ./logSearcher

lint:
	golint

test:	lint
	@echo Running tests
	go list -f '{{if or (len .TestGoFiles) (len .XTestGoFiles)}} MONGO_LOG=mongodb://localhost:27017/log-test grc go test -parallel 8 -test.v -test.timeout=120s -covermode=count -coverprofile={{.Name}}_{{len .Imports}}_{{len .Deps}}.coverprofile -coverpkg $(PKGS_DELIM) {{.ImportPath}}{{end}}' $(PKGS) | xargs -I {} bash -c {}
	gocovmerge `ls *.coverprofile` > cover.out
	rm *.coverprofile

cover:
	go tool cover -html cover.out

doc:	$(PKGS)
	godoc -html $(PKGS}

.PHONY: cover test dev build
