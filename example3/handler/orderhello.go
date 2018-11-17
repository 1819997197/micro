package handler

import (
	"context"
	"fmt"
	"will/micro/example3/proto/model"
)

type Order struct {}

func (s *Order) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)

	var head map[string]*model.Pair
	head = make(map[string]*model.Pair)
	head["name"] = &model.Pair{Key:1}
	rsp.Header = head
	rsp.Msg = "hello will"
	rsp.Values = []string{"a", "b", "c"}
	rsp.Type = model.RespType_NONE
	return nil
}