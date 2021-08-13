package main

import (
	"flag"
	"fmt"
	"github.com/longjoy/tag-service/proto"
	"github.com/longjoy/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
)
var port string

func init()  {
	err  := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err %v", err)
	}
}

func setupFlag() error  {
	flag.StringVar(&port, "port", "8001", "启动端口")
	flag.Parse()
	return nil
}

func main()  {
	s := grpc.NewServer()
	proto.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp",":"+port)
	if err != nil {
		log.Fatalf("net.Listen err %v", err)
	}
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("server.Server err: %v", err)
	}
	fmt.Printf("server.Serve, listen port:%v", port)
}
