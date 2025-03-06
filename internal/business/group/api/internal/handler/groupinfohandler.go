package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/lhh-gh/IM/internal/business/group/api/internal/logic"
	"github/lhh-gh/IM/internal/business/group/api/internal/svc"
	"github/lhh-gh/IM/internal/business/group/api/internal/types"
	"github/lhh-gh/IM/pkg/response"
)

func groupInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.HttpFail(r, w, err)
			return
		}

		l := logic.NewGroupInfoLogic(r.Context(), svcCtx)
		resp, err := l.GroupInfo(&req)
		if err != nil {
			response.HttpFail(r, w, err)
		} else {
			response.HttpSuccess(r, w, resp)
		}
	}
}
