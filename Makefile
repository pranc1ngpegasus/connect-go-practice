.PHONY: buf.mod.update
buf.mod.update:
	@go run github.com/bufbuild/buf/cmd/buf mod update proto

.PHONY: buf.lint
buf.lint:
	@go run github.com/bufbuild/buf/cmd/buf lint

.PHONY: buf.format
buf.format:
	@go run github.com/bufbuild/buf/cmd/buf format -w

.PHONY: generate.buf
generate.buf: buf.mod.update buf.format
	@go run github.com/bufbuild/buf/cmd/buf generate
	$(call goimports)

define goimports
	@go run golang.org/x/tools/cmd/goimports -w ..
endef
