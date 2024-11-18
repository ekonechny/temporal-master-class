.PHONY: worker server

export PROTOC_GEN_GO_TEMPORAL_VERSION=v1.14.3
export TEMPORAL_DEBUG=true

init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/cludden/protoc-gen-go-temporal/cmd/protoc-gen-go_temporal@${PROTOC_GEN_GO_TEMPORAL_VERSION}

gen: gen-temporal gen-server

gen-temporal:
	protoc \
     -I ./proto \
     -I ${GOPATH}/pkg/mod/github.com/cludden/protoc-gen-go-temporal@${PROTOC_GEN_GO_TEMPORAL_VERSION}/proto \
     --go_out=../ \
     --go_opt=paths=import \
     --plugin=protoc-gen-go-temporal \
     --go_temporal_out=../ \
     --go_temporal_opt="cli-categories=true" \
     --go_temporal_opt="cli-enabled=true" \
     --go_temporal_opt="workflow-update-enabled=true" \
     proto/temporal.proto

gen-server:
	protoc \
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

deps:
	go mod tidy

worker:
	go run worker/main.go worker

server:
	go run server/main.go