pkg_folder=$$(find ./pkg -mindepth 1 -type d -not -path './pkg/pkg*')

all: install build
.PHONY: all

generate-uml:
	goplantuml -recursive . > diagram.puml

version:
	goversion

install:
	go get -t -u ./...
	sudo /usr/local/go/bin/go mod vendor
	sudo chown -R $${USER}:$${USER} go.sum go.mod vendor
	go mod tidy

clean-force: clean
	sudo rm -rf vendor go.sum

clean:
	rm search search.db

build: test build-pkg post-build

build-pkg:
	go build -o search .

pre-test:
	rm -f search || exit 0
	rm -f search.db || exit 0
	rm -f nohup.out || exit 0

post-build: insert-texts

test-lint:
	go fmt . ${pkg_folder}
	go vet . ${pkg_folder}

test-cover:
	go test -cover -race . ${pkg_folder}

test: pre-test test-lint test-cover

insert-texts:
	./scripts/insert_text.sh data
