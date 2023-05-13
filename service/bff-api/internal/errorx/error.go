package errorx

import "MaoerMovie/service/bff-api/internal/types"

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.Msg
}

func ToStatus(resp interface{}, err error) types.Status {
	var status types.Status
	if err != nil {
		status.Code = -1
		status.Msg = err.Error()
	} else {
		status.Code = 200
		status.Msg = "OK"
	}
	return status
}
