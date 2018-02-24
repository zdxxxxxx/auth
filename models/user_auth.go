package models

type UserAuth struct {
	BaseModel
	UId    int
	AuthId int
}


func (o *UserAuth) Insert() error {
	return DB.Create(o).Error
}

func (o *UserAuth) Update() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"u_id":   o.UId,
		"auth_id": o.AuthId,
	}).Error
}

func (o *UserAuth) Delete() error {
	return DB.Delete(o).Error
}