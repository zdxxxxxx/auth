package models

type UserAuth struct {
	BaseModel
	Uid    string
	Auth   Auth `gorm:"foreignkey:AuthId"`
	AuthId uint
}

func (o *UserAuth) Insert() error {
	return DB.Create(o).Error
}

func (o *UserAuth) Delete() error {
	return DB.Delete(o).Error
}
