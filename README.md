# Temporal Master Class

Запуск dev сервера для Temporal
```shell
make temporal-dev-server
```

Создать аттрибуты для поиска
```shell
make create-search-attributes
```

Запуск temporal-worker
```shell
make worker
```

Запуск grpc-server
```shell
make server
```

Примеры запросов в `./server.http`

### 01. Hello world

Пишем простейший рабочий процесс с помощью Temporal

Переключаемся на ветку
```shell
git switch 01-helloworld
```


### 02. CRUD 

Учимся создавать и завершать рабочие процессы, а так же получать информацию из них и обновлять ее

```shell
git switch 02-crud
```

### 03. Root Workflow

Создаем рутовый воркфлоу, оркестрирующий основные действия

```shell
git switch 03-root
```

### 04. Activities

Создаем основные активити и дочерние процессы

```shell
git switch 04-activities
```

### 05. Versioning

Учимся работать с версиями

```shell
git switch 05-versioning
```

### 06. Saga

Реализация компенсаций

```shell
git switch 06-saga
```

### 07. Updates

Работа с сигналами и апдейтами

```shell
git switch 07-updates
```

### 08. Search

Индексация и поиск рабочих процессов

```shell
git switch 08-search
```

### 09. Testing

Создание юнит-тестов и replay-тестов

```shell
git switch 09-testing
```

### 10. Observability

Работа с метриками и логированием

```shell
git switch 10-observability
```