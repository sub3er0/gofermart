package service

import (
	"gofermart/internal/repository"
)

type OrderService struct {
	OrderRepository *repository.OrderRepository
}

func (or *OrderService) IsOrderExist(orderNumber string) bool {
	isExist, err := or.OrderRepository.IsOrderExist(orderNumber)

	if err != nil {
		return false
	}

	return isExist
}

func (or *OrderService) SaveOrder(orderNumber string, userID int) error {
	err := or.OrderRepository.SaveOrder(orderNumber, userID)

	if err != nil {
		return err
	}

	return nil
}

func (or *OrderService) GetUserOrders(userID int) ([]repository.OrderData, error) {
	orderData, err := or.OrderRepository.GetUserOrders(userID)

	if err != nil {
		return orderData, err
	}

	return orderData, nil
}
