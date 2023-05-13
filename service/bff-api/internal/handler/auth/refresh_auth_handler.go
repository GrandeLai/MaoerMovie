package auth

import (
	"MaoerMovie/service/bff-api/internal/errorx"
	"net/http"

	"MaoerMovie/service/bff-api/internal/logic/auth"
	"MaoerMovie/service/bff-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewRefreshAuthLogic(r.Context(), svcCtx)
		resp, err := l.RefreshAuth(r.Header.Get("Authorization"))
		resp.Status = errorx.ToStatus(resp, err)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
