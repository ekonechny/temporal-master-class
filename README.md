# Temporal Master Class

Проверьте, что у вас установлено: `go1.23`, `Makefile`

## Первые шаги
Установите [Temporal](https://temporal.io/setup/install-temporal-cli):
```shell
brew install temporal
```

Установите [protoc](https://grpc.io/docs/protoc-installation/):
```shell
brew install protobuf
```

Установите [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/):
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Установите [protoc-gen-go-temporal](https://github.com/cludden/protoc-gen-go-temporal):
```shell
go install github.com/cludden/protoc-gen-go-temporal/cmd/protoc-gen-go_temporal@${PROTOC_GEN_GO_TEMPORAL_VERSION}
```

Подтяните зависимости
```shell
make deps
```

## Структура проекта

```
├── Makefile
├── README.md
├── generated --> Результаты кодгена
│   ├── server
│   │   ├── server.pb.go
│   │   └── server_grpc.pb.go
│   └── temporal
│       ├── order.pb.go
│       └── order_temporal.pb.go
├── go.mod
├── go.sum
├── proto
│   ├── temporal.proto --> Protofile для Temporal, описывающий рабочий процесс
│   └── server.proto -->  Простейший GRPC-сервер
├── server
│   └── main.go --> Команда, для запуска GRPC-сервера
├── server.http --> Примеры запросов к GRPC-серверу
├── worker
│   └── main.go --> Команда, для запуска воркера
└── workflow.go --> Основной workflow
```

## Запуск

Запустить Temporal Dev Server
```shell
make temporal-dev-server
```

Запуск воркера
```shell
make worker
```

Запуск GRPC-сервера
```shell
make server
```

Дальше можно отправлять GRPC-запросы из `server.http`

## Перейти на следующий шаг

```shell
git switch 03-root
```