package handler

import (
	"MaoerMovie/common/response"
	"net/http"

	"MaoerMovie/service/user/api/internal/logic"
	"MaoerMovie/service/user/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)
	}
}
