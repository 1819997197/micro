package handler

import (
	"context"
	"fmt"
	pb "micro/dcxt/user/proto"
)

type User struct {}

func InitUser() *User {
	return new(User)
}

func (u *User) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq, rsp *pb.GetUserInfoRes) error {
	fmt.Println("received_user: ", req.Id)
	//1.校验参数(略)

	rsp.Msg = "user"
	rsp.Values = []string{"will", "allen", "martin"}
	return nil
}
