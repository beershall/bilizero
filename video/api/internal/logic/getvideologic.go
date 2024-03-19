package logic

import (
	"bilizero/user/rpc/types/user"
	"context"
	"fmt"

	"bilizero/video/api/internal/svc"
	"bilizero/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoRes, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "888",
	})
	if err != nil {
		fmt.Printf("call 00 err:%s\n", err.Error())
		return nil, err
	}

	return &types.VideoRes{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
