package models

type Auth struct {
	BaseModel
	ResourceId  int
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
