package ports

import (
	"context"

	"github.com/NathanHarden/gorder-v2/common/genproto/orderpb"
	"github.com/NathanHarden/gorder-v2/order/app"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct{
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer{
	return &GRPCServer{app:app}
}

func (g GRPCServer) CreateOrder(c context.Context, re *orderpb.CreateOrderRequest) (*emptypb.Empty, error){
	panic("unimplement me")
}

func (g GRPCServer) GetOrder(c context.Context, re *orderpb.GetOrderRequest) (*orderpb.Order, error){
	panic("unimplement me")
}

func (g GRPCServer) UpdateOrder(c context.Context,re *orderpb.Order) (*emptypb.Empty, error){
	panic("unimplement me")
}