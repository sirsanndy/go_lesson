package entity

type Customer struct {
	Id       int32
	Email    string
	UserName string
	Name     string
	IsActive bool
	Balance  int64
	Married  bool
}
