# Go Task API booking-service 🐹

API RESTful escrita en Go, con soporte para PostgreSQL y MongoDB.

## ✨ Features
- GIN como framework web
- GORM para PostgreSQL
- MongoDB con driver oficial
- Docker & Docker Compose
- Arquitectura limpia (handler, service, repository)
- Selección de DB vía variable de entorno

## 🚀 Ejecutar localmente

```bash
docker-compose up --build
```


## 🧪 Endpoints
```bash
GET /tasks

POST /tasks { "title": "string" }

GET /tasks/:id
```


## Make Commands
```bash
# Ejecutar app local
make run
# Compilar binario
make build
# Ejecutar en Docker
make up
# Limpiar binarios
make clean
# Test
make test
# Linter
make lint
# Generar dependencias
make deps
```