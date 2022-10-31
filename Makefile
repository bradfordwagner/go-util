clean:
	@rm -rf ./dist

test:
	@go test ./...

d: dev
c: clean

# watch / develop
dev_pipeline: test
watch:
	@watchexec -cr -f "*.go" -- make dev_pipeline
dev: watch

