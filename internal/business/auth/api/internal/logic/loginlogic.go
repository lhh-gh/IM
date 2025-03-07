package logic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github/lhh-gh/IM/internal/business/auth/api/internal/svc"
	"github/lhh-gh/IM/internal/business/auth/api/internal/types"
	"github/lhh-gh/IM/pkg/encrypt"
	"github/lhh-gh/IM/pkg/jwt"
	"gorm.io/gorm"
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

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	userInfo, err := l.svcCtx.MySQL.GetUserAuthInfo(l.ctx, req.ID)
	if err != nil {
		logx.Infof("User %d login failed, err: %v\n", req.ID, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在或密码错误")
		}
		return nil, errors.New("登陆失败！请稍后再试")
	}
	// 校验密码
	password := userInfo.Password
	if !encrypt.CheckPassword(req.Password, password) {
		return nil, errors.New("用户不存在或密码错误")
	}

	// 生成并分发token
	token, err := jwt.GenToken(
		jwt.PayLoad{ID: req.ID, Username: userInfo.Username},
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
	)
	if err != nil {
		logx.Errorf("User %d generate token error, err: %v\n", req.ID, err)
		return nil, errors.New("登陆失败！请稍后再试")
	}

	return &types.LoginResponse{Token: token}, nil
}
