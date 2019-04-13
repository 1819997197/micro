package models

type OrderModel struct {
	Id    int64 `json:"id" form:"id"`
	Price int64 `json:"price" form:"price"`
}

func (p OrderModel) TableName() string {
	return "orders"
}
