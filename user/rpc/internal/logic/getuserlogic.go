package logic

import (
	"context"
	"fmt"

	"bilizero/user/rpc/internal/svc"
	"bilizero/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Printf(">>>>>> call  in:%s\n", in.Id)
	return &user.UserResponse{
		Id:     "123",
		Name:   "beshall",
		Gender: true,
	}, nil
}
