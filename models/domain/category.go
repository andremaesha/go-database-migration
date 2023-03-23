package domain

type Category struct {
	Id   int `gorm:"primarykey"`
	Name string
}
