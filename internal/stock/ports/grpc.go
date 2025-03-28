package ports

import (
	"context"
	"github.com/NathanHarden/gorder-v2/common/genproto/orderpb"
	"github.com/NathanHarden/gorder-v2/common/genproto/stockpb"
	"github.com/NathanHarden/gorder-v2/stock/app"
	"github.com/sirupsen/logrus"
)
type GRPCServer struct{
	app app.Application
}

func NewGRPCServer (app app.Application) *GRPCServer{
	return &GRPCServer{app:app}
}

func (g GRPCServer) GetItems(c context.Context, st *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error){
	logrus.Info("rpc request in,stock.GetItems")
	defer func(){
		logrus.Info("rpc request out,stock.GetItems")
	}()
	fake:=[]*orderpb.Item{
		{
			ID:"fake-item-from-GetItems",
		},
	}
	return &stockpb.GetItemsResponse{Items: fake},nil
}

func (g GRPCServer) CheckIfItemsInStock(c context.Context, st *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error){
	logrus.Info("rpc request in,stock.checkIfItemsInstock")
	defer func(){
		logrus.Info("rpc request out,stock.checkIfItemsInstock")
	}()
	return nil,nil
}






