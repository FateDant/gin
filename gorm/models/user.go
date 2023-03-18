package models

type User struct {
	Id    int
	Name  string
	Age   string
	Pic   string
	Phone string
}

func (User) TableName() string {
	//表名重命名
	return "rename_user"
}
