package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github/lhh-gh/IM/internal/business/info/api/internal/svc"
	"github/lhh-gh/IM/internal/business/info/api/internal/types"
)

type MessageRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageRecordLogic {
	return &MessageRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageRecordLogic) MessageRecord(req *types.MessageRecordRequest) (resp *types.MessageRecordResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
