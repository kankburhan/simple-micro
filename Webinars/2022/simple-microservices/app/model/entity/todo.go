package entity

type Todo struct {
	ID int `gorm:"primary_key;auto_increment;not_null"`

	Todo string `json:"todo"`
}

func (Todo) TableName() string {
	return "todos"
}
