SRC = $(shell find . -name "*.go")

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=${BUILD_TIME} -X main.GitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.GitBranch=${GIT_BRANCH}"

bin/damx: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/damx cmd/damx/main.go

.PHONY: install
install: bin/damx
	go run build/damx/install.go $(CURDIR)
	cp bin/damx ${HOME}/.local/bin/

.PHONY: dev_install
dev_install: bin/dpdocs bin/damx
	-rm -r ./testdata
	mkdir -p ./testdata/config
	mkdir -p ./testdata/cache
	BCSPORTAL_CONFIG=./testdata/config BCSPORTAL_CACHE=./testdata/cache go run build/damx/install.go $(CURDIR)
	BCSPORTAL_CONFIG=./testdata/config BCSPORTAL_CACHE=./testdata/cache go run build/dpdocs/install.go $(CURDIR)

.PHONY: dev_db_up
dev_db_up:
	cd tools/migrate/ && go run migrate.go --seed up

.PHONY: dev_db_down
dev_db_down:
	cd tools/migrate/ && go run migrate.go down

.PHONY: dev_db_reset
dev_db_reset:
	cd tools/migrate/ && go run migrate.go --seed reset

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
