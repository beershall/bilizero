package logic

import (
	"context"

	"bilizero/apistudy/user/api/internal/svc"
	"bilizero/apistudy/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserInfoResponse{}

	resp.Code = 201
	resp.Message = "ok2"
	resp.Data = types.UserInfo{
		UserName: "me",
		Addr:     "here",
		Id:       8,
	}

	return
}
