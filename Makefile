SRC = $(shell find . -name "*.go")

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=${BUILD_TIME} -X main.GitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.GitBranch=${GIT_BRANCH}"

bin/damx: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/damx cmd/damx/main.go

bin/dpdocs: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/dpdocs cmd/dpdocs/main.go

.PHONY: install
install: bin/dpdocs bin/damx
	go run build/damx/install.go $(CURDIR)
	go run build/dpdocs/install.go $(CURDIR)
	cp bin/dpdocs ${HOME}/.local/bin/
	cp bin/damx ${HOME}/.local/bin/

.PHONY: install_dev
install_dev: bin/dpdocs bin/damx
	-rm -r ./testdata
	mkdir -p ./testdata/config
	mkdir -p ./testdata/cache
	BCSPORTAL_CONFIG=./testdata/config BCSPORTAL_CACHE=./testdata/cache go run build/damx/install.go $(CURDIR)
	BCSPORTAL_CONFIG=./testdata/config BCSPORTAL_CACHE=./testdata/cache go run build/dpdocs/install.go $(CURDIR)

.PHONY: devdb_migrate
devdb_migrate:
	cd tools/devdb/ && go run devdb.go reset

.PHONY: devdb_reset
devdb_reset:
	cd tools/devdb/ && go run devdb.go --seed reset

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: clean
clean:
	rm -rf bin/


# For future reference when I set this up with a server:
#.PHONY: run
#run: bin/bcs-portal
#	air -d -c .air.conf
