package ports

import (
	"context"

	"github.com/NathanHarden/gorder-v2/common/genproto/stockpb"
	"github.com/NathanHarden/gorder-v2/stock/app"
)
type GRPCServer struct{
	app app.Application
}

func NewGRPCServer (app app.Application) *GRPCServer{
	return &GRPCServer{app:app}
}

func (g GRPCServer) GetItems(c context.Context, st *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error){
	panic("unimplement me")
}

func (g GRPCServer) CheckIfItemsInStock(c context.Context, st *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error){
	panic("unimplement me")
}






