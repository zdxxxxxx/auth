package models

type Operation struct {
	Id     int    `json:"id" gorm:"primary_key"`
	Name   string `json:"name" gorm:"not null;unique"`
	Value  string `json:"value" gorm:"not null;unique"`
	Status int    `json:"status" gorm:"default:'0'"`
}

func (o *Operation) Insert() error {
	return DB.Create(o).Error
}

func (o *Operation) Update() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"name":  o.Name,
		"value": o.Value,
	}).Error
}

func (o *Operation) Delete() error {
	return DB.Delete(o).Error
}

func (o *Operation) UpdateStatus() error {
	return DB.Model(o).Updates(map[string]interface{}{
		"status": o.Status,
	}).Error
}

func GetOperationById(id int) (*Operation, error) {
	var o Operation
	err := DB.First(&o, "id = ?", id).Error
	return &o, err
}

func GetOperationAll() ([]*Operation, error) {
	var ops []*Operation
	err := DB.Find(&ops).Error
	return ops, err
}
