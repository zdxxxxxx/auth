package models

type UserAuth struct {
	BaseModel
	Uid    string
	AuthId uint
	AppId  uint
	Path   string
}

func (o *UserAuth) Insert() error {
	return DB.Create(o).Error
}

func (o *UserAuth) Delete() error {
	return DB.Delete(o).Error
}

func GetUserAuthById(id uint) (*UserAuth, error) {
	var o UserAuth
	err := DB.First(&o, "id = ?", id).Error
	return &o, err
}

func GetUserAuthByAuthId(id uint) ([]*UserAuth, error) {
	var o []*UserAuth
	err := DB.Find(&o, "auth_id = ?", id).Error
	return o, err
}

func GetUserAuthByUid(uid string) ([]*UserAuth, error) {
	var o []*UserAuth
	err := DB.Find(&o, "uid = ?", uid).Error
	return o, err
}

func GetUserAuth(params ...interface{}) ([]*UserAuth, error) {
	var str = make([]string, 0)
	var err error
	var value = make([]interface{}, 0)
	var o []*UserAuth

	filter := []string{"uid", "app_id", "path"}
	for index, param := range params {
		if param == "" || param == 0 {
			continue
		}
		str = append(str, filter[index])
		value = append(value, param)
	}
	if len(str) == 1 {
		err = DB.Find(&o, str[0]+"= ?", value[0]).Error
	} else if len(str) == 2 {
		err = DB.Find(&o, str[0]+"= ? AND "+str[1]+"= ? ", value[0], value[1]).Error
	} else if len(str) == 3 {
		err = DB.Find(&o, str[0]+"= ? AND "+str[1]+"= ? AND "+str[2]+"= ? ", value[0], value[1], value[3]).Error

	}
	return o, err;
}
