package models

type Operation struct {
	Id     int    `json:"id" gorm:"primary_key"`
	Name   string `json:"name" gorm:"not null;unique"`
	Value  string `json:"value" gorm:"not null;unique"`
	Status int    `json:"status" gorm:"default:'1'"`
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

func AddOp() {
	o := []map[string]string{{
		"name":  "查询",
		"value": "READ",
	}, {
		"name":  "新建",
		"value": "CREATE",
	}, {
		"name":  "修改",
		"value": "UPDATE",
	}, {
		"name":  "删除",
		"value": "DELETE",
	}, {
		"name":  "执行",
		"value": "EXECUTE",
	}, {
		"name":  "上传",
		"value": "UPLOAD",
	}, {
		"name":  "下载",
		"value": "DOWNLOAD",
	}}
	for _, item := range o {
		r := &Operation{Name: item["name"], Value: item["value"]}
		r.Insert()
	}
}
