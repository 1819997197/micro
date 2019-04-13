package dao

import (
	"github.com/jinzhu/gorm"
	"micro/ch03/conn"
	"micro/ch03/models"
)

type OrderDao struct{}

func InitOrderDao() *OrderDao {
	return new(OrderDao)
}

//gorm查询单条记录
func (p *OrderDao) GetPerson(id int64) (*models.OrderModel, error) {
	item := &models.OrderModel{}

	err := conn.SqlDB.Where("id = ? ", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return item, nil
}
