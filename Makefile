.PHONY: worker

export TEMPORAL_DEBUG=true

temporal-dev-server:
	temporal server start-dev \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true" \
      --ui-port 8080

deps:
	go mod tidy

worker:
	go run worker/main.go

start:
	go run starter/main.go
