package handler

import (
	"MaoerMovie/common/response"
	"net/http"

	"MaoerMovie/service/user/api/internal/logic"
	"MaoerMovie/service/user/api/internal/svc"
	"MaoerMovie/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EmailCodeSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailCodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEmailCodeSendLogic(r.Context(), svcCtx)
		resp, err := l.EmailCodeSend(&req)
		response.Response(w, resp, err)
	}
}
