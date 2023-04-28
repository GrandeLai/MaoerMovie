package handler

import (
	"MaoerMovie/common/response"
	"net/http"

	"MaoerMovie/service/user/api/internal/logic"
	"MaoerMovie/service/user/api/internal/svc"
)

func RefreshAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewRefreshAuthLogic(r.Context(), svcCtx)
		resp, err := l.RefreshAuth(r.Header.Get("Authorization"))
		response.Response(w, resp, err)
	}
}
