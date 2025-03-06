package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github/lhh-gh/IM/internal/business/group/api/internal/svc"
	"github/lhh-gh/IM/internal/business/group/api/internal/types"
)

type GroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInfoLogic {
	return &GroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupInfoLogic) GroupInfo(req *types.GroupInfoRequest) (resp *types.GroupInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
