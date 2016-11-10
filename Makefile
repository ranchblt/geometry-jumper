# https://stackoverflow.com/questions/11354518/golang-application-auto-build-versioning
# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

# https://gist.github.com/Stratus3D/a5be23866810735d7413

default: build

build: vet dev-assets
	go build -v

release: clean vet assets
	go build ${LDFLAGS} -v

linux: clean vet assets
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -v

vet:
	go vet $(go list ./... | grep -v /vendor/)

assets:
	go-bindata $(BINDATA_OPTS) -prefix "resource" -pkg resource -o resource/bindata.go resource/...

dev-assets:
	@$(MAKE) --no-print-directory assets BINDATA_OPTS="-debug"

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint .

clean:
	find . -type f -name "bindata.go" -not -path "./vendor/*" -delete
	go clean

run: clean build
	./geometry-jumper