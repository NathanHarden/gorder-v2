package main

import (
	"log"

	"github.com/NathanHarden/gorder-v2/common/config"
	"github.com/spf13/viper"
)

func init(){
	if err:=config.NewViperConfig();err!=nil{
		log.Fatal(err)
	}
}
func main(){
	log.Printf("%v\n",viper.Get("order"))
}