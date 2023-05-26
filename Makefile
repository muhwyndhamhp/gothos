dep:
	@go mod download
	@go mod vendor

run:
	@air -c .air.toml --build.cmd "go build -ldflags \"$(LDFLAGS)\" -o ./tmp/main ."


create-migration:
	@echo "\nfunc migrate_$(shell date +%Y%m%d_%H%M%S)(db *gorm.DB) error {\n\tpanic(\"Unimplemented\")\n}" >> migrations/migration.go

migrate:
	go build -o ./migrations/migration ./migrations/migration.go && chmod 775 ./migrations/migration && ./migrations/migration

lint:
	@golangci-lint run
