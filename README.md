# Wallet Service

Тестовое задание сервис кошелек от ITK Academy

## API Endpoints

- `POST /api/v1/wallet` - обновление кошелька (DEPOSIT/WITHDRAW)
- `GET /api/v1/wallets/{id}` - получение баланса

## Запуск

### Локально
```bash
go run cmd/main.go
```

### Docker
```bash
docker-compose up --build -d
```

## Тестирование 
```bash
go test ./tests -v
```

## Переменные окружения в файле config.env
```txt
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=w4ll3t
POSTGRES_PASSWORD=w4ll3t
POSTGRES_DB=w4ll3t
SERVER_PORT=8080
```
