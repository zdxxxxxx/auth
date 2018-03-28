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

type UserAuthReqData struct {
	Id        int    `json:"id"`
	Uid       string `json:"uid"`
	App       string `json:"app"`
	Path      string `json:"path"`
	Operation string `json:"operation"`
}

type AppReqData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (r *ReqData) SetResult(s int, d interface{}) {
	r.Code = s
	r.Msg = status[s]
	r.Data = d
}
