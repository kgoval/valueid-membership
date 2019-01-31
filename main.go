package main

import (
	"fmt"
	"github.com/kgoval/envparser"
	"github.com/kgoval/erresp"
	"github.com/kgoval/valueid-membership/boot"
	"github.com/kgoval/valueid-membership/handler"
	local "github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"strings"

	//"github.com/micro/go-micro"
)

func init()  {
	//-- registering error message
	erresp.Register(ErrorMessage)

	//-- parse config from env
	if err := envparser.Parse(boot.Config); err != nil{
		panic(err)
	}

}



func main(){

	fmt.Println("engine started..")
	var service micro.Service

	opts := micro.Option(func(*micro.Options){
		micro.Name(boot.Config.ServiceName)
	})

	if strings.ToLower(boot.Config.Environment) == "local" {
		service = local.NewService(opts)
	}else{
		service = k8s.NewService(opts)
	}

	service.Init()

	handler.Register(service.Server())

	fmt.Println("running..")
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
