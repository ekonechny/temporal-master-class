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

Подтяните зависимости
```shell
make deps
```

## Структура проекта

```
├── Makefile --> Здесь команды запуска
├── README.md                  
├── go.mod
├── go.sum
├── helloworld.go --> Workflow, выполняющий Hello и Bye Activity
├── starter
│   └── main.go --> Команда для запуска нового экземпляра workflow
└── worker
    └── main.go --> Команда для запуска воркера, который будет исполнять workflow
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

Запуск нового экземпляра рабочего процесса
```shell
make start
```

## Перейти на следующий шаг

```shell
git switch 02-crud
```