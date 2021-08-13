package client

import (
	"context"
	"github.com/longjoy/tag-service/proto"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:8004", nil)
	defer clientConn.Close()
	tagServiceClient := proto.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(
			ctx,
			&proto.GetTagListRequest{Name: "Go"},
		)
	log.Printf("resp:%v",resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error)  {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}

