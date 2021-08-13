package server

import (
	"context"
	"github.com/longjoy/tag-service/pkg/bapi"
	"github.com/longjoy/tag-service/pkg/errcode"
	"github.com/longjoy/tag-service/proto"

	"encoding/json"
)


type TagServer struct {}

func NewTagServer() *TagServer  {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *proto.GetTagListRequest) (*proto.GetTagListReply, error) {
	//resp, err := http.Get(
	//	"http://127.0.0.1:8000/api/v1/tags?name=" + r.GetName(),
	//	)
	//if err != nil {
	//	return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	//}
	api := 	bapi.NewAPI("127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagList := proto.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return &tagList, nil
}