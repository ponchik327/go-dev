# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:12345678@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up ${COUNT}

# Откат миграций
migrate-down:
	$(MIGRATE) down ${COUNT}
	
# Принудительно устанавливает нужную версию
migrate-force:
	$(MIGRATE) force ${VERSION}

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

lint:
	golangci-lint config -c linter/.golangci.yaml > /dev/null
	golangci-lint run --out-format=colored-line-number