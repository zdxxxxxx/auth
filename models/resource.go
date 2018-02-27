package models

import (
	"strings"
)

type Resource struct {
	BaseModel
	AppId uint   `json:"app_id"`
	Name  string `json:"name"`
	Path  string `json:"path" gorm:"not null;unique"`
}

func (o *Resource) Insert() (error) {
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

func GetResourceById(id uint) (*Resource, error) {
	var o Resource
	err := DB.First(&o, "id = ?", id).Error
	return &o, err
}

func DeleteResourceByAppId(id uint) error {
	return DB.Delete(&Resource{}, "app_id = ?", id).Error
}
func GetResourceByAppId(id uint) ([]*Resource, error) {
	var rows []*Resource
	err := DB.Find(&rows, "app_id = ?", id).Error
	return rows, err
}

func DeleteResources(id uint) error {
	var e error
	r, err := GetResourceById(id)
	if err == nil {
		var rows []*Resource
		_err := DB.Find(&rows, "app_id=? AND path LIKE ?", r.AppId, "%"+r.Path+"%").Error
		if _err == nil {
			for _, i := range rows {
				e = i.Delete()
				if e != nil {
					break
				}
			}
		} else {
			e = _err
		}
	} else {
		e = err
	}
	return e
}
func UpdateResources(id uint, path string) error {
	var e error
	r, err := GetResourceById(id)
	if err == nil {
		var rows []*Resource
		_err := DB.Find(&rows, "app_id=? AND path LIKE ?", r.AppId, "%"+r.Path+"%").Error
		if _err == nil {
			for _, i := range rows {
				i.Path = strings.Replace(i.Path, r.Path, path, 1)
				e = i.Update()
				if e != nil {
					break
				}
			}
		} else {
			e = _err
		}
	} else {
		e = err
	}
	return e
}
