// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github/lhh-gh/IM/internal/business/user/api/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/friends",
				Handler: addFriendsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friends",
				Handler: getFriendsHandler(serverCtx),
			},
		},
	)
}
