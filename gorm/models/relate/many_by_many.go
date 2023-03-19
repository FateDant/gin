package relate

type ManyOne struct {
	Id      int
	Name    string
	ManyTwo []ManyTwo `gorm:"many2many:one_two_one"` //多对多关系 必须带着个标签
}

type ManyTwo struct {
	Id  int
	Age int
}
