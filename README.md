Инструкция запуска:

## Сборка proto файфлов

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/datakeeper/data_keeper_service.proto`

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user/user_service.proto`

## Запуск `User` приложения

# Config

`.env|console flag`

`ADDR|--a` - Адрес сервера. По умолчанию `:1234`
`DATA_BASE_DSN|--d` - Адрес базы данных. По умолчанию `postgres://zzman:@localhost:5432/postgres`

Команда для запуска `go run ./cmd/userServer/main`

## Запуск `Storage` приложения

# Config

`.env|console flag`

`ADDR|--a` - Адрес сервера. По умолчанию `:5678`
`USER_SERVICE_ADDR|--ua` - Адрес `User` сервера. По умолчанию `:1234`
`DATA_BASE_DSN|d` - Адрес базы данных. По умолчанию `postgres://zzman:@localhost:5432/postgres`

Команда для запуска `go run ./cmd/storageServer/main`

## Сценарий работы

`User` сервер обращается к базе данных `users`.
Для добавления нового пользователя необходимо добавить в таблицу запись с установленным полем `secret`.
Например следующей командой `INSERT INTO users (secret) VALUES ('уникальный_секрет_пользователя')`.

После того как поднят сервер `User` необходимо запустить сервер `Storage`.
`Storage` сервер обращается к `User` серверу для аутетификации пользователей.

Команда для запуска `Client` приложения `go run ./cmd/client/main`
