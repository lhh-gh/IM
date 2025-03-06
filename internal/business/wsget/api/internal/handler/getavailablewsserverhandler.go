package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/lhh-gh/IM/internal/business/wsget/api/internal/logic"
	"github/lhh-gh/IM/internal/business/wsget/api/internal/svc"
	"github/lhh-gh/IM/internal/business/wsget/api/internal/types"
	"github/lhh-gh/IM/pkg/response"
)

func getAvailableWSServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebsocketServerGetRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.HttpFail(r, w, err)
			return
		}

		l := logic.NewGetAvailableWSServerLogic(r.Context(), svcCtx)
		resp, err := l.GetAvailableWSServer(&req)
		if err != nil {
			response.HttpFail(r, w, err)
		} else {
			response.HttpSuccess(r, w, resp)
		}
	}
}
