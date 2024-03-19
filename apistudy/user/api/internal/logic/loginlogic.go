package logic

import (
	"context"

	"bilizero/apistudy/user/api/internal/svc"
	"bilizero/apistudy/user/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.Response, err error) {
	resp = &types.Response{}

	resp.Data = req.UserName + "=>" + req.Password
	resp.Code = 200
	resp.Message = "ok"

	return
}
