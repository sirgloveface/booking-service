# Ejecutar app local
run:
	go run ./cmd/main.go

# Compilar binario
build:
	go build -o bin/app ./cmd/main.go

# Ejecutar en Docker
up:
	docker-compose up --build

# Limpiar binarios
clean:
	rm -rf bin/

# Test
test:
	go test ./...

# Linter
lint:
	golangci-lint run

# Generar dependencias
deps:
	go mod tidy
