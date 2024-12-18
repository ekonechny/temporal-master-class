package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"

	"temporal-master-class/internal/generated/server"
	"temporal-master-class/internal/generated/temporal"
	"temporal-master-class/internal/workflows/processing"
	tcl_query_builder "temporal-master-class/pkg/tcl-query-builder"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type srv struct {
	server.CustomerServer
	tcl client.Client
	cc  temporal.CustomerClient
	pc  temporal.ProcessingClient
}

func (s *srv) PaymentCallback(ctx context.Context, request *server.PaymentCallbackRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.pc.PaymentCallback(
		ctx,
		evalProcessingWorkflowID(request),
		"",
		&temporal.PaymentCallbackRequest{
			Status: request.Status,
		},
	)
}

func (s *srv) VendorOrderCallback(ctx context.Context, request *server.VendorOrderCallbackRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.pc.VendorOrderCallback(
		ctx,
		evalProcessingWorkflowID(request),
		"",
		&temporal.VendorOrderCallbackRequest{
			Status: request.Status,
		},
	)
}

func (s *srv) GetProfile(ctx context.Context, request *server.GetProfileRequest) (*temporal.Profile, error) {
	return s.cc.GetProfile(
		ctx,
		evalCustomerWorkflowID(request),
		"",
	)
}

func (s *srv) UpdateProfile(ctx context.Context, request *server.UpdateProfileRequest) (*temporal.Profile, error) {
	return s.cc.UpdateProfile(
		ctx,
		evalCustomerWorkflowID(request),
		"",
		&temporal.UpdateProfileRequest{
			Name: request.Name,
		})
}

func (s *srv) DeleteProfile(ctx context.Context, request *server.DeleteProfileRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{},
		s.cc.DeleteProfile(
			ctx,
			evalCustomerWorkflowID(request),
			"",
		)

}

func (s *srv) SetAddress(ctx context.Context, request *server.SetAddressRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{},
		s.cc.SetAddress(
			ctx,
			evalCustomerWorkflowID(request),
			"",
			&temporal.SetAddressRequest{Address: &temporal.Address{
				Title: request.Title,
				Lat:   request.Lat,
				Long:  request.Long,
			}})
}

func (s *srv) GetCart(ctx context.Context, request *server.GetCartRequest) (*temporal.Cart, error) {
	return s.cc.GetCart(
		ctx,
		evalCustomerWorkflowID(request),
		"",
	)
}

func (s *srv) UpdateCart(ctx context.Context, request *server.UpdateCartRequest) (*temporal.Cart, error) {
	return s.cc.UpdateCart(
		ctx,
		evalCustomerWorkflowID(request),
		"",
		&temporal.UpdateCartRequest{
			Products: request.Products,
		})
}

func (s *srv) DeleteCart(ctx context.Context, request *server.DeleteCartRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{},
		s.cc.DeleteCart(
			ctx,
			evalCustomerWorkflowID(request),
			"",
		)

}

func (s *srv) GetOrder(ctx context.Context, request *server.GetOrderRequest) (*temporal.Order, error) {
	return s.pc.GetOrder(ctx, evalProcessingWorkflowID(request), "")
}

func (s *srv) GetOrders(ctx context.Context, request *server.GetOrdersRequest) (*server.GetOrdersResponse, error) {
	qb := tcl_query_builder.NewQueryBuilder()
	qb.AddCondition(tcl_query_builder.Eq("WorkflowType", temporal.ProcessingFlowWorkflowName))
	if request.Phone != "" {
		qb.AddCondition(tcl_query_builder.Eq(processing.CustomerPhone.GetName(), request.Phone))
	}
	if request.CustomerId != "" {
		qb.AddCondition(tcl_query_builder.Eq(processing.CustomerId.GetName(), request.CustomerId))
	}
	if request.Address != "" {
		qb.AddCondition(tcl_query_builder.Eq(processing.CustomerAddress.GetName(), request.Address))
	}

	tQuery := &workflowservice.ListWorkflowExecutionsRequest{
		Query: qb.Query(),
	}

	listResponse, err := s.tcl.ListWorkflow(ctx, tQuery)
	if err != nil {
		return nil, err
	}

	items := make([]*server.ListOrderItem, 0, len(listResponse.Executions))

	conv := converter.GetDefaultDataConverter()
	for _, ex := range listResponse.Executions {
		attrs := ex.GetSearchAttributes()
		item := &server.ListOrderItem{
			Id:        ex.Execution.WorkflowId,
			CreatedAt: ex.StartTime,
		}

		if err = conv.FromPayload(attrs.GetIndexedFields()[processing.CustomerPhone.GetName()], &item.CustomerPhone); err != nil {
			return nil, err
		}
		if err = conv.FromPayload(attrs.GetIndexedFields()[processing.CustomerId.GetName()], &item.CustomerId); err != nil {
			return nil, err
		}
		if err = conv.FromPayload(attrs.GetIndexedFields()[processing.CustomerAddress.GetName()], &item.CustomerAddress); err != nil {
			return nil, err
		}
		items = append(items, item)

	}

	return &server.GetOrdersResponse{
		Orders: items,
	}, nil
}

func (s *srv) Checkout(ctx context.Context, request *server.CheckoutRequest) (*temporal.Order, error) {
	return s.cc.Checkout(
		ctx,
		evalCustomerWorkflowID(request),
		"",
		&temporal.CheckoutRequest{
			PaymentType: request.PaymentType,
		})
}

func (s *srv) NewCustomer(ctx context.Context, in *server.NewCustomerRequest) (*temporal.Profile, error) {
	run, err := s.cc.CustomerFlowAsync(ctx, &temporal.CustomerFlowRequest{
		Name:  in.Name,
		Phone: in.Phone,
	})
	if err != nil {
		return nil, err
	}
	profile, err := s.cc.GetProfile(ctx, run.ID(), run.RunID())
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func evalCustomerWorkflowID(msg interface{ ProtoReflect() protoreflect.Message }) string {
	workflowID, _ := expression.EvalExpression(temporal.CustomerFlowIdexpression, msg.ProtoReflect())
	return workflowID
}

func evalProcessingWorkflowID(msg interface{ ProtoReflect() protoreflect.Message }) string {
	workflowID, _ := expression.EvalExpression(temporal.ProcessingFlowIdexpression, msg.ProtoReflect())
	return workflowID
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	tcl, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	cc := temporal.NewCustomerClient(tcl)
	pc := temporal.NewProcessingClient(tcl)
	s := grpc.NewServer()
	server.RegisterCustomerServer(s, &srv{
		tcl: tcl,
		cc:  cc,
		pc:  pc,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
