package main

import (
	"log"
	"github.com/NathanHarden/gorder-v2/common/config"
	"github.com/NathanHarden/gorder-v2/common/genproto/orderpb"
	"github.com/NathanHarden/gorder-v2/common/server"
	"github.com/NathanHarden/gorder-v2/order/ports"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init(){
	if err:=config.NewViperConfig();err!=nil{
		log.Fatal(err)
	}
}

type HTTPServer struct{

}

func (s HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerId string){

}

func (s HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string){

}
func main(){
	serviceName:=viper.GetString("order.service-name")
		go server.RunGRPCServer(serviceName,func(server *grpc.Server){
			svc:=ports.NewGRPCServer()
			orderpb.RegisterOrderServiceServer(server,svc)

		})
		
		server.RunHTTPServer(serviceName,func (router *gin.Engine){
			ports.RegisterHandlersWithOptions(router,HTTPServer{},ports.GinServerOptions{
				BaseURL: "/api",
				Middlewares: nil,
				ErrorHandler: nil,
			})
		})
}