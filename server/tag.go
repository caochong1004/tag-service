package server

import (
	"context"
	"errors"

	"github.com/longjoy/tag-service/pkg/bapi"
	"github.com/longjoy/tag-service/proto"

	"encoding/json"
)


type TagServer struct {}

func NewTagServer() *TagServer  {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *proto.GetTagListRequest) (*proto.GetTagListReply, error) {
	api := 	bapi.NewAPI("127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagList := proto.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errors.New("grpc 失败")
	}
	return &tagList, nil
}