SHELL := /bin/bash

PREFIX=qingstor-sdk-go
VERSION=$(shell cat version.go | grep "Version\ =" | sed -e s/^.*\ //g | sed -e s/\"//g)
DIRS_TO_CHECK=$(shell ls -d */ | grep -vE "vendor|test")
PKGS_TO_CHECK=$(shell go list ./... | grep -v "/vendor/")
PKGS_TO_RELEASE=$(shell go list ./... | grep -vE "/vendor/|/test")
FILES_TO_RELEASE=$(shell find . -name "*.go" | grep -vE "/vendor/|/test|.*_test.go")
FILES_TO_RELEASE_WITH_VENDOR=$(shell find . -name "*.go" | grep -vE "/test|.*_test.go")

.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  all               to check, build, test and release this SDK"
	@echo "  check             to vet and lint the SDK"
	@echo "  update            to update git submodules"
	@echo "  generate          to generate service code"
	@echo "  build             to build the SDK"
	@echo "  unit              to run all sort of unit tests except runtime"
	@echo "  unit-test         to run unit test"
	@echo "  unit-benchmark    to run unit test with benchmark"
	@echo "  unit-coverage     to run unit test with coverage"
	@echo "  unit-race         to run unit test with race"
	@echo "  unit-runtime      to run test with go1.7, go1.6, go 1.5 in docker"
	@echo "  test              to run service test"
	@echo "  release           to build and release current version"
	@echo "  release-source    to pack the source code"
	@echo "  clean             to clean the coverage files"

.PHONY: all
all: check build unit release

.PHONY: check
check: vet lint

.PHONY: vet
vet:
	@echo "go tool vet, skipping vendor packages"
	@go tool vet -all ${DIRS_TO_CHECK}
	@echo "ok"

.PHONY: lint
lint:
	@echo "golint, skipping vendor packages"
	@lint=$$(for pkg in ${PKGS_TO_CHECK}; do golint $${pkg}; done); \
	 lint=$$(echo "$${lint}"); \
	 if [[ -n $${lint} ]]; then echo "$${lint}"; exit 1; fi
	@echo "ok"

.PHONY: update
update:
	git submodule update --remote
	@echo "ok"

.PHONY: generate
generate:
	@if [[ ! -f "$$(which snips)" ]]; then \
		echo "ERROR: Command \"snips\" not found."; \
	fi
	snips \
		--service=qingstor --service-api-version=latest \
		--spec="./specs" --template="./template" --output="./service"
	gofmt -w .
	@echo "ok"

.PHONY: build
build:
	@echo "build the SDK"
	GOOS=linux GOARCH=amd64 go build ${PKGS_TO_CHECK}
	GOOS=darwin GOARCH=amd64 go build ${PKGS_TO_CHECK}
	GOOS=windows GOARCH=amd64 go build ${PKGS_TO_CHECK}
	@echo "ok"

.PHONY: unit
unit: unit-test unit-benchmark unit-coverage unit-race

.PHONY: unit-test
unit-test:
	@echo "run unit test"
	go test -v ${PKGS_TO_CHECK}
	@echo "ok"

.PHONY: unit-benchmark
unit-benchmark:
	@echo "run unit test with benchmark"
	go test -v -bench=. ${PKGS_TO_CHECK}
	@echo "ok"

.PHONY: unit-coverage
unit-coverage:
	@echo "run unit test with coverage"
	for pkg in ${PKGS_TO_CHECK}; do \
		output="coverage$${pkg#github.com/yunify/qingstor-sdk-go}"; \
		mkdir -p $${output}; \
		go test -v -cover -coverprofile="$${output}/profile.out" $${pkg}; \
		if [[ -e "$${output}/profile.out" ]]; then \
			go tool cover -html="$${output}/profile.out" -o "$${output}/profile.html"; \
		fi; \
	done
	@echo "ok"

.PHONY: unit-race
unit-race:
	@echo "run unit test with race"
	go test -v -race -cpu=1,2,4 ${PKGS_TO_CHECK}
	@echo "ok"

.PHONY: unit-runtime
unit-runtime: unit-runtime-go-1.8 unit-runtime-go-1.7 unit-runtime-go-1.6 unit-runtime-go-1.5

export define DOCKERFILE_GO_1_8
FROM golang:1.8

ADD . /go/src/github.com/yunify/qingstor-sdk-go
WORKDIR /go/src/github.com/yunify/qingstor-sdk-go

CMD ["make", "build", "unit"]
endef

.PHONY: unit-runtime-go-1.8
unit-runtime-go-1.8:
	@echo "run test in go 1.8"
	echo "$${DOCKERFILE_GO_1_8}" > "dockerfile_go_1.8"
	docker build -f "./dockerfile_go_1.8" -t "${PREFIX}:go-1.8" .
	rm -f "./dockerfile_go_1.8"
	docker run --name "${PREFIX}-go-1.8-unit" -t "${PREFIX}:go-1.8"
	docker rm "${PREFIX}-go-1.8-unit"
	docker rmi "${PREFIX}:go-1.8"
	@echo "ok"

export define DOCKERFILE_GO_1_7
FROM golang:1.7

ADD . /go/src/github.com/yunify/qingstor-sdk-go
WORKDIR /go/src/github.com/yunify/qingstor-sdk-go

CMD ["make", "build", "unit"]
endef

.PHONY: unit-runtime-go-1.7
unit-runtime-go-1.7:
	@echo "run test in go 1.7"
	echo "$${DOCKERFILE_GO_1_7}" > "dockerfile_go_1.7"
	docker build -f "./dockerfile_go_1.7" -t "${PREFIX}:go-1.7" .
	rm -f "./dockerfile_go_1.7"
	docker run --name "${PREFIX}-go-1.7-unit" -t "${PREFIX}:go-1.7"
	docker rm "${PREFIX}-go-1.7-unit"
	docker rmi "${PREFIX}:go-1.7"
	@echo "ok"

export define DOCKERFILE_GO_1_6
FROM golang:1.6

ADD . /go/src/github.com/yunify/qingstor-sdk-go
WORKDIR /go/src/github.com/yunify/qingstor-sdk-go

CMD ["make", "build", "unit"]
endef

.PHONY: unit-runtime-go-1.6
unit-runtime-go-1.6:
	@echo "run test in go 1.6"
	echo "$${DOCKERFILE_GO_1_6}" > "dockerfile_go_1.6"
	docker build -f "./dockerfile_go_1.6" -t "${PREFIX}:go-1.6" .
	rm -f "./dockerfile_go_1.6"
	docker run --name "${PREFIX}-go-1.6-unit" -t "${PREFIX}:go-1.6"
	docker rm "${PREFIX}-go-1.6-unit"
	docker rmi "${PREFIX}:go-1.6"
	@echo "ok"

export define DOCKERFILE_GO_1_5
FROM golang:1.5
ENV GO15VENDOREXPERIMENT="1"

ADD . /go/src/github.com/yunify/qingstor-sdk-go
WORKDIR /go/src/github.com/yunify/qingstor-sdk-go

CMD ["make", "build", "unit"]
endef

.PHONY: unit-runtime-go-1.5
unit-runtime-go-1.5:
	@echo "run test in go 1.5"
	echo "$${DOCKERFILE_GO_1_5}" > "dockerfile_go_1.5"
	docker build -f "dockerfile_go_1.5" -t "${PREFIX}:go-1.5" .
	rm -f "dockerfile_go_1.5"
	docker run --name "${PREFIX}-go-1.5-unit" -t "${PREFIX}:go-1.5"
	docker rm "${PREFIX}-go-1.5-unit"
	docker rmi "${PREFIX}:go-1.5"
	@echo "ok"

.PHONY: test
test:
	pushd "./test"; go run *.go; popd
	@echo "ok"

.PHONY: release
release: release-source release-source-with-vendor

.PHONY: release-source
release-source:
	@echo "pack the source code"
	mkdir -p "release"
	zip -FS "release/${PREFIX}-source-v${VERSION}.zip" ${FILES_TO_RELEASE}
	@echo "ok"

.PHONY: release-source-with-vendor
release-source-with-vendor:
	@echo "pack the source code"
	mkdir -p "release"
	zip -FS "release/${PREFIX}-source-with-vendor-v${VERSION}.zip" ${FILES_TO_RELEASE_WITH_VENDOR}
	@echo "ok"

.PHONY: clean
clean:
	rm -rf $${PWD}/coverage
	@echo "ok"
