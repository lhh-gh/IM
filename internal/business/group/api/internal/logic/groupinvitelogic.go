package logic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github/lhh-gh/IM/internal/business/group/api/internal/svc"
	"github/lhh-gh/IM/internal/business/group/api/internal/types"
)

type GroupInviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInviteLogic {
	return &GroupInviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GroupInvite 拉人进群，暂时定为邀请人就能直接拉进群而不需要别人同意
// 此处是创建群聊，不会有事先缓存的情况发生，无需更新redis
func (l *GroupInviteLogic) GroupInvite(req *types.GroupInviteRequest) (*types.GroupInviteResponse, error) {
	// 直接塞mysql里
	err := l.svcCtx.MySQL.InsertGroupMembers(l.ctx, req.GroupID, req.Members)
	if err != nil {
		logx.Errorf("[GroupInvite] Insert members to group %d failed, error: %v", req.GroupID, err)
		return nil, errors.New("邀请成员入群失败")
	}
	return nil, nil
}
