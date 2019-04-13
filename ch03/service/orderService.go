package service

import (
	"fmt"
	"micro/ch03/dao"
)

type OrderService struct {
	OrderDao *dao.OrderDao
}

func InitOrderService() *OrderService {
	o := new(OrderService)
	o.OrderDao = dao.InitOrderDao()
	return o
}

func (o *OrderService) GetOrderById(id int64) (string, error) {
	order, err := o.OrderDao.GetPerson(id)
	if err != nil {
		fmt.Println("o.OrderDao.GetPerson err:", err)
		return "", err
	}

	if order == nil {
		return "", nil

	}
	return fmt.Sprintf("id:%d, price:%d", order.Id, order.Price), nil
}
