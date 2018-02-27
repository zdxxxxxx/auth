package models

type Auth struct {
	BaseModel
	Resource    Resource `gorm:"foreignkey:ResourceId"`
	ResourceId  uint
	OperationId int
}

func (o *Auth) Insert() error {
	return DB.Create(o).Error
}

func (o *Auth) Update() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"resource_id":  o.ResourceId,
		"operation_id": o.OperationId,
	}).Error
}

func (o *Auth) Delete() error {
	return DB.Delete(o).Error
}

func GetAuthById(id uint) (*Auth, error) {
	var o Auth
	err := DB.First(&o, "id = ?", id).Error
	return &o, err
}

func GetAuthsByResourceId(id uint) ([]*Auth, error) {
	var ops []*Auth
	err := DB.Find(&ops, "resource_id = ?", id).Error
	return ops, err
}
func DeleteAuthByResourceId(id uint) error {
	return DB.Delete(&Auth{}, "resource_id=?", id).Error
}
