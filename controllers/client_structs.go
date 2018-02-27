package controllers

type OperationClientJson struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AppClientJson struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Path    string `json:"path"`
}

type AuthClientJson struct {
	AppId      int    `json:"app_id"`
	Path       string `json:"path"`
	Operations []int  `json:"operations"`
}

type UserAuthClientJson struct {
	Uid  string `json:"uid"`
	Auth []int  `json:"auth"`
}

type CheckAuthClientJson struct {
	Uid       string `json:"uid"`
	Path      string `json:"path"`
	Operation string `json:"operation"`
}
