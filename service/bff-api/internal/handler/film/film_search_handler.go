package film

import (
	"MaoerMovie/service/bff-api/internal/errorx"
	"net/http"

	"MaoerMovie/service/bff-api/internal/logic/film"
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FilmSearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FilmSearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := film.NewFilmSearchLogic(r.Context(), svcCtx)
		resp, err := l.FilmSearch(&req)
		resp.Status = errorx.ToStatus(resp, err)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
