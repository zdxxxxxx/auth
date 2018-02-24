package utils

var (
	status = map[int]string{
		0:   "success",
		100: "服务端异常",
		101: "参数错误",
	}
)

type ReqData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *ReqData) SetResult(s int, d interface{}) {
	r.Code = s
	r.Msg = status[s]
	r.Data = d
}
