package models

type Resource struct {
	BaseModel
	AppId int
	Name  string `json:"name" gorm:"not null;unique"`
	Path  string `json:"path" gorm:"not null;unique"`
}

func (o *Resource) Insert() error {
	return DB.Create(o).Error
}

func (o *Resource) Update() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"name":   o.Name,
		"path":   o.Path,
		"app_id": o.AppId,
	}).Error
}

func (o *Resource) Delete() error {
	return DB.Delete(o).Error
}
