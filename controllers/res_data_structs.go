package controllers

type reqAppData struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type reqResourceData struct {
	ResourceId uint   `json:"resource_id"`
	Path       string `json:"path"`
	Operations []int  `json:"operations"`
}

type reqAuthData struct {
	AuthId    uint   `json:"auth_id"`
	Path      string `json:"path"`
	Operation string `json:"operation"`
}

type IdsData struct {
	Ids []int `json:"ids"`
}
