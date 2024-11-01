package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"temporal-master-class/generated/server"
	"temporal-master-class/generated/temporal"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type srv struct {
	server.OrderServer
	tcl temporal.OrderClient
}

func (s *srv) Create(ctx context.Context, in *temporal.CreateOrderRequest) (*temporal.Order, error) {
	run, err := s.tcl.CreateOrderAsync(ctx, in)
	if err != nil {
		return nil, err
	}
	order, err := s.tcl.Read(ctx, run.ID(), run.RunID())
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *srv) Read(ctx context.Context, in *server.ReadOrderRequest) (*temporal.Order, error) {
	id, err := expression.EvalExpression(temporal.CreateOrderIdexpression, in.ProtoReflect())
	if err != nil {
		return nil, err
	}
	read, err := s.tcl.Read(ctx, id, "")
	if err != nil {
		return nil, err
	}
	return read, nil
}

func (s *srv) Update(ctx context.Context, in *server.UpdateOrderRequest) (*temporal.Order, error) {
	id, err := expression.EvalExpression(temporal.CreateOrderIdexpression, in.ProtoReflect())
	if err != nil {
		return nil, err
	}
	update, err := s.tcl.Update(ctx, id, "", in.Body)
	if err != nil {
		return nil, err
	}
	return update, nil
}

func (s *srv) Delete(ctx context.Context, in *server.DeleteOrderRequest) (*emptypb.Empty, error) {
	err := s.tcl.Delete(ctx, in.Id, "")
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	tcl := temporal.NewOrderClient(c)
	s := grpc.NewServer()
	server.RegisterOrderServer(s, &srv{tcl: tcl})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
