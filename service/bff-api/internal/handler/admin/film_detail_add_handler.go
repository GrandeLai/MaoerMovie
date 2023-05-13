package admin

import (
	"MaoerMovie/service/bff-api/internal/errorx"
	"io/ioutil"
	"net/http"

	"MaoerMovie/service/bff-api/internal/logic/admin"
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FilmDetailAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FilmDetailAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		r.ParseMultipartForm(32 << 20) // 32MB max size
		files := r.MultipartForm.File["files"]

		var filmByte [][]byte
		var filmNameList []string
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// 读取文件内容并保存
			data, err := ioutil.ReadAll(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			filmByte = append(filmByte, data)
			filmNameList = append(filmNameList, fileHeader.Filename)
		}

		l := admin.NewFilmDetailAddLogic(r.Context(), svcCtx)
		resp, err := l.FilmDetailAdd(&req, filmByte, filmNameList)
		resp.Status = errorx.ToStatus(resp, err)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
