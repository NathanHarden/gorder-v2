package main

import (
	"context"

	"github.com/NathanHarden/gorder-v2/common/config"
	"github.com/NathanHarden/gorder-v2/common/genproto/orderpb"
	"github.com/NathanHarden/gorder-v2/common/server"
	"github.com/NathanHarden/gorder-v2/order/app"
	"github.com/NathanHarden/gorder-v2/order/ports"
	"github.com/NathanHarden/gorder-v2/order/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init(){
	if err:=config.NewViperConfig();err!=nil{
		logrus.Fatal(err)
	}
}

type HTTPServer struct{
	app app.Application
}

func (s HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerId string){

}

func (s HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string){

}
func main(){
	serviceName:=viper.GetString("order.service-name")
	ctx,cancel:=context.WithCancel(context.Background())
	app:=service.NewApplication(ctx)
	defer cancel()

		go server.RunGRPCServer(serviceName,func(server *grpc.Server){
			svc:=ports.NewGRPCServer(app)
			orderpb.RegisterOrderServiceServer(server,svc)

		})
		
		server.RunHTTPServer(serviceName,func (router *gin.Engine){
			ports.RegisterHandlersWithOptions(router,HTTPServer{app:app},ports.GinServerOptions{
				BaseURL: "/api",
				Middlewares: nil,
				ErrorHandler: nil,
			})
		})
}