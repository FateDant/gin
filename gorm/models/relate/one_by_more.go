package relate

import "github.com/jinzhu/gorm"

type One struct {
	gorm.Model
	Name  string
	Mores []More
}

type More struct {
	gorm.Model
	Hobbies string
	OneID   int //关联键 一定得是 模型+ID   OId（×）
}
