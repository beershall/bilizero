package handler

import (
	"bilizero/apistudy/user/api_v2/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func PackResponse(r *http.Request, w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		httpx.WriteJson(w, http.StatusOK, types.Response{
			Code: 8,
			Msg:  err.Error(),
			Data: nil,
		})
	} else {
		httpx.WriteJson(w, http.StatusOK, types.Response{
			Code: 0,
			Msg:  "ok",
			Data: data,
		})
	}
}
