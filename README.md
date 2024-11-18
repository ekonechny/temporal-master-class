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
├── cmd
│   ├── server
│   │   └── main.go --> Команда, для запуска GRPC-сервера
│   └── worker
│       └── main.go  --> Команда, для запуска воркера
├── go.mod
├── go.sum
├── internal
│   ├── generated --> Результаты кодгена
│   │   ├── server
│   │   │   ├── server.pb.go
│   │   │   └── server_grpc.pb.go
│   │   └── temporal
│   │       ├── checkout.pb.go
│   │       ├── checkout_temporal.pb.go
│   │       ├── common.pb.go
│   │       ├── customer.pb.go
│   │       ├── customer_temporal.pb.go
│   │       ├── processing.pb.go
│   │       └── processing_temporal.pb.go
│   ├── services --> Моки вызовов в сервисы
│   │   ├── assortment
│   │   │   └── client.go
│   │   ├── payment
│   │   │   └── client.go
│   │   └── vendors
│   │       └── client.go
│   ├── utils
│   │   └── utils.go --> Полезные общие команды
│   └── workflows
│       ├── checkout
│       │   ├── activities
│       │   │   ├── activities.go --> Основная структура активити и все зависимости
│       │   │   ├── activity_assortment_reserve.go
│       │   │   └── activity_create_payment.go
│       │   └── workflow.go --> Workflow, создающий заказ
│       ├── customer
│       │   ├── activities
│       │   │   ├── activities.go --> Основная структура активити и все зависимости
│       │   │   └── activity_assortment_get_products.go  --> Получение продуктов из сервиса ассортиментов
│       │   ├── handler_checkout.go --> Handler для создания заказа
│       │   ├── handler_update_cart.go --> Handler для обновления корзины
│       │   └── workflow.go --> Workflow, описывающий жизненный цикл пользователя
│       └── processing
│           ├── activities
│           │   ├── activities.go --> Основная структура активити и все зависимости
│           │   ├── activity_create_vendor_order.go --> Создание заказа у вендора
│           │   ├── activity_get_payment.go --> Получение статуса платежа
│           │   └── activity_get_vendor_order.go --> Получение статуса заказа у вендора
│           └── workflow.go  --> Workflow для процессинга заказов
├── proto
│   ├── checkout.proto --> Workflow создания заказа
│   ├── common.proto --> Общие сущности
│   ├── customer.proto --> Workflow, описывающий жизненный цикл пользователя
│   ├── processing.proto --> Workflow для процессинга заказа
│   └── server.proto
└── server.http  --> Примеры запросов к GRPC-серверу
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
git switch 05-versioning
```