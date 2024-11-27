.PHONY: worker server

export PROTOC_GEN_GO_TEMPORAL_VERSION=v1.14.3
export TEMPORAL_DEBUG=true

init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/cludden/protoc-gen-go-temporal/cmd/protoc-gen-go_temporal@${PROTOC_GEN_GO_TEMPORAL_VERSION}

gen: gen-temporal gen-server

gen-temporal:
	@mkdir -p "./internal/generated/temporal"
	@protoc \
     -I ./proto \
     -I ${GOPATH}/pkg/mod/github.com/cludden/protoc-gen-go-temporal@${PROTOC_GEN_GO_TEMPORAL_VERSION}/proto \
     --go_out=../ \
     --go_opt=paths=import \
     --plugin=protoc-gen-go-temporal \
     --go_temporal_out=../ \
     --go_temporal_opt="cli-categories=true" \
     --go_temporal_opt="cli-enabled=true" \
     --go_temporal_opt="workflow-update-enabled=true" \
     proto/common.proto proto/processing.proto proto/customer.proto proto/checkout.proto

gen-server:
	@mkdir -p "./internal/generated/server"
	@protoc \
     -I ./proto \
     -I ${GOPATH}/pkg/mod/github.com/cludden/protoc-gen-go-temporal@${PROTOC_GEN_GO_TEMPORAL_VERSION}/proto \
     --go_out=../ \
     --go-grpc_out=../ \
     --plugin=protoc-gen-go-temporal \
     proto/server.proto

temporal-dev-server:
	temporal server start-dev \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true" \
      --ui-port 8080

create-search-attributes: sleep
	temporal operator search-attribute create --namespace "default" \
            --name="CustomerPhone" --type="Text" \
            --name="CustomerId" --type="Text" \
            --name="CustomerAddress" --type="Text"

deps:
	go mod tidy

worker:
	go run cmd/worker/main.go worker

server:
	go run cmd/server/main.go

test:
	@go test ./internal/... -cover -short -count=1

# make -j4 all
all: temporal-dev-server create-search-attributes worker server

# Ленивый хак, чтобы дождаться пока запустится temporal-dev-server и создать там индексы
sleep:
	@sleep 1

