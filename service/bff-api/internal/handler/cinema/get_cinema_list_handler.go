package cinema

import (
	"MaoerMovie/service/bff-api/internal/errorx"
	"net/http"

	"MaoerMovie/service/bff-api/internal/logic/cinema"
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCinemaListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CinemaListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cinema.NewGetCinemaListLogic(r.Context(), svcCtx)
		resp, err := l.GetCinemaList(&req)
		resp.Status = errorx.ToStatus(resp, err)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
