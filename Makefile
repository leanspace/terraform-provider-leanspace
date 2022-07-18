TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=leanspace.io
NAMESPACE=io
NAME=leanspace
BINARY=terraform-provider-${NAME}.exe
VERSION=0.3
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
    ifneq ($(filter aarch6%,$(UNAME_P)),)
        ARCHITECTURE=arm64
    endif
endif

OS_ARCH=${PLATFORM}_${ARCHITECTURE}

default: install

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_darwin_arm64
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=linux GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_linux_arm64
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64
	GOOS=windows GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_windows_arm64

install: build
		mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
		mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

install-windows: build
	if not exist %APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS_ARCH} mkdir %APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS_ARCH}
	move ${BINARY} %APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   