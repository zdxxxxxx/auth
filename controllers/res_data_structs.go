package controllers

type reqAppData struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type reqAuthsData struct {
	ResourceId uint   `json:"resource_id"`
	Path       string `json:"path"`
	Operations []int  `json:"operations"`
}
