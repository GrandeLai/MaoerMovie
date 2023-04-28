package handler

import (
	"MaoerMovie/common/response"
	"net/http"

	"MaoerMovie/service/user/api/internal/logic"
	"MaoerMovie/service/user/api/internal/svc"
	"MaoerMovie/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdatePassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePassRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdatePassLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePass(&req)
		response.Response(w, resp, err)
	}
}
