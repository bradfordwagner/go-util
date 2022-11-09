install_binaries:
	go install github.com/golang/mock/mockgen@v1.6.0

generate_mocks:
	go generate ./...

install: install_binaries generate_mocks

clean:
	@rm -rf ./dist

test:
	@go test ./...

d: dev
c: clean
i: install

# watch / develop
dev_pipeline: test
watch:
	@watchexec -cr -f "*.go" -- make dev_pipeline
dev: watch

