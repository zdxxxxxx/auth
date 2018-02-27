package models

type App struct {
	BaseModel
	Name       string `json:"name" gorm:"not null;unique"`
	Status     int    `json:"status" gorm:"default:'1'"`
	Desc       string `json:"desc"`
	ResourceId uint   `json:"resource_id"`
	Content    string `json:"content" gorm:"size:4096"`
}

func (o *App) Insert() error {
	return DB.Create(o).Error
}

func (o *App) Update() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"name": o.Name,
		"desc": o.Desc,
	}).Error
}

func (o *App) UpdateContent() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"content": o.Content,
	}).Error
}
func (o *App) UpdateResourceId() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"resource_id": o.ResourceId,
	}).Error
}
func (o *App) UpdateStatus() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"status": o.Status,
	}).Error
}

func (o *App) Delete() error {
	return DB.Delete(o).Error
}

func GetAppById(id uint) (*App, error) {
	var o App
	err := DB.First(&o, "id = ?", id).Error
	return &o, err
}
func GetAppAll() ([]*App, error) {
	var ops []*App
	err := DB.Find(&ops).Error
	return ops, err
}
