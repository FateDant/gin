package relate

type UserMaster struct {
	Id   int
	Name string
	Age  int
	Addr string
}

type UserSlave struct {
	Id           int
	Pic          string
	Phone        string
	UserMasterID int
	UserMaster   UserMaster //关联关系 这个字段名得和结构体保持一致 真的坑！！！
}
