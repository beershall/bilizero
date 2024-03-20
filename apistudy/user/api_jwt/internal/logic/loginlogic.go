package logic

import (
	"bilizero/common/jwts"
	"context"

	"bilizero/apistudy/user/api_jwt/internal/svc"
	"bilizero/apistudy/user/api_jwt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	auth := l.svcCtx.Config.Auth

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserId:   1,
		Username: "me",
		Role:     3,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "", err
	}

	return token, nil
}
