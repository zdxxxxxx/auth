package models

type App struct {
	BaseModel
	Name   string `json:"name" gorm:"not null;unique"`
	Status int    `json:"status" gorm:"default:'0'"`
	Path   string `json:"path" gorm:"not null;unique"`
	Desc   string `json:"desc"`
}

func (o *App) Insert() error {
	return DB.Create(o).Error
}

func (o *App) Update() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"name":   o.Name,
		"status": o.Status,
		"path":   o.Path,
		"desc":   o.Desc,
	}).Error
}

func (o *App) Delete() error {
	return DB.Delete(o).Error
}
