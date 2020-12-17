TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=terraform.gorillastack.com
NAMESPACE=gorillastack
NAME=gorillastack
BINARY=terraform-provider-${NAME}
VERSION=0.3
OS_ARCH=linux_amd64

default: install

build:
	go build -o ${BINARY}
release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   



# TEST?=$$(go list ./... | grep -v 'vendor')
# PKG_NAME=gorillastack
# export GO111MODULE=on

# # Last tagged version
# VERSION = $$(git tag --sort=v:refname | tail -1)

# default: build

# # Builds a binary for current OS and Arch in the plugins dir
# build: fmtcheck
# 	@mkdir -p ~/.terraform.d/plugins/
# 	@go build -o "${HOME}/.terraform.d/plugins/terraform-provider-gorillastack_${VERSION}"

# # Builds a binary for current OS and Arch in the plugins dir
# build-dev: fmtcheck
# 	@mkdir -p ~/.terraform.d/plugins/
# 	@go build -o "${HOME}/.terraform.d/plugins/terraform-provider-gorillastack_${VERSION}_dev"

# # Builds a binary for Linux, Windows, and OSX and installs it in the default terraform plugins directory
# build-plugins: fmtcheck
# 	@mkdir -p ~/.terraform.d/plugins/
# 	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" \
# 	  -output="${HOME}/.terraform.d/plugins/{{.OS}}_{{.Arch}}/terraform-provider-gorillastack_${VERSION}" .

# test: fmtcheck
# 	go test -i $(TEST) || exit 1
# 	echo $(TEST) | \
# 		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

# testacc: fmtcheck
# 	TF_ACC=1 TF_SCHEMA_PANIC_ON_ERROR=1 go test $(TEST) -v $(TESTARGS) -timeout 240m -ldflags="-X=github.com/terraform-providers/terraform-provider-gorillastack/version.ProviderVersion=acc"

# fmt:
# 	@echo "==> Fixing source code with gofmt..."
# 	gofmt -w -s ./$(PKG_NAME)

# # Currently required by tf-deploy compile
# fmtcheck:
# 	@echo "==> Checking source code against gofmt..."
# 	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

# lint:
# 	@echo "==> Checking source code against linters..."
# 	@golangci-lint run ./$(PKG_NAME)

# tools:
# 	@echo "==> installing required tooling..."
# 	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
# 	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

# test-compile:
# 	@if [ "$(TEST)" = "./..." ]; then \
# 		echo "ERROR: Set TEST to a specific package. For example,"; \
# 		echo "  make test-compile TEST=./$(PKG_NAME)"; \
# 		exit 1; \
# 	fi
# 	go test -c $(TEST) $(TESTARGS)

# .PHONY: build test testacc vet fmt fmtcheck lint tools errcheck test-compile

