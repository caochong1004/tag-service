package main

import (
	"net"
	"log"
)

func main()  {
	s := grpc.NewServer()
	proto.RegisterTagServiceServer(s, server.NewTagServer())
	lis, err := net.Listen("tcp",":"+port)
	if err != nil {
		log.Fatalf("net.Listen err %v", err)
	}
	err = s.Server(lis)
	if err != nil {
		 log.Fatalf("server.Server err: %v", err)
	}
}
