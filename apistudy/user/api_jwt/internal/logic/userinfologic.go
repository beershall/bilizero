package logic

import (
	"context"
	"encoding/json"

	"bilizero/apistudy/user/api_jwt/internal/svc"
	"bilizero/apistudy/user/api_jwt/internal/types"

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
	userId := l.ctx.Value("user_id").(json.Number)
	userName := l.ctx.Value("username").(string)
	uid, err := userId.Int64()
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		Id:       uint(uid),
		UserName: userName,
	}, nil
}
