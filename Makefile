TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=leanspace
NAME=leanspace
BINARY=terraform-provider-${NAME}
VERSION?=0.0.4
PLATFORM=""
ARCHITECTURE=""

ifeq ($(OS),Windows_NT)
    PLATFORM=windows
    ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
        ARCHITECTURE=amd64
    else
        ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
            ARCHITECTURE=amd64
        endif
		ifeq ($(PROCESSOR_ARCHITECTURE),ARM64)
            ARCHITECTURE=arm64
        endif
    endif
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        PLATFORM=linux
    endif
    ifeq ($(UNAME_S),Darwin)
        PLATFORM=darwin
    endif
    UNAME_P := $(shell uname -p)
    ifeq ($(UNAME_P),x86_64)
        ARCHITECTURE=amd64
    endif
    ifeq ($(UNAME_P),i386)
        ARCHITECTURE=amd64
    endif
    ifneq ($(filter aarch6%,$(UNAME_P)),)
        ARCHITECTURE=arm64
    endif
    ifeq ($(UNAME_P),arm)
        ARCHITECTURE=arm64
    endif
endif

ifeq ($(OS_ARCH),)
    OS_ARCH=${PLATFORM}_${ARCHITECTURE}
endif


default: install

build:
	go build -o ${BINARY}

build-windows:
	go build -o ${BINARY}.exe

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

install-windows: build-windows
	if not exist %APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS_ARCH} mkdir %APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS_ARCH}
	move ${BINARY}.exe %APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
