package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/lhh-gh/IM/internal/business/user/api/internal/logic"
	"github/lhh-gh/IM/internal/business/user/api/internal/svc"
	"github/lhh-gh/IM/internal/business/user/api/internal/types"
	"github/lhh-gh/IM/pkg/response"
)

func getFriendsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFriendsRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.HttpFail(r, w, err)
			return
		}

		l := logic.NewGetFriendsLogic(r.Context(), svcCtx)
		resp, err := l.GetFriends(&req)
		if err != nil {
			response.HttpFail(r, w, err)
		} else {
			response.HttpSuccess(r, w, resp)
		}
	}
}
